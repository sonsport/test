package inviter

import (
	"context"
	"strings"
	"time"

	"github.com/druidcaesa/ztool"
	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/consts/config"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/entity"
	"fuya-ark/internal/service"
	"fuya-ark/utility"
)

type sInviter struct{}

func init() {
	service.RegisterInviter(New())
}

func New() *sInviter {
	return &sInviter{}
}

// ShareChannels 分享
func (s *sInviter) ShareChannels(ctx context.Context, in model.ShareChannelsReq) (out model.ShareChannelsRes, err error) {
	userId := gconv.Int64(ctx.Value(consts.CtxUserId))
	//appName := gconv.String(ctx.Value(consts.CtxAppName))
	shareId, err := GetShareId(ctx, userId)
	if err != nil || shareId <= 0 {
		return model.ShareChannelsRes{}, err
	}
	//ad融合邀新
	source := "ad_merge"
	shareUrl := strings.ReplaceAll(consts.AdjustShareUrl, "#adj_label", gconv.String(userId))
	urlStr := getShortUrl(ctx, userId, shareUrl, source)
	return model.ShareChannelsRes{Link: urlStr}, err
}

// NewUserLoginReward 新用户绑定
func (s *sInviter) NewUserLoginReward(ctx context.Context, in model.NewUserLoginRewardReq) (out model.NewUserLoginRewardRes, err error) {
	userId := gconv.Int64(ctx.Value(consts.CtxUserId))
	//appName := gconv.String(ctx.Value(consts.CtxAppName))
	appChannel := gconv.String(ctx.Value(consts.CtxAppChannel))
	deviceId := gconv.String(ctx.Value(consts.CtxDeviceId))
	clientIP := gconv.String(ctx.Value(consts.CtxClientIP))
	clientMode := gconv.String(ctx.Value(consts.CtxClientMode))

	inviterId := in.InviterId

	redisKey := consts.RedisLockFeedLike + gconv.String(userId) + gconv.String(inviterId)
	redisLock := consts.RedisLock.NewMutex(redisKey)
	err = redisLock.Lock()
	if err != nil {
		return out, err
	}
	defer func(redisLock *redsync.Mutex) {
		_, err := redisLock.Unlock()
		if err != nil {
		}
	}(redisLock)
	////检查用户邀新权限
	if checkNewUserPermission(ctx, inviterId) <= 0 { //邀新权限不足
		SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 9)
		return model.NewUserLoginRewardRes{ErrCode: -66}, nil
	}

	if inviterId == userId {
		g.Log().Warningf(ctx, "chyInfo 邀请新人 自己不能邀请自己,邀请者 = %v，被邀请者 = %v", inviterId, userId)
		return model.NewUserLoginRewardRes{ErrCode: -65}, nil
	}

	//判断设备号是否正常
	if len(deviceId) == 0 || inviterId == 0 || len(clientMode) <= 5 {
		g.Log().Warningf(ctx, "chyInfo 邀请新人 邀请设备为空/InviterId为空，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
		SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 1)
		return model.NewUserLoginRewardRes{ErrCode: 4}, nil
	}
	guildAnchor, err := dao.GuildAnchor.GetGuildAnchor(ctx, inviterId, consts.GuildAnchorStatusOk)
	startTime := ztool.DateUtils.Now().StartTimeOfDay().Time().Unix()
	invitedCount, err := dao.ShareNewUserInfo.GetInvitedUserCountByTime(ctx, inviterId, startTime)

	if guildAnchor == nil || guildAnchor.StarLevel <= 0 { //不是星级主播 邀新限制2个
		balanceInfo := dao.BalanceInfo.GetByUserId(ctx, inviterId)
		if balanceInfo != nil && balanceInfo.TotalFee < 20000000 && invitedCount >= 2 { //充值小于200k 每天限制2个
			g.Log().Warningf(ctx, "chyInfo 充值小于200k 每天限制2个，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 2)
			return model.NewUserLoginRewardRes{ErrCode: 66}, nil
		}
	}

	//判断当天邀请人数大于20
	if invitedCount >= 20 {
		g.Log().Warningf(ctx, "chyInfo 邀请新人 当天邀请人数大于20，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
		//rongcloud.SendSystemByText(inviterId, fmt.Sprintf("Hari ini Anda telah mencapai batas undangan %v orang, silakan coba lagi besok.", invitedCount))
		SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 3)
		return model.NewUserLoginRewardRes{ErrCode: 66}, nil
	}
	if invitedCount > 15 { //发送运营号通知
		err = ShareMaxNotice(ctx, inviterId, invitedCount)
		if err != nil {
			g.Log().Errorf(ctx, "通知运营号邀新用户 %v 邀新达到上限失败 err=%v", userId, err)
		}
	}
	// 校验新用户是否已参加过活动
	shareNewUserInfo, _ := dao.ShareNewUserInfo.GetShareNewUserInfo(ctx, userId, 0)
	isOldUser := false
	userInfo, err := dao.UserInfo.GetUserById(ctx, userId)
	balance := dao.BalanceInfo.GetByUserId(ctx, userId)
	if appChannel == "fuya_google" && time.Now().Unix()-userInfo.LastHeartbeatAt > consts.OneDaySeconds*30 &&
		balance != nil && balance.TotalFee > 0 { //邀请时间大于30天的fuya充值用户 可以再次绑定一个邀请 邀请大于30天的充值用户
		isOldUser = true
	}
	//4星性感主播邀请新用户给2000钻石
	if guildAnchor != nil && guildAnchor.StarLevel >= consts.AnchorStar4 &&
		utility.IsCharmAnchor(ctx, guildAnchor.Label) {
		isOldUser = true
	}
	aAGuildCount, _ := dao.AnchorLevel.GetCountByGuildId(ctx, inviterId)
	//有AA类主播的工会长邀新奖励增至2000钻
	if aAGuildCount > 0 {
		g.Log().Warningf(ctx, "chyInfo 有AA类主播的工会长邀新奖励增至2000钻/人  ip = %v，aAGuildCount: =%v , 被邀请者 = %v", clientIP, aAGuildCount, userId)
		isOldUser = true
	}
	if shareNewUserInfo == nil && !isOldUser { //检查是否印尼、马来ip
		region, _ := utility.SearchIp(clientIP)
		if dao.AreaInfo.GetByTitleCnAndIsShow(ctx, region) == nil {
			g.Log().Warningf(ctx, "chyInfo 邀请新人 检查是否支持的地区  ip = %v，region: =%v , 被邀请者 = %v", clientIP, region, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 4)
			return model.NewUserLoginRewardRes{ErrCode: 11}, nil
		}

		// 检测该设备是否已经参加过奖励
		shareNewUserInfoByDeviceId, _ := dao.ShareNewUserInfo.GetShareNewUserInfoBySmId(ctx, deviceId)
		if shareNewUserInfoByDeviceId != nil {
			g.Log().Warningf(ctx, "chyInfo 邀请新人 设备已经完成绑定,设备号= %v,邀请者 = %v，被邀请者 = %v", deviceId, inviterId, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 5)
			return model.NewUserLoginRewardRes{ErrCode: 12}, nil
		}

		//查找相似或类似的ip地址,一个ip地址只能被成功邀请一次，不能多次邀请
		same := SameTheIP(ctx, inviterId, clientIP)
		if same {
			g.Log().Warningf(ctx, "chyInfo 邀请新人 该IP地址已经被邀请过，ip = %v，不能重复邀请,邀请者 = %v，被邀请者 = %v", clientIP, inviterId, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 6)
			return model.NewUserLoginRewardRes{ErrCode: 14}, nil
		}

		//检查邀请者是否有分享权限,当前默认全开放
		ok := CheckInviterRole(ctx, inviterId)
		if !ok {
			g.Log().Warningf(ctx, "chyInfo 邀请新人 邀请人邀请资格失效,邀请者 = %v，被邀请者 = %v", inviterId, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 7)
			return model.NewUserLoginRewardRes{ErrCode: 8}, nil
		} else if shareNewUserInfo != nil {
			g.Log().Warningf(ctx, "chyInfo 邀请新人 用户已经完成绑定,邀请者 = %v，被邀请者 = %v", inviterId, userId)
			SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 8)
			return model.NewUserLoginRewardRes{ErrCode: -10}, nil
		}

		////记录被邀请者/邀请者信息 运营号10001和10002的分享链接替换成了30001、30002
		if inviterId == 10001 {
			inviterId = 30001
		} else if inviterId == 10002 {
			inviterId = 30002
		}
		if err = NewUserLoad(
			ctx, userId, inviterId, clientIP, deviceId, clientMode, isOldUser); err != nil {
			return model.NewUserLoginRewardRes{ErrCode: -11}, err
		}
		newUserRewardDiamonds := config.ConfigS.UrlConfig.NewUserRewardDiamonds
		//加邀请奖励
		if err = NewUserReward(ctx, userId, inviterId, newUserRewardDiamonds, isOldUser, consts.ShareByWatchLive); err != nil {
			return model.NewUserLoginRewardRes{ErrCode: -99}, err
		}
		//A类主播邀请过来的用户发放奖励 7天--坐骑--Pink Romance 3天--头像框--Blue tears
		isAAnchor, _ := dao.AnchorLevel.IsALevelAnchor(ctx, inviterId)
		if isAAnchor > 0 {
			g.Log().Warningf(ctx, "chyInfo A类主播邀请过来的用户发放奖2000钻石 ip = %v，aAGuildCount: =%v , 被邀请者 = %v", clientIP, aAGuildCount, userId)
			isOldUser = true
			//todo
			//mallDomain.GiftFreeGood(ctx, userId, 1, 7)
			//mallDomain.GiftFreeGood(ctx, userId, 19, 3)
			//text := fmt.Sprintf("Pengguna terhormat %v, Selamat join di Fuya Live, anda diundang dari teman %v dan Selamat anda mendapat hadiah efek khusus dari event Fuya Live : 7 hari Naiki Pink Romance dan 3 hari Avatar Frame Blue Tears!!", userId, inviterId)
			//rongcloud.SendOfficialSystemMsgV2(ctx, consts.TitleTypeSystem, 0, 0, 0, userInfo, userInfo, nil, text)
		}

		//运营号绑定
		inviterInfo, _ := dao.UserInfo.GetUserById(ctx, inviterId)
		userBaseInfo := entity.UserBaseInfo{}
		if inviterInfo != nil {
			if inviterInfo.Role == consts.UserServer { //运营号
				userBaseInfo.ServerId = inviterId
			} else {
				guildInfo, _ := dao.GuildInfo.GetGuildServiceIdByUserId(ctx, inviterId, consts.GuildAnchorStatusOk)
				if guildInfo != nil && guildInfo.ServerId > 0 { //公会长、主播有运营号
					userBaseInfo.ServerId = guildInfo.ServerId
				} else { //用户
					baseInfo, _ := dao.UserBaseInfo.GetUserBaseInfoByUserId(ctx, inviterId)
					if baseInfo != nil && baseInfo.ServerId != 0 {
						userBaseInfo.ServerId = baseInfo.ServerId
					} else {
						userBaseInfo.ServerId = 10000
					}
				}
			}
		}
		data, err := dao.UserBaseInfo.GetUserBaseInfo(ctx, userId)
		if err != nil {
			return model.NewUserLoginRewardRes{}, err
		}
		if data == nil {
			err = dao.UserBaseInfo.SaveUserBaseInfo(ctx, userBaseInfo, nil)
			if err != nil {
				g.Log().Errorf(ctx, "domain/user/user_base.go:35 err:%v", err)
				return model.NewUserLoginRewardRes{}, err
			}
		} else {
			err = dao.UserBaseInfo.UpdateUserBaseServer(ctx, userId, userBaseInfo.ServerId)
			if err != nil {
				g.Log().Error(ctx, "更新用户运营号失败 err = %v", err)
				return model.NewUserLoginRewardRes{}, err
			}
		}
	}
	return model.NewUserLoginRewardRes{}, err
}

// GetShareId 获取分享ID
func GetShareId(ctx context.Context, userId int64) (shareId int64, err error) {
	var inviterInfo *entity.InviterInfo
	//这里还需要判断用户是否具备邀请新人权限，如果不满足条件，需要记录一下，但不限制分享
	err = dao.InviterInfo.Ctx(ctx).Where(dao.InviterInfo.Columns().UserId, userId).Scan(ctx, &inviterInfo)
	if err != nil {
		g.Log().Warningf(ctx, "GetInviterInfo error:%s", err)
		return -1, err
	}
	if inviterInfo == nil {
		inviterInfo.UserId = userId
		inviterInfo.Stat = consts.InviteStatPass
		inviterInfo.CreateTime = time.Now().Unix()
		inviterInfo.UpdateTime = time.Now().Unix()

		shareId, err = dao.InviterInfo.Ctx(ctx).Data(inviterInfo).InsertAndGetId()
		if err != nil {
			g.Log().Warningf(ctx, "SaveInviterInfo error:%s", err)
			return -1, err
		}
		return shareId, nil
	}

	return int64(inviterInfo.Id), nil
}

// 获取短链接
func getShortUrl(ctx context.Context, userId int64, rawUrl string, source string) string {
	shortCode := utility.TransTo62(userId*1000, source)
	var shortUrls *entity.ShortUrls
	err := dao.ShortUrls.Ctx(ctx).
		Where(dao.ShortUrls.Columns().ShortCode, shortCode).
		Where(dao.ShortUrls.Columns().Source, source).Scan(&shortUrls)

	if err != nil {
		return shortUrls.RawUrl
	} else {
		shortUrls = &entity.ShortUrls{
			UserId: userId, RawUrl: rawUrl, Source: source, ShortCode: shortCode, CreateTime: time.Now().Unix(),
		}
		_, err = dao.ShortUrls.Ctx(ctx).Data().Insert()
		if err != nil {
			return ""
		}
		//shortUrl := config.ConfigS.UrlConfig.ShareDomain + "/" + shortCode
		return rawUrl
	}
}

// checkNewUserPermission 检查用户邀新权限
func checkNewUserPermission(ctx context.Context, userId int64) int {
	orderColumns := dao.OrderInfo.Columns()
	orderCount, err := dao.OrderInfo.Ctx(ctx).
		Where(orderColumns.UserId, userId).Where(orderColumns.State, 1).Count()
	if err != nil {
		return 0
	}

	if orderCount > 0 {
		return 1
	}
	//2、工会长必须有标星主播
	startLevelCount := dao.GuildAnchor.GetAnchorCountByStartLevel(ctx, userId)
	if startLevelCount > 0 {
		return 1
	}
	//3、标星主播才有邀新资格
	guildAnchorModel, _ := dao.GuildAnchor.GetGuildAnchor(ctx, userId, consts.GuildAnchorStatusOk)
	if guildAnchorModel != nil && guildAnchorModel.StarLevel > 0 {
		return 1
	}
	return 0
}

// CheckInviterRole 检查邀请者是否有分享权限,当前默认全开放
func CheckInviterRole(ctx context.Context, inviterId int64) bool {
	inviterInfo, err := dao.InviterInfo.GetInviterInfo(ctx, inviterId)
	if err != nil {
		g.Log().Errorf(ctx, "GetInviterInfo error:%s", err)
		return false
	}
	return inviterInfo != nil && inviterInfo.Stat == model.InviteStatPass
}

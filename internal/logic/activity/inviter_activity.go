package activity

import (
	"context"
	"strings"
	"time"

	"github.com/druidcaesa/ztool"
	"github.com/go-redsync/redsync/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/entity"
	"fuya-ark/internal/service"
	"fuya-ark/utility"
)

type sInviterActivity struct{}

func init() {
	service.RegisterInviterActivity(New())
}

func New() *sInviterActivity {
	return &sInviterActivity{}
}

// ShareChannels 分享
func (s *sInviterActivity) ShareChannels(ctx context.Context, in model.ShareChannelsInput) (out model.ShareChannelsOutput, err error) {
	userId := gconv.Int64(ctx.Value(consts.CtxUserId))
	//appName := gconv.String(ctx.Value(consts.CtxAppName))
	shareId, err := GetShareId(ctx, userId)
	if err != nil || shareId <= 0 {
		return model.ShareChannelsOutput{}, err
	}
	//ad融合邀新
	source := "ad_merge"
	shareUrl := strings.ReplaceAll(consts.AdjustShareUrl, "#adj_label", gconv.String(userId))
	urlStr := getShortUrl(ctx, userId, shareUrl, source)
	return model.ShareChannelsOutput{Link: urlStr}, err
}

// NewUserLoginReward 新用户绑定
func (s *sInviterActivity) NewUserLoginReward(ctx context.Context, in model.NewUserLoginRewardInput) (out model.NewUserLoginRewardOutput, err error) {
	userId := gconv.Int64(ctx.Value(consts.CtxUserId))
	//appName := gconv.String(ctx.Value(consts.CtxAppName))
	//appChannel := gconv.String(ctx.Value(consts.CtxAppChannel))
	deviceId := gconv.String(ctx.Value(consts.CtxDeviceId))
	//clientIP := gconv.String(ctx.Value(consts.CtxClientIP))
	clientMode := gconv.String(ctx.Value(consts.CtxClientMode))

	inviterId := in.InviterId
	out = model.NewUserLoginRewardOutput{}

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
		//domain.SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 9)
		out.ErrCode = -66
		return out, nil
	}

	if inviterId == userId {
		g.Log().Warningf(ctx, "chyInfo 邀请新人 自己不能邀请自己,邀请者 = %v，被邀请者 = %v", inviterId, userId)
		out.ErrCode = -65
		return out, nil
	}

	//判断设备号是否正常
	if len(deviceId) == 0 || inviterId == 0 || len(clientMode) <= 5 {
		g.Log().Warningf(ctx, "chyInfo 邀请新人 邀请设备为空/InviterId为空，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
		//domain.SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 1)
		out.ErrCode = 4
		return out, err
	}
	guildAnchor, err := dao.GuildAnchor.GetGuildAnchor(ctx, inviterId, consts.GuildAnchorStatusOk)
	startTime := ztool.DateUtils.Now().StartTimeOfDay().Time().Unix()
	invitedCount, err := dao.ShareNewUserInfo.GetInvitedUserCountByTime(ctx, inviterId, startTime)

	if guildAnchor == nil || guildAnchor.StarLevel <= 0 { //不是星级主播 邀新限制2个
		balanceInfo := dao.BalanceInfo.GetByUserId(ctx, inviterId)
		if balanceInfo != nil && balanceInfo.TotalFee < 20000000 && invitedCount >= 2 { //充值小于200k 每天限制2个
			g.Log().Warningf(ctx, "chyInfo 充值小于200k 每天限制2个，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
			//rongcloud.SendSystemByText(inviterId, fmt.Sprintf("Hari ini Anda telah mencapai batas undangan %v orang, silakan coba lagi besok.", invitedCount))
			//domain.SaveInvalidUserRecord(ctx, userId, inviterId, clientIP, deviceId, clientMode, 2)
			return model.NewUserLoginRewardOutput{ErrCode: 66}, err
		}
	}

	////判断当天邀请人数大于20
	//if invitedCount >= 20 {
	//	g.Log().Warningf(ctx, "chyInfo 邀请新人 当天邀请人数大于20，不参与邀请结算,邀请者 = %v，被邀请者 = %v", inviterId, userId)
	//	//rongcloud.SendSystemByText(inviterId, fmt.Sprintf("Hari ini Anda telah mencapai batas undangan %v orang, silakan coba lagi besok.", invitedCount))
	//	domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 3)
	//	return NewUserLoginRewardRes{ErrCode: 66}, code.InviteErr
	//}
	//if invitedCount > 15 { //发送运营号通知
	//	err := domain.ShareMaxNotice(ctx, inviterId, invitedCount)
	//	if err != nil {
	//		g.Log().Errorf(ctx, "通知运营号邀新用户 %v 邀新达到上限失败 err=%v", userId, err)
	//	}
	//}
	//// 校验新用户是否已参加过活动
	//shareNewUserInfo, _ := domain.GetShareNewUserInfo(ctx, userId)
	//isOldUser := false
	//userInfo := userDomain.GetUserInfoById(ctx, userId)
	//balance := balanceDomain.GetByUserId(ctx, userId)
	//if head.AppChannel == "fuya_google" && time.Now().Unix()-userInfo.LastHeartbeatAt > utilsCommon.OneDaySeconds*30 &&
	//	balance != nil && balance.TotalFee > 0 { //邀请时间大于30天的fuya充值用户 可以再次绑定一个邀请 邀请大于30天的充值用户
	//	isOldUser = true
	//}
	////4星性感主播邀请新用户给2000钻石
	//if guildAnchor != nil && guildAnchor.StarLevel >= common.AnchorStar4 &&
	//	utils.IsCharmAnchor(ctx, guildAnchor.Label) {
	//	isOldUser = true
	//}
	//aAGuildCount, _ := dao.AnchorLevelSingle().GetCountByGuildId(ctx, inviterId)
	////有AA类主播的工会长邀新奖励增至2000钻
	//if aAGuildCount > 0 {
	//	g.Log().Warningf(ctx, "chyInfo 有AA类主播的工会长邀新奖励增至2000钻/人  ip = %v，aAGuildCount: =%v , 被邀请者 = %v", head.ClientIP, aAGuildCount, userId)
	//	isOldUser = true
	//}
	//if shareNewUserInfo == nil && !isOldUser { //检查是否印尼、马来ip
	//	region, _ := ip2region.SearchIp(head.ClientIP)
	//	areaInfoDao := dao.AreaInfo{}
	//	if areaInfoDao.GetByTitleCnAndIsShow(ctx, region) == nil {
	//		g.Log().Warningf(ctx, "chyInfo 邀请新人 检查是否支持的地区  ip = %v，region: =%v , 被邀请者 = %v", head.ClientIP, region, userId)
	//		domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 4)
	//		return NewUserLoginRewardRes{ErrCode: 11}, code.InviteErr
	//	}
	//
	//	// 检测该设备是否已经参加过奖励
	//	shareNewUserInfoByDeviceId, _ := domain.GetShareNewUserInfoBySmId(ctx, head.DeviceId)
	//	if shareNewUserInfoByDeviceId != nil {
	//		g.Log().Warningf(ctx, "chyInfo 邀请新人 设备已经完成绑定,设备号= %v,邀请者 = %v，被邀请者 = %v", head.DeviceId, inviterId, userId)
	//		domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 5)
	//		return NewUserLoginRewardRes{ErrCode: 12}, code.InviteErr
	//	}
	//
	//	//查找相似或类似的ip地址,一个ip地址只能被成功邀请一次，不能多次邀请
	//	same := domain.SameTheIP(ctx, inviterId, head.ClientIP)
	//	if same {
	//		g.Log().Warningf(ctx, "chyInfo 邀请新人 该IP地址已经被邀请过，ip = %v，不能重复邀请,邀请者 = %v，被邀请者 = %v", head.ClientIP, inviterId, userId)
	//		domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 6)
	//		return NewUserLoginRewardRes{ErrCode: 14}, code.InviteErr
	//	}
	//
	//	//检查邀请者是否有分享权限,当前默认全开放
	//	ok = domain.CheckInviterRole(ctx, inviterId)
	//	if !ok {
	//		g.Log().Warningf(ctx, "chyInfo 邀请新人 邀请人邀请资格失效,邀请者 = %v，被邀请者 = %v", inviterId, userId)
	//		domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 7)
	//		return NewUserLoginRewardRes{ErrCode: 8}, code.InviteErr
	//	}
	//} else if shareNewUserInfo != nil {
	//	g.Log().Warningf(ctx, "chyInfo 邀请新人 用户已经完成绑定,邀请者 = %v，被邀请者 = %v", inviterId, userId)
	//	domain.SaveInvalidUserRecord(ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, 8)
	//	return NewUserLoginRewardRes{ErrCode: -10}, code.InviteErr
	//}
	//
	////记录被邀请者/邀请者信息 运营号10001和10002的分享链接替换成了30001、30002
	//if inviterId == 10001 {
	//	inviterId = 30001
	//} else if inviterId == 10002 {
	//	inviterId = 30002
	//}
	//if errCode := domain.NewUserLoad(
	//	ctx, userId, inviterId, head.ClientIP, head.DeviceId, head.ClientMode, isOldUser); errCode != code.Succeed {
	//	return nil, errCode
	//}
	//newUserRewardDiamonds := config.ConfigS.UrlConfig.NewUserRewardDiamonds
	////加邀请奖励
	//if errCode :=
	//	domain.NewUserReward(ctx, userId, inviterId, newUserRewardDiamonds, isOldUser, common.ShareByWatchLive); errCode != code.Succeed {
	//	return nil, errCode
	//}
	////A类主播邀请过来的用户发放奖励 7天--坐骑--Pink Romance 3天--头像框--Blue tears
	//isAAnchor, _ := dao.AnchorLevelSingle().IsALevelAnchor(ctx, inviterId)
	//if isAAnchor > 0 {
	//	g.Log().Warningf(ctx, "chyInfo A类主播邀请过来的用户发放奖2000钻石 ip = %v，aAGuildCount: =%v , 被邀请者 = %v", head.ClientIP, aAGuildCount, userId)
	//	isOldUser = true
	//	mallDomain.GiftFreeGood(ctx, userId, 1, 7)
	//	mallDomain.GiftFreeGood(ctx, userId, 19, 3)
	//	text := fmt.Sprintf("Pengguna terhormat %v, Selamat join di Fuya Live, anda diundang dari teman %v dan Selamat anda mendapat hadiah efek khusus dari event Fuya Live : 7 hari Naiki Pink Romance dan 3 hari Avatar Frame Blue Tears!!", userId, inviterId)
	//	rongcloud.SendOfficialSystemMsgV2(ctx, common.TitleTypeSystem, 0, 0, 0, userInfo, userInfo, nil, text)
	//}
	//
	////运营号绑定
	//userDao := &dao.UserInfoModel{}
	//inviterInfo, _ := userDao.GetUserById(ctx, inviterId)
	//userBaseDao := &dao.UserBaseInfo{UserId: userId}
	//
	//if inviterInfo != nil {
	//	if inviterInfo.Role == utilsCommon.UserServer { //运营号
	//		userBaseDao.ServerId = inviterId
	//	} else {
	//		guildInfoDao := &dao.GuildInfoModel{}
	//		guildInfo, _ := guildInfoDao.GetGuildServiceIdByUserId(ctx, inviterId, utilsCommon.GuildAnchorStatusOk)
	//		if guildInfo != nil && guildInfo.ServerId > 0 { //公会长、主播有运营号
	//			userBaseDao.ServerId = guildInfo.ServerId
	//		} else { //用户
	//			baseDao := &dao.UserBaseInfo{}
	//			baseInfo, _ := baseDao.GetUserBaseInfoByUserId(ctx, inviterId)
	//			if baseInfo != nil && baseInfo.ServerId != 0 {
	//				userBaseDao.ServerId = baseInfo.ServerId
	//			} else {
	//				userBaseDao.ServerId = 10000
	//			}
	//		}
	//	}
	//}
	//baseDomain := &userDomain.BaseUserDomain{}
	//data := baseDomain.GetUserBaseInfo(ctx, userId, nil)
	//if data == nil {
	//	err := userBaseDao.SaveUserBaseInfo(ctx, nil)
	//	if err != nil {
	//		g.Log().Errorf(ctx, "domain/user/user_base.go:35 err:%v", err)
	//		return nil, code.SqlResultFail
	//	}
	//} else {
	//	err := userBaseDao.UpdateUserBaseServer(ctx)
	//	if err != nil {
	//		g.Log().Error(ctx, "更新用户运营号失败 err = %v", err)
	//		return nil, code.SqlResultFail
	//	}
	//}
	return model.NewUserLoginRewardOutput{}, err
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
	//startLevelCount, err := dao.GuildAnchor.Ctx(ctx).GetAnchorCountByStartLevel(userId)
	//if startLevelCount > 0 {
	//	return 1
	//}
	////3、标星主播才有邀新资格
	//guildAnchorModelDao := dao.GuildAnchorModel{}
	//guildAnchorModel, _ := guildAnchorModelDao.GetGuildAnchor(ctx, userId, utilsCommon.GuildAnchorStatusOk)
	//if guildAnchorModel != nil && guildAnchorModel.StarLevel > 0 {
	//	return 1
	//}
	return 0
}

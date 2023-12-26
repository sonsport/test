package inviter

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/consts/config"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/do"
	"fuya-ark/internal/service"
)

// SaveInvalidUserRecord 保存无效用户记录
func SaveInvalidUserRecord(ctx context.Context, userId, inviterId int64, ip, SMId, deviceName string, fullReasonCode int) error {

	if dao.ShareNewUserInfo.GetCountByUser(ctx, userId, inviterId) > 0 {
		return nil
	}
	now := time.Now().Unix()
	shareNewUserInfo := do.ShareNewUserInfo{
		UserId: userId, OwnerUserId: inviterId, ShareId: inviterId, RegisterTime: now, Ip: ip, Smid: SMId,
		DeviceName: deviceName, WatchLiveRewards: 2, IsOldUser: 0, Status: -1, FullReasonCode: fullReasonCode,
		CreateTime: now, UpdateTime: now,
	}

	err := dao.ShareNewUserInfo.SaveShareNewUserInfo(ctx, shareNewUserInfo)
	if err != nil {
		g.Log().Errorf(ctx, "SaveShareNewUserInfo error:%s", err)
		return err
	}
	inviterUserInfo, err := dao.InviterInfo.GetInviterInfo(ctx, inviterId) //更新邀请人信息
	if err != nil || inviterUserInfo == nil {
		g.Log().Errorf(ctx, "dao.GetInviterInfo err:%s ", err)
		return err
	}
	dao.InviterInfo.SetCount(ctx, inviterUserInfo.Id)
	return nil
}

// ShareMaxNotice 邀新超过上限通知运营号
func ShareMaxNotice(ctx context.Context, userId int64, inviteCount int) (err error) {
	//guilDao := &dao.GuildAnchorModel{}
	//guildInfo, _ := guilDao.GetGuildInfoByUserId(ctx, userId)
	//if guildInfo != nil && guildInfo.ServerId != 0 {
	//	noticModel := &dao.ServerNotice{
	//		UserId:       userId,
	//		ServerId:     guildInfo.ServerId,
	//		NoticeType:   common.NoticeTypeSystemShare,
	//		MessageType:  common.MessageTypeSystemShareMax,
	//		SystemRemark: fmt.Sprintf("Agency Id:%v ,%v people have been invited,Maximum of 15 per day", guildInfo.UserId, inviteCount),
	//	}
	//	err = noticModel.InsertServerNotice(ctx)
	//}
	return err
}

// SameTheIP 判断ip地址是否有重复
func SameTheIP(ctx context.Context, userId int64, strIp string) bool {

	if len(strIp) == 0 || strIp == "183.157.13.206" {
		return true
	}

	count := dao.ShareNewUserInfo.SameTheIP(ctx, strIp)
	if count > 0 {
		return true
	}

	index := strings.LastIndex(strIp, ".")
	str := strIp[:(index + 1)]
	if len(str) == 0 {
		return false
	}

	//一个相似的ip地址一天只能邀请两个
	count = dao.ShareNewUserInfo.SimilarTheIP(ctx, userId, str+"%")
	if count > 2 {
		return true
	}

	////查看邀请人一天内 手机机型连续不超过5个相同的
	deviceList := dao.ShareNewUserInfo.SimilarThePhone(ctx, userId)
	theSame := ""
	if len(deviceList) > 4 {
		isSame := true
		for i, device := range deviceList {
			if i == 0 {
				theSame = device[:5]
			} else {
				if theSame != device[:5] {
					isSame = false
					break
				}
			}
		}
		if isSame {
			return true
		}
	}

	return false
}

func NewUserLoad(ctx context.Context, userId, inviterId int64, ip, SMId, deviceName string, isOldUser bool) error {
	// 判断邀新用户看播奖励是否有效
	watchLiveRewards := 1
	illegalUserCount, err := dao.LoginRecord.GetIllegalUserCount(ctx, userId)
	if err != nil {
		g.Log().Errorf(ctx, "NewUserLoad GetIllegalUserCount error:%s", err)
		return err
	}
	if illegalUserCount > 0 {
		watchLiveRewards = 2
	}
	now := time.Now().Unix()
	shareNewUserInfo := do.ShareNewUserInfo{
		UserId: userId, OwnerUserId: inviterId, ShareId: int(inviterId), RegisterTime: now, Ip: ip,
		Smid: SMId, DeviceName: deviceName, WatchLiveRewards: watchLiveRewards,
	}
	if isOldUser {
		shareNewUserInfo.IsOldUser = 1
	}
	shareNewUserInfo.IsOldUser = 0
	shareNewUserInfo.CreateTime = now
	shareNewUserInfo.UpdateTime = now

	err = dao.ShareNewUserInfo.SaveShareNewUserInfo(ctx, shareNewUserInfo)
	if err != nil {
		g.Log().Errorf(ctx, "SaveShareNewUserInfo error:%s", err)
		return err
	}

	inviterUserInfo, err := dao.InviterInfo.GetInviterInfo(ctx, inviterId) //更新邀请人信息
	if err != nil || inviterUserInfo == nil {
		g.Log().Errorf(ctx, "dao.GetInviterInfo err:%s ", err)
		return err
	}
	dao.InviterInfo.SetCount(ctx, inviterUserInfo.Id)
	return nil
}

// NewUserReward  新用户充值奖励/新用户观看直播奖励
func NewUserReward(ctx context.Context, userId, inviterId, diamonds int64, isOldUser bool, typ int) error {

	newUserInfo, err := dao.ShareNewUserInfo.GetShareNewUserInfo(ctx, userId, inviterId)
	if err != nil {
		g.Log().Errorf(ctx, "dao.GetNewUserRecordInfo err:%s ", err)
		return err
	}
	if newUserInfo == nil {
		return fmt.Errorf("用户绑定信息不存在 newUserInfo:%v ", newUserInfo)
	}
	if newUserInfo.Status == -1 || newUserInfo.FullReasonCode > 0 {
		g.Log().Debugf(ctx, "用户邀新无效 newUserInfo:%v ", newUserInfo)
		return fmt.Errorf("用户邀新无效 newUserInfo:%v ", newUserInfo)
	}
	// 判断是否新用户是否能获得看播奖励
	if typ == consts.ShareByWatchLive && newUserInfo.WatchLiveRewards == consts.InvalidRewardForWatchLive && !isOldUser {
		g.Log().Warningf(ctx, "判断是否新用户是否能获得看播奖励 userId %v  isOldUser %v", userId, isOldUser)
		return fmt.Errorf("判断是否新用户是否能获得看播奖励 userId %v  isOldUser %v", userId, isOldUser)
	}

	// 已经加过的不再重复加
	if newUserInfo.Status == consts.ShareLiveSuccess && typ == consts.ShareByWatchLive {
		g.Log().Warningf(ctx, "已经加过的不再重复加 userId %v  isOldUser %v", userId, isOldUser)
		return fmt.Errorf("已经加过的不再重复加 userId %v  isOldUser %v", userId, isOldUser)
	}

	if typ == consts.ShareByWatchLive && !isOldUser { //看播奖励
		illegalUserCount, err := dao.LoginRecord.GetIllegalUserCount(ctx, userId)
		if err != nil {
			g.Log().Errorf(ctx, "NewUserLoad GetIllegalUserCount error:%s", err)
			return err
		}
		if illegalUserCount > 0 { // 判断邀新用户看播奖励是否有效
			dao.ShareNewUserInfo.UpdateWatchLiveRewards(ctx, userId, consts.InvalidRewardForWatchLive)
			return fmt.Errorf("判断邀新用户看播奖励是否有效")
		}
	}
	// 邀请人信息
	inviterUserInfo, err := dao.InviterInfo.GetInviterInfo(ctx, newUserInfo.OwnerUserId)
	if err != nil {
		g.Log().Errorf(ctx, "dao.GetInviterInfo err:%s ", err)
		return err
	}
	if inviterUserInfo == nil || inviterUserInfo.Stat != model.InviteStatPass {
		g.Log().Warning(ctx, "分享用户信息不存在/或者邀请人资格失效")
		return fmt.Errorf("分享用户信息不存在/或者邀请人资格失效")
	}
	curDate := time.Now().Format("20060102")
	// 分享充值奖励计算
	if typ == consts.ShareByRecharge {
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 被邀请人一个月内充值了，邀请人获得10%奖励（被邀请人充值钻石*10%），次数不限
			nowTime := time.Now().Unix()
			if nowTime-newUserInfo.CreateTime <= consts.TimeOneMonth {
				ownerExtraRewardDiamond := diamonds / 10 // 邀请者获得百分之十奖励
				deptType := consts.RechargeRewards
				_, err := service.Balances().Recharge(
					ctx, inviterUserInfo.UserId, newUserInfo.UserId, 0, ownerExtraRewardDiamond, newUserInfo.ShareId, uint(deptType), tx,
				)
				if err != nil {
					g.Log().Warning(ctx, "充值奖励dao.Recharge err:%s ", err)
					return err
				}
				g.Redis().IncrBy(ctx, consts.StatisticInviteAward+curDate, ownerExtraRewardDiamond)
				g.Redis().Expire(ctx, consts.StatisticInviteAward+curDate, consts.Board3Day)
			}
			//成功充值且还没有收到第一次观看直播奖励（第一次登录开始的24小时内），邀请人（主播或用户）获得1000钻石，被邀请人获得200钻石；
			if newUserInfo.Status != consts.ShareLiveSuccess {
				// 24小时内的有效
				if nowTime-newUserInfo.CreateTime <= consts.TimeDayHour { //在这里判断
					err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
						if err = UpdateReward(ctx, inviterUserInfo.UserId, newUserInfo.UserId, int64(newUserInfo.ShareId), false, tx); err != nil {
							return errors.New("UpdateReward")
						}
						return nil
					})
				}
			}
			return nil
		})
	} else if typ == consts.ShareByWatchLive {
		//分享观看直播奖励计算
		err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			return UpdateReward(ctx, inviterUserInfo.UserId, newUserInfo.UserId, int64(newUserInfo.ShareId), isOldUser, tx)
		})
	}
	if err != nil {
		g.Log().Errorf(ctx, "NewUserPaymentReward  新用户充值奖励 异常 err:%v", err)
		return err
	}
	return nil
}

func UpdateReward(ctx context.Context, inviterUserId int64, newUserId int64, shareId int64, isOldUser bool, tx gdb.TX) error {
	inviterDiamonds := config.ConfigS.UrlConfig.OwnerRewardDiamonds
	invitedDiamonds := config.ConfigS.UrlConfig.NewUserRewardDiamonds
	//userInfo, _ := dao.UserInfo.GetUserById(ctx, newUserId)
	//inviterUserInfo, _ := dao.UserInfo.GetUserById(ctx, inviterUserId)
	if isOldUser { //付费老用户召回  发送优惠券
		inviterDiamonds = 2000
		//todo
		//couponInfoDomain := couponDomain.NewCouponInfoDomain()
		//couponInfoDomain.SaveCouponInfo(ctx, newUserId, 1, shareId, int64(10000000), int64(50000000), 10)
		//text := "Selamat datang di rumah, kami berikan kepada Anda kupon hadiah berlian, yang bisa langsung digunakan saat Anda melakukan pengisian ulang."
		//rongcloud.SendOfficialSystemMsgV2(ctx, common.TitleTypeShare, 0, common.LastTitleTypeInviter, 0, userInfo, userInfo, inviterUserInfo, text)
	}
	curDate := time.Now().Format("20060102")
	//邀请人的奖励
	_, errCode := service.Balances().Recharge(ctx, inviterUserId, newUserId, 0, inviterDiamonds, shareId, consts.InviteRewards, tx)
	if errCode != nil {
		return errors.New("dao.UpdateShareOnwerUserInfoRechargeReward")
	}
	//被邀请信息更新
	if err := dao.ShareNewUserInfo.SetInvitedInfoStatus(ctx, newUserId, consts.ShareLiveSuccess, tx); err != nil {
		return errors.New("dao.UpdateShareNewUserInfoStatus")
	}
	//被邀请的奖励
	_, errCode = service.Balances().Recharge(ctx, newUserId, inviterUserId, 0, invitedDiamonds, shareId, consts.InviteRewards, tx)
	if errCode != nil {
		return errors.New("dao.UpdateShareOnwerUserInfoRechargeReward")
	}
	//发送消息通知邀请成功
	//邀请人发送邀新成功消息
	//text := fmt.Sprintf(utils.TranslateByKey(ctx, "invitationSuccessful"), inviterDiamonds)
	//rongcloud.SendOfficialSystemMsgV2(ctx, common.TitleTypeShare, common.SubTitleTypeDiamondReward, common.LastTitleTypeInvited, inviterDiamonds, inviterUserInfo, userInfo, inviterUserInfo, text)
	////被邀请人发送被邀请成功消息
	//text = fmt.Sprintf(utils.TranslateByKey(ctx, "invitationSuccessful"), invitedDiamonds)
	//rongcloud.SendOfficialSystemMsgV2(ctx, common.TitleTypeShare, common.SubTitleTypeDiamondReward, common.LastTitleTypeInviter, invitedDiamonds, userInfo, userInfo, inviterUserInfo, text)

	g.Redis().IncrBy(ctx, consts.StatisticInviteAward+curDate, inviterDiamonds)
	g.Redis().Expire(ctx, consts.StatisticInviteAward+curDate, consts.Board3Day)
	return nil
}

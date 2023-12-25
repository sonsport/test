package inviter

import (
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"fuya-ark/internal/dao"
	"fuya-ark/internal/model/do"
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

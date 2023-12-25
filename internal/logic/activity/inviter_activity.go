package activity

import (
	"context"
	"strings"
	"time"

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

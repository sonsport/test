// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"fuya-ark/internal/model"
)

type (
	IInviter interface {
		// ShareChannels 分享
		ShareChannels(ctx context.Context, in model.ShareChannelsReq) (out model.ShareChannelsRes, err error)
		// NewUserLoginReward 新用户绑定
		NewUserLoginReward(ctx context.Context, in model.NewUserLoginRewardReq) (out model.NewUserLoginRewardRes, err error)
	}
)

var (
	localInviter IInviter
)

func Inviter() IInviter {
	if localInviter == nil {
		panic("implement not found for interface IInviter, forgot register?")
	}
	return localInviter
}

func RegisterInviter(i IInviter) {
	localInviter = i
}

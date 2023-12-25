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
	IInviterActivity interface {
		// ShareChannels 分享
		ShareChannels(ctx context.Context, in model.ShareChannelsInput) (out model.ShareChannelsOutput, err error)
		// NewUserLoginReward 新用户绑定
		NewUserLoginReward(ctx context.Context, in model.NewUserLoginRewardInput) (out model.NewUserLoginRewardOutput, err error)
	}
)

var (
	localInviterActivity IInviterActivity
)

func InviterActivity() IInviterActivity {
	if localInviterActivity == nil {
		panic("implement not found for interface IInviterActivity, forgot register?")
	}
	return localInviterActivity
}

func RegisterInviterActivity(i IInviterActivity) {
	localInviterActivity = i
}

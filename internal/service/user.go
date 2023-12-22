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
	IUser interface {
		// Register 注册
		Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error)
		// UpdatePassword 修改密码
		UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error)
		// UserDetail 获取用户详情
		UserDetail(ctx context.Context, in model.UserDetailInput) (out model.UserDetailOutput, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}

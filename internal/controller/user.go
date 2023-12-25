package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/model"
	"fuya-ark/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *model.RegisterInput) (res *model.RegisterOutput, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &model.RegisterOutput{Id: out.Id}, nil
}

func (*cUser) UpdatePassword(ctx context.Context, req *model.UpdatePasswordInput) (res *model.UpdatePasswordOutput, err error) {
	data := model.UpdatePasswordInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UpdatePassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &model.UpdatePasswordOutput{Id: out.Id}, nil
}

func (*cUser) UserDetail(ctx context.Context, req *model.UserDetailInput) (res *model.UserDetailOutput, err error) {
	data := model.UserDetailInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().UserDetail(ctx, data)
	if err != nil {
		return nil, err
	}
	err = gconv.Struct(out, &res)
	return res, err
}

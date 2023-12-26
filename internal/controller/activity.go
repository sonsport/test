package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/model"
	"fuya-ark/internal/service"
	"fuya-ark/utility"
)

var Activity = cActivity{}

type cActivity struct{}

func (c *cActivity) ShareChannels(ctx context.Context, r RequestParam) (resp interface{}, gCode gcode.Code) {
	data := model.ShareChannelsReq{}
	err := gconv.Struct(r, &data)
	if err != nil {
		return nil, utility.ParamError
	}
	out, err := service.InviterActivity().ShareChannels(ctx, data)
	if err != nil {
		return nil, utility.Internal
	}
	return &model.ShareChannelsRes{Link: out.Link}, utility.Succeed
}

func (c *cActivity) NewUserLoginReward(ctx context.Context, r RequestParam) (resp interface{}, gCode gcode.Code) {
	data := model.NewUserLoginRewardReq{}
	err := gconv.Struct(r, &data)
	if err != nil {
		return nil, utility.ParamError
	}
	out, err := service.InviterActivity().NewUserLoginReward(ctx, data)
	if err != nil {
		return nil, utility.Internal
	}
	return out, utility.Succeed
}

package controller

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "fuya-ark/api/v1"
	"fuya-ark/internal/model"
	"fuya-ark/internal/service"
)

var Activity = cActivity{}

type cActivity struct{}

func (c *cActivity) ShareChannels(ctx context.Context, req *v1.ShareChannelsReq) (res *v1.ShareChannelsRes, err error) {
	data := model.ShareChannelsInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.InviterActivity().ShareChannels(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.ShareChannelsRes{Link: out.Link}, nil
}

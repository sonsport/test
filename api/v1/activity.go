package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShareChannelsReq 获取分享渠道及分享链接
type ShareChannelsReq struct {
	g.Meta      `path:"/share/channels" method:"get" tag:"api" summary:"获取分享渠道及分享链接"`
	UserId      int64  `json:"userId"  v:"userId"   description:"userId"`
	Source      string `json:"source" validate:"required" description:"区分不同分享方式"`
	CustomQuery string `json:"customQuery" description:"自定义请求字符串 json"`
}

type ShareChannelsRes struct {
	Link string `json:"link"  description:"分享链接"`
}

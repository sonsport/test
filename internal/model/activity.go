package model

// ShareChannelsInput 获取分享渠道及分享链接
type ShareChannelsInput struct {
	UserId      int64  `json:"userId"  v:"userId"   description:"userId"`
	Source      string `json:"source" validate:"required" description:"区分不同分享方式"`
	CustomQuery string `json:"custom_query" description:"自定义请求字符串 json"`
}

type ShareChannelsOutput struct {
	Link string `json:"link"  description:"分享链接"`
}

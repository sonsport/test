package model

// ShareChannelsReq 获取分享渠道及分享链接
type ShareChannelsReq struct {
	UserId      int64  `json:"userId"  v:"userId"   description:"userId"`
	Source      string `json:"source" validate:"required" description:"区分不同分享方式"`
	CustomQuery string `json:"custom_query" description:"自定义请求字符串 json"`
}

type ShareChannelsRes struct {
	Link string `json:"link"  description:"分享链接"`
}

// NewUserLoginRewardReq 分享登录绑定
type NewUserLoginRewardReq struct {
	InviterId int64 `json:"inviterId"  description:"邀请用户id"`
}

type NewUserLoginRewardRes struct {
	ErrCode int `json:"errCode"`
}

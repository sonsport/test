// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BannerInfo is the golang structure for table banner_info.
type BannerInfo struct {
	Bid        int64  `json:"bid"        description:"bid"`
	Ranking    int    `json:"ranking"    description:"排序"`
	Title      string `json:"title"      description:"标题"`
	ClientOs   int    `json:"clientOs"   description:"0 全部；1 android；2 iOS"`
	AppName    string `json:"appName"    description:"app名称"`
	Cover      string `json:"cover"      description:"封面URL"`
	Target     int    `json:"target"     description:"跳转类型，0 网页，1 app内页"`
	Link       string `json:"link"       description:"跳转URL"`
	ShowType   int    `json:"showType"   description:"展示类型，0 全部，1 未充值用户 2 充值用户 3 主播"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
	State      int    `json:"state"      description:"状态1有效、2失效"`
}

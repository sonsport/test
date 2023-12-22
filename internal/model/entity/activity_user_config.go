// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityUserConfig is the golang structure for table activity_user_config.
type ActivityUserConfig struct {
	Id         int    `json:"id"         description:""`
	AcId       int    `json:"acId"       description:"活动Id activity_config 自增id"`
	ShowUser   string `json:"showUser"   description:"展示对象：ALL 全部、UNPAID未充值用户 、PAID已经充值用户 、HOST主播公会 后续扩展"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

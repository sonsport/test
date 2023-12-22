// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityChannelConfig is the golang structure for table activity_channel_config.
type ActivityChannelConfig struct {
	Id          int    `json:"id"          description:""`
	AcId        int    `json:"acId"        description:"活动Id activity_config 自增id"`
	ShowChannel string `json:"showChannel" description:"渠道：ALL全部渠道、指定渠道"`
	CreateTime  int64  `json:"createTime"  description:""`
	UpdateTime  int64  `json:"updateTime"  description:""`
}

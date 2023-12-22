// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityChannelConfig is the golang structure of table activity_channel_config for DAO operations like Where/Data.
type ActivityChannelConfig struct {
	g.Meta      `orm:"table:activity_channel_config, do:true"`
	Id          interface{} //
	AcId        interface{} // 活动Id activity_config 自增id
	ShowChannel interface{} // 渠道：ALL全部渠道、指定渠道
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

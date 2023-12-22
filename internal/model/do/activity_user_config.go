// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityUserConfig is the golang structure of table activity_user_config for DAO operations like Where/Data.
type ActivityUserConfig struct {
	g.Meta     `orm:"table:activity_user_config, do:true"`
	Id         interface{} //
	AcId       interface{} // 活动Id activity_config 自增id
	ShowUser   interface{} // 展示对象：ALL 全部、UNPAID未充值用户 、PAID已经充值用户 、HOST主播公会 后续扩展
	CreateTime interface{} //
	UpdateTime interface{} //
}

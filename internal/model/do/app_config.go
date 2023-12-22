// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AppConfig is the golang structure of table app_config for DAO operations like Where/Data.
type AppConfig struct {
	g.Meta     `orm:"table:app_config, do:true"`
	Id         interface{} // 自增Id
	System     interface{} // 所属系统；0所有,1安卓 2ios
	Version    interface{} // 版本
	Channel    interface{} // 渠道
	Key        interface{} // key
	Value      interface{} // 值
	Remark     interface{} // 备注
	CreateTime interface{} //
	UpdateTime interface{} //
}

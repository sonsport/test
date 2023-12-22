// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdjustMessage is the golang structure of table adjust_message for DAO operations like Where/Data.
type AdjustMessage struct {
	g.Meta      `orm:"table:adjust_message, do:true"`
	Id          interface{} //
	InstallTime interface{} // install_time
	UserId      interface{} // 用户Id
	UserName    interface{} //
	OrderId     interface{} // 用户Id
	EventToken  interface{} //
	GaId        interface{} // 订单id
	DeviceId    interface{} // 设备id
	AppName     interface{} //
	AppId       interface{} //
	Tracker     interface{} //
	Network     interface{} //
	AdjustId    interface{} //
	Currency    interface{} //
	Revenue     interface{} //
	CreateTime  interface{} //
}

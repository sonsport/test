// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LoginRecord is the golang structure of table login_record for DAO operations like Where/Data.
type LoginRecord struct {
	g.Meta     `orm:"table:login_record, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户id
	AppName    interface{} // app名称
	AppChannel interface{} // app渠道
	AppVersion interface{} // app版本
	DeviceId   interface{} // 设备号
	Ip         interface{} // 登陆ip
	Language   interface{} // 语言
	CreateTime interface{} //
	UpdateTime interface{} //
	AppToken   interface{} // app登录token
	LoginType  interface{} // 1 gg 2 fb 3 apple
	AppPhone   interface{} // app手机型号
}

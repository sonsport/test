// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BanUserInfo is the golang structure of table ban_user_info for DAO operations like Where/Data.
type BanUserInfo struct {
	g.Meta     `orm:"table:ban_user_info, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 封禁用户id
	AdminId    interface{} // 管理员id
	LinkId     interface{} // 关联警告表主键id
	Type       interface{} // 封禁类型，1禁发私信，2禁发公屏聊天，3禁开播，4禁止登录，5封禁设备
	BanTime    interface{} // 封禁时间
	DeviceId   interface{} // 封禁设备
	BanReason  interface{} // 封禁原因
	CreateTime interface{} //
	UpdateTime interface{} //
}

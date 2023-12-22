// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomAdmin is the golang structure of table room_admin for DAO operations like Where/Data.
type RoomAdmin struct {
	g.Meta     `orm:"table:room_admin, do:true"`
	Id         interface{} // ID
	UserId     interface{} // 用户ID 存储主播的id
	ManagerId  interface{} // 管理员ID 主播指定的管理员
	Valid      interface{} // 是否有效，0无效 1有效
	CreateTime interface{} //
	UpdateTime interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomRemoveList is the golang structure of table room_remove_list for DAO operations like Where/Data.
type RoomRemoveList struct {
	g.Meta     `orm:"table:room_remove_list, do:true"`
	Id         interface{} // ID
	UserId     interface{} // UID
	OperatorId interface{} // 操作者ID
	RoomId     interface{} // 房间ID
	EndTime    interface{} // 移出结束时间,为0就是永久拉黑
	CreateTime interface{} //
	UpdateTime interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomUnlockRecord is the golang structure of table room_unlock_record for DAO operations like Where/Data.
type RoomUnlockRecord struct {
	g.Meta       `orm:"table:room_unlock_record, do:true"`
	Id           interface{} //
	RoomId       interface{} //
	UserId       interface{} // 用户id
	AnchorUserId interface{} // 主播id
	IsUnlock     interface{} // 是否解锁
	CreateTime   interface{} //
	UpdateTime   interface{} //
	UnlockPrice  interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLiveCall is the golang structure of table room_live_call for DAO operations like Where/Data.
type RoomLiveCall struct {
	g.Meta       `orm:"table:room_live_call, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户id
	AnchorUserId interface{} // 主播用户id
	RoomId       interface{} // 房间id
	CallStates   interface{} // 连麦状态 0 排队中 1连麦中 2 已结束
	StreamId     interface{} // 连麦流id
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomMultipleUser is the golang structure of table room_multiple_user for DAO operations like Where/Data.
type RoomMultipleUser struct {
	g.Meta           `orm:"table:room_multiple_user, do:true"`
	Id               interface{} //
	UserId           interface{} // 用户id
	Nickname         interface{} // 昵称
	Portrait         interface{} // 头像
	RoomId           interface{} // 房间id
	AnchorUserId     interface{} // 主播用户id
	LiveControl      interface{} // 直播控制 1 正常 2 禁视频
	SoundControl     interface{} // 直播控制 1 正常 2 静音
	LiveStates       interface{} // 多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略
	StreamId         interface{} // 流id 房主的为流混流拉流id
	MultipleUserRole interface{} // 多人房身份 1 多人房主播 2 上麦者
	MainStatus       interface{} // 主麦状态
	CreateTime       interface{} //
	UpdateTime       interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomUserRecord is the golang structure of table room_user_record for DAO operations like Where/Data.
type RoomUserRecord struct {
	g.Meta     `orm:"table:room_user_record, do:true"`
	Id         interface{} // 自增长ID
	UserId     interface{} // 用户ID
	RoomId     interface{} // 房间ID
	ExitType   interface{} // 用户退出房间原因，0 自己退出；1 房间管理员；2 主播提出房间 3 超级管理踢出房间；4 主播拉黑；5 异常退出房间
	Extend     interface{} // 额外信息
	IsChat     interface{} // 是否有聊天
	IsWaitLong interface{} // 是否待满5分钟
	Diamonds   interface{} // 在此房间消费的钻石流水
	BeginTime  interface{} // 进入房间时间
	EndTime    interface{} // 退出房间时间
	CreateTime interface{} //
	UpdateTime interface{} //
}

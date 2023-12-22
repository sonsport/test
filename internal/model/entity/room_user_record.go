// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomUserRecord is the golang structure for table room_user_record.
type RoomUserRecord struct {
	Id         uint64 `json:"id"         description:"自增长ID"`
	UserId     int64  `json:"userId"     description:"用户ID"`
	RoomId     string `json:"roomId"     description:"房间ID"`
	ExitType   int    `json:"exitType"   description:"用户退出房间原因，0 自己退出；1 房间管理员；2 主播提出房间 3 超级管理踢出房间；4 主播拉黑；5 异常退出房间"`
	Extend     string `json:"extend"     description:"额外信息"`
	IsChat     int    `json:"isChat"     description:"是否有聊天"`
	IsWaitLong int    `json:"isWaitLong" description:"是否待满5分钟"`
	Diamonds   int64  `json:"diamonds"   description:"在此房间消费的钻石流水"`
	BeginTime  int64  `json:"beginTime"  description:"进入房间时间"`
	EndTime    int64  `json:"endTime"    description:"退出房间时间"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

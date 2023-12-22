// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomAdmin is the golang structure for table room_admin.
type RoomAdmin struct {
	Id         uint64 `json:"id"         description:"ID"`
	UserId     int64  `json:"userId"     description:"用户ID 存储主播的id"`
	ManagerId  int64  `json:"managerId"  description:"管理员ID 主播指定的管理员"`
	Valid      int    `json:"valid"      description:"是否有效，0无效 1有效"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

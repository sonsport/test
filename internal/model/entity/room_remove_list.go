// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomRemoveList is the golang structure for table room_remove_list.
type RoomRemoveList struct {
	Id         uint64 `json:"id"         description:"ID"`
	UserId     int64  `json:"userId"     description:"UID"`
	OperatorId int64  `json:"operatorId" description:"操作者ID"`
	RoomId     int64  `json:"roomId"     description:"房间ID"`
	EndTime    int64  `json:"endTime"    description:"移出结束时间,为0就是永久拉黑"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

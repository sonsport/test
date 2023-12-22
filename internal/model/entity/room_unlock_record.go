// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomUnlockRecord is the golang structure for table room_unlock_record.
type RoomUnlockRecord struct {
	Id           int    `json:"id"           description:""`
	RoomId       string `json:"roomId"       description:""`
	UserId       int64  `json:"userId"       description:"用户id"`
	AnchorUserId int64  `json:"anchorUserId" description:"主播id"`
	IsUnlock     int    `json:"isUnlock"     description:"是否解锁"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
	UnlockPrice  int    `json:"unlockPrice"  description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomLiveCall is the golang structure for table room_live_call.
type RoomLiveCall struct {
	Id           int64  `json:"id"           description:""`
	UserId       int64  `json:"userId"       description:"用户id"`
	AnchorUserId int64  `json:"anchorUserId" description:"主播用户id"`
	RoomId       string `json:"roomId"       description:"房间id"`
	CallStates   int    `json:"callStates"   description:"连麦状态 0 排队中 1连麦中 2 已结束"`
	StreamId     string `json:"streamId"     description:"连麦流id"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

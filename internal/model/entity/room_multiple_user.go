// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomMultipleUser is the golang structure for table room_multiple_user.
type RoomMultipleUser struct {
	Id               int64  `json:"id"               description:""`
	UserId           int64  `json:"userId"           description:"用户id"`
	Nickname         string `json:"nickname"         description:"昵称"`
	Portrait         string `json:"portrait"         description:"头像"`
	RoomId           string `json:"roomId"           description:"房间id"`
	AnchorUserId     int64  `json:"anchorUserId"     description:"主播用户id"`
	LiveControl      int    `json:"liveControl"      description:"直播控制 1 正常 2 禁视频"`
	SoundControl     int    `json:"soundControl"     description:"直播控制 1 正常 2 静音"`
	LiveStates       int    `json:"liveStates"       description:"多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略"`
	StreamId         string `json:"streamId"         description:"流id 房主的为流混流拉流id"`
	MultipleUserRole int    `json:"multipleUserRole" description:"多人房身份 1 多人房主播 2 上麦者"`
	MainStatus       int    `json:"mainStatus"       description:"主麦状态"`
	CreateTime       int64  `json:"createTime"       description:""`
	UpdateTime       int64  `json:"updateTime"       description:""`
}

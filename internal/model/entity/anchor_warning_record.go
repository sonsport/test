// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorWarningRecord is the golang structure for table anchor_warning_record.
type AnchorWarningRecord struct {
	Id          int64  `json:"id"          description:"自增id"`
	Icon        string `json:"icon"        description:"警告图片证据"`
	Appeal      string `json:"appeal"      description:"申诉"`
	State       int    `json:"state"       description:"0待申诉，1申诉中，2成功，3失败"`
	UserId      int64  `json:"userId"      description:"主播id"`
	RoomId      string `json:"roomId"      description:"直播间id"`
	WarnType    int    `json:"warnType"    description:"警告类型"`
	Reason      string `json:"reason"      description:"警告原因"`
	WarnBigType int    `json:"warnBigType" description:""`
	AdminId     int64  `json:"adminId"     description:"提醒大类"`
	CreateTime  int64  `json:"createTime"  description:""`
	UpdateTime  int64  `json:"updateTime"  description:""`
}

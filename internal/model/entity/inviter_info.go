// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// InviterInfo is the golang structure for table inviter_info.
type InviterInfo struct {
	Id              uint  `json:"id"              description:""`
	UserId          int64 `json:"userId"          description:""`
	Stat            int   `json:"stat"            description:""`
	InviteUserCount int   `json:"inviteUserCount" description:""`
	CreateTime      int64 `json:"createTime"      description:""`
	UpdateTime      int64 `json:"updateTime"      description:""`
}

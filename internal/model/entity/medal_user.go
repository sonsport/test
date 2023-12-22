// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MedalUser is the golang structure for table medal_user.
type MedalUser struct {
	Id          int64 `json:"id"          description:""`
	UserId      int64 `json:"userId"      description:"用户id"`
	MedalId     int64 `json:"medalId"     description:"勋章id"`
	AccessCount int   `json:"accessCount" description:"获得次数"`
	CreateTime  int64 `json:"createTime"  description:""`
	UpdateTime  int64 `json:"updateTime"  description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SafeInfo is the golang structure for table safe_info.
type SafeInfo struct {
	Id         int64  `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:"用户Id"`
	Content    string `json:"content"    description:"内容"`
	SafeType   int    `json:"safeType"   description:"类型，1头像，2昵称，3签名，后续扩展"`
	State      int    `json:"state"      description:"状态，0待审核，1审核通过，2审核不通过"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

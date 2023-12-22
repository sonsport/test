// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserCheckFirst is the golang structure for table user_check_first.
type UserCheckFirst struct {
	Id         int64 `json:"id"         description:""`
	UserId     int64 `json:"userId"     description:"用户Id"`
	CheckType  int   `json:"checkType"  description:"检查类型"`
	CreateTime int64 `json:"createTime" description:"创建时间"`
	UpdateTime int64 `json:"updateTime" description:"最后更新时间"`
}

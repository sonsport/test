// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BlackList is the golang structure for table black_list.
type BlackList struct {
	Id         uint64 `json:"id"         description:"ID"`
	UserId     int64  `json:"userId"     description:"用户ID"`
	OperatorId int64  `json:"operatorId" description:"操作者ID"`
	BlockedId  int64  `json:"blockedId"  description:"被拉黑的用户ID"`
	EndTime    int64  `json:"endTime"    description:"拉黑结束时间,为0就是永久拉黑"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

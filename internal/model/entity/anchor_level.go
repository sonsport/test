// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorLevel is the golang structure for table anchor_level.
type AnchorLevel struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	UserId     int64  `json:"userId"     description:"用户ID"`
	Level      string `json:"level"      description:"等级"`
	Remark     string `json:"remark"     description:"等级"`
	State      int    `json:"state"      description:"状态1有效，2无效"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

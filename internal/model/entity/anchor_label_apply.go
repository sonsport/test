// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorLabelApply is the golang structure for table anchor_label_apply.
type AnchorLabelApply struct {
	Id         int    `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:""`
	Label      string `json:"label"      description:"标签"`
	States     int    `json:"states"     description:"状态 1 待审核 2 已通过 3 已驳回"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

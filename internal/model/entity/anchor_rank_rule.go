// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorRankRule is the golang structure for table anchor_rank_rule.
type AnchorRankRule struct {
	Id         uint  `json:"id"         description:"主键ID"`
	Level      int   `json:"level"      description:"等级从0开始, 0为等级1"`
	Min        int64 `json:"min"        description:"最小值"`
	Max        int64 `json:"max"        description:"最大值"`
	CreateTime int64 `json:"createTime" description:""`
	UpdateTime int64 `json:"updateTime" description:""`
}

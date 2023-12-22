// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BalanceRecord is the golang structure for table balance_record.
type BalanceRecord struct {
	Id            uint64 `json:"id"            description:""`
	SourceUserId  uint64 `json:"sourceUserId"  description:"用户Id"`
	TargetUserId  uint64 `json:"targetUserId"  description:"主播Id"`
	DepletionType uint   `json:"depletionType" description:"消费类型"`
	Diamonds      uint64 `json:"diamonds"      description:"钻石变动数额"`
	BeforeChange  uint64 `json:"beforeChange"  description:"变更之前余额"`
	AfterChange   uint64 `json:"afterChange"   description:"变更之后余额"`
	LinkId        uint64 `json:"linkId"        description:"引起余额变更的id"`
	Type          uint   `json:"type"          description:"消费类型 1支出 2收入"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

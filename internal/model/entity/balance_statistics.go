// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BalanceStatistics is the golang structure for table balance_statistics.
type BalanceStatistics struct {
	Id            uint64 `json:"id"            description:""`
	DepletionType uint   `json:"depletionType" description:"消费类型"`
	Diamonds      uint64 `json:"diamonds"      description:"钻石变动数额"`
	BeginTime     int64  `json:"beginTime"     description:"开始时间"`
	EndTime       int64  `json:"endTime"       description:"结束时间"`
	TimeLevel     int    `json:"timeLevel"     description:"时间维度：1.每小时，2.每天"`
	BigType       uint   `json:"bigType"       description:"消费类型 1支出 2收入"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

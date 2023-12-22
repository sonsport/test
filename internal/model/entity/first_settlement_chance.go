// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FirstSettlementChance is the golang structure for table first_settlement_chance.
type FirstSettlementChance struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	Week       int    `json:"week"       description:"周，一年中的第几周 例如:202230 2022年的第30周"`
	UserId     int64  `json:"userId"     description:"用户ID"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

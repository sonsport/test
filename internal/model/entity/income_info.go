// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// IncomeInfo is the golang structure for table income_info.
type IncomeInfo struct {
	Id                         uint64 `json:"id"                         description:"id"`
	UserId                     int64  `json:"userId"                     description:"用户id"`
	TotalIncome                int64  `json:"totalIncome"                description:"总收益 单位分"`
	BalancesIncome             int64  `json:"balancesIncome"             description:"收益余额 单位分"`
	TotalCommissionDiamonds    int64  `json:"totalCommissionDiamonds"    description:""`
	BalancesCommissionDiamonds int64  `json:"balancesCommissionDiamonds" description:""`
	CreateTime                 int64  `json:"createTime"                 description:"总佣金收入 扩大10000"`
	UpdateTime                 int64  `json:"updateTime"                 description:"佣金收入 扩大10000"`
}

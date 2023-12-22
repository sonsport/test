// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BalanceInfo is the golang structure for table balance_info.
type BalanceInfo struct {
	UserId              int64  `json:"userId"              description:"用户Id"`
	TotalFee            uint64 `json:"totalFee"            description:"总花费、只有付费购买才累计到这里 印尼盾 单位：分"`
	TotalPayoutDiamonds uint64 `json:"totalPayoutDiamonds" description:"总流水，凡是用户花费的钻石都记录到这里 累计"`
	TotalPayDiamonds    uint64 `json:"totalPayDiamonds"    description:"总购买钻石 只累计购买的"`
	RemainDiamonds      uint64 `json:"remainDiamonds"      description:"剩余钻石数"`
	CreateTime          int64  `json:"createTime"          description:""`
	UpdateTime          int64  `json:"updateTime"          description:""`
}

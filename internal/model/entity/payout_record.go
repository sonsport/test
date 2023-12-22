// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PayoutRecord is the golang structure for table payout_record.
type PayoutRecord struct {
	Id          int64  `json:"id"          description:""`
	PayoutId    int64  `json:"payoutId"    description:"代付id"`
	OrderNo     string `json:"orderNo"     description:"系统订单号"`
	TradeNo     string `json:"tradeNo"     description:"交易订单号"`
	Status      string `json:"status"      description:"状态"`
	FailureCode string `json:"failureCode" description:"失败状态"`
	CreateTime  int64  `json:"createTime"  description:"创建时间"`
	UpdateTime  int64  `json:"updateTime"  description:"修改时间"`
}

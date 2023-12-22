// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TradesInfo is the golang structure for table trades_info.
type TradesInfo struct {
	OrderId    uint64 `json:"orderId"    description:"订单Id"`
	Tid        string `json:"tid"        description:"票据Id"`
	PayType    int    `json:"payType"    description:"支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay"`
	TotalFee   uint64 `json:"totalFee"   description:"支付费用，三方支付通知有，没有则填充订单的对应货币金额 单位分"`
	Currency   string `json:"currency"   description:"实际支付货币，三方通知有说明，比如：IDR，没有则填充对应货币单位"`
	State      int    `json:"state"      description:"1, 交易成功0, 未知的交易状态 2,   退款 3,  续订  4, 等待支付"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

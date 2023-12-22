// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OrderStatistics is the golang structure for table order_statistics.
type OrderStatistics struct {
	Id                int    `json:"id"                description:""`
	BeginTime         int64  `json:"beginTime"         description:"开始时间"`
	EndTime           int64  `json:"endTime"           description:"结束时间"`
	TimeLevel         int    `json:"timeLevel"         description:"时间维度：1.每小时，2.每天"`
	AppChannel        string `json:"appChannel"        description:"app渠道"`
	PayType           int    `json:"payType"           description:"支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay"`
	ChannelCode       int    `json:"channelCode"       description:"支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定"`
	TotalFee          uint64 `json:"totalFee"          description:"商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段"`
	AllOrderCount     int64  `json:"allOrderCount"     description:"总订单数"`
	SuccessOrderCount int64  `json:"successOrderCount" description:"付费订单数"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OrderStatistics is the golang structure of table order_statistics for DAO operations like Where/Data.
type OrderStatistics struct {
	g.Meta            `orm:"table:order_statistics, do:true"`
	Id                interface{} //
	BeginTime         interface{} // 开始时间
	EndTime           interface{} // 结束时间
	TimeLevel         interface{} // 时间维度：1.每小时，2.每天
	AppChannel        interface{} // app渠道
	PayType           interface{} // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	ChannelCode       interface{} // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	TotalFee          interface{} // 商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段
	AllOrderCount     interface{} // 总订单数
	SuccessOrderCount interface{} // 付费订单数
}

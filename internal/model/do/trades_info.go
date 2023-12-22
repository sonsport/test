// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TradesInfo is the golang structure of table trades_info for DAO operations like Where/Data.
type TradesInfo struct {
	g.Meta     `orm:"table:trades_info, do:true"`
	OrderId    interface{} // 订单Id
	Tid        interface{} // 票据Id
	PayType    interface{} // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	TotalFee   interface{} // 支付费用，三方支付通知有，没有则填充订单的对应货币金额 单位分
	Currency   interface{} // 实际支付货币，三方通知有说明，比如：IDR，没有则填充对应货币单位
	State      interface{} // 1, 交易成功0, 未知的交易状态 2,   退款 3,  续订  4, 等待支付
	CreateTime interface{} //
	UpdateTime interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PaymentInfo is the golang structure of table payment_info for DAO operations like Where/Data.
type PaymentInfo struct {
	g.Meta       `orm:"table:payment_info, do:true"`
	Id           interface{} //
	AppChannel   interface{} // app渠道
	PayType      interface{} // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	PaymentName  interface{} // 支付渠道：比如Google、OVO、DANA等等
	CurrencyCode interface{} // 国家数字码，比如印尼：360；马来：458  0为默认
	PayIcon      interface{} // 支付图标
	ChannelCode  interface{} // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	Outside      interface{} // 是否外部网页打开：0不需要，1需要
	State        interface{} // 状态：0下架、1上架
	SystemOs     interface{} // 所属系统；0所有,1安卓 2ios
	IsDefault    interface{} // 是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付
	Sort         interface{} // 排序，顺序排列
	Remark       interface{} // 说明备注
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

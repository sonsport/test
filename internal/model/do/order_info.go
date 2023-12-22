// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta             `orm:"table:order_info, do:true"`
	Oid                interface{} //
	Tid                interface{} // 票据Id
	UserId             interface{} // 用户Id
	Product            interface{} // 商品Id 对应commodity_info主键
	ProductValue       interface{} // 商品钻石
	ProductBonus       interface{} // 商品赠送钻石
	ProductDescription interface{} // 商品描述
	TotalFee           interface{} // 商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段
	Currency           interface{} // 价格对应的单位 比如：IDR、MYR 前台支付  对应commodity_info的currency字段
	CurrencyPrice      interface{} // 价格 分 对应商品手机支付当地货币价格
	IsFirstTopup       interface{} // 是否首次充值 0非首充 1首充  对应商品表的is_first_topup字段
	State              interface{} // 状态 0为待支付 1为已支付
	SystemOs           interface{} // 所属系统；0未知,1安卓 2ios
	AppChannel         interface{} // 所属渠道
	AppName            interface{} // 所属app名称
	PayType            interface{} // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	ChannelCode        interface{} // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	Remarks            interface{} // 备注说明
	CouponId           interface{} // 优惠券id
	ReferrerId         interface{} // 推荐人id
	CertificateUrl     interface{} // 代理充值凭证
	CreateTime         interface{} //
	UpdateTime         interface{} //
}

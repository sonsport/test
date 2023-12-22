// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	Oid                uint64 `json:"oid"                description:""`
	Tid                string `json:"tid"                description:"票据Id"`
	UserId             uint64 `json:"userId"             description:"用户Id"`
	Product            uint   `json:"product"            description:"商品Id 对应commodity_info主键"`
	ProductValue       uint64 `json:"productValue"       description:"商品钻石"`
	ProductBonus       uint   `json:"productBonus"       description:"商品赠送钻石"`
	ProductDescription string `json:"productDescription" description:"商品描述"`
	TotalFee           uint64 `json:"totalFee"           description:"商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段"`
	Currency           string `json:"currency"           description:"价格对应的单位 比如：IDR、MYR 前台支付  对应commodity_info的currency字段"`
	CurrencyPrice      uint64 `json:"currencyPrice"      description:"价格 分 对应商品手机支付当地货币价格"`
	IsFirstTopup       int    `json:"isFirstTopup"       description:"是否首次充值 0非首充 1首充  对应商品表的is_first_topup字段"`
	State              int    `json:"state"              description:"状态 0为待支付 1为已支付"`
	SystemOs           int    `json:"systemOs"           description:"所属系统；0未知,1安卓 2ios"`
	AppChannel         string `json:"appChannel"         description:"所属渠道"`
	AppName            string `json:"appName"            description:"所属app名称"`
	PayType            int    `json:"payType"            description:"支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay"`
	ChannelCode        int    `json:"channelCode"        description:"支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定"`
	Remarks            string `json:"remarks"            description:"备注说明"`
	CouponId           int64  `json:"couponId"           description:"优惠券id"`
	ReferrerId         int64  `json:"referrerId"         description:"推荐人id"`
	CertificateUrl     string `json:"certificateUrl"     description:"代理充值凭证"`
	CreateTime         int64  `json:"createTime"         description:""`
	UpdateTime         int64  `json:"updateTime"         description:""`
}

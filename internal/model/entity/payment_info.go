// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PaymentInfo is the golang structure for table payment_info.
type PaymentInfo struct {
	Id           uint   `json:"id"           description:""`
	AppChannel   string `json:"appChannel"   description:"app渠道"`
	PayType      int    `json:"payType"      description:"支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay"`
	PaymentName  string `json:"paymentName"  description:"支付渠道：比如Google、OVO、DANA等等"`
	CurrencyCode int    `json:"currencyCode" description:"国家数字码，比如印尼：360；马来：458  0为默认"`
	PayIcon      string `json:"payIcon"      description:"支付图标"`
	ChannelCode  int    `json:"channelCode"  description:"支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定"`
	Outside      int    `json:"outside"      description:"是否外部网页打开：0不需要，1需要"`
	State        int    `json:"state"        description:"状态：0下架、1上架"`
	SystemOs     int    `json:"systemOs"     description:"所属系统；0所有,1安卓 2ios"`
	IsDefault    int    `json:"isDefault"    description:"是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付"`
	Sort         int    `json:"sort"         description:"排序，顺序排列"`
	Remark       string `json:"remark"       description:"说明备注"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

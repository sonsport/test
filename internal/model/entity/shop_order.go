// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShopOrder is the golang structure for table shop_order.
type ShopOrder struct {
	Id         uint64 `json:"id"         description:""`
	ShopId     uint   `json:"shopId"     description:"商品ID"`
	Num        uint   `json:"num"        description:"购买数量（月）"`
	BuyUid     int64  `json:"buyUid"     description:"购买人uid"`
	ReceiveUid int64  `json:"receiveUid" description:"接收人"`
	CoinNumber uint64 `json:"coinNumber" description:"购买金额 钻石数（不扩大）"`
	Desc       string `json:"desc"       description:"订单描述"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

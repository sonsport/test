// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrder is the golang structure of table shop_order for DAO operations like Where/Data.
type ShopOrder struct {
	g.Meta     `orm:"table:shop_order, do:true"`
	Id         interface{} //
	ShopId     interface{} // 商品ID
	Num        interface{} // 购买数量（月）
	BuyUid     interface{} // 购买人uid
	ReceiveUid interface{} // 接收人
	CoinNumber interface{} // 购买金额 钻石数（不扩大）
	Desc       interface{} // 订单描述
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

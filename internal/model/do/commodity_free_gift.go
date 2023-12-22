// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommodityFreeGift is the golang structure of table commodity_free_gift for DAO operations like Where/Data.
type CommodityFreeGift struct {
	g.Meta   `orm:"table:commodity_free_gift, do:true"`
	Id       interface{} //
	Cid      interface{} //
	State    interface{} // 状态 1 启用
	GiftType interface{} // 赠送类型	1 装扮  2 经验
	GiftLink interface{} // 赠送关联	商品id
	GiftVal  interface{} // 赠送数	装扮为天 经验为经验值
}

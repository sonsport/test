// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommodityInfo is the golang structure of table commodity_info for DAO operations like Where/Data.
type CommodityInfo struct {
	g.Meta       `orm:"table:commodity_info, do:true"`
	Cid          interface{} // 商品Id 唯一约束
	PaymentId    interface{} // 关联payment_info表主键
	Name         interface{} // 名称 比如80钻
	Price        interface{} // 价格 分
	Value        interface{} // 钻石数
	Bonus        interface{} // 赠送钻石数
	Currency     interface{} // 价格对应的单位 比如：IDR、MYR 前台支付
	TotalFeeCent interface{} // 对应印尼盾价格 单位分 统计使用
	IsFirstTopup interface{} // 是否首次充值优惠档位，0默认，1首充优惠档位、2非首充档位
	Sort         interface{} // 排序、顺序
	State        interface{} // 状态；0下架；1上架
	IsDefault    interface{} // 是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付
	OriginBonus  interface{} // 折扣前赠送
	DiscountType interface{} // 折扣类型，1.首充折扣
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

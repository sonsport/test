// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopGoods is the golang structure of table shop_goods for DAO operations like Where/Data.
type ShopGoods struct {
	g.Meta         `orm:"table:shop_goods, do:true"`
	Id             interface{} // 自增ID
	CategoryId     interface{} // 类型
	Name           interface{} // 商品名称
	Status         interface{} // 商品状态 0下架 1上架
	IconUrl        interface{} // 图标地址
	ResUrl         interface{} // 资源地址
	IconSize       interface{} // 图标尺寸
	Price          interface{} // 单价 钻石数
	ForeverPrice   interface{} // 永久售卖价值 后续使用对于有永久售卖的商品
	Desc           interface{} // 商品描述
	DiscountPrice  interface{} // 折扣 json 配置不同天数的价格
	UserLevelLimit interface{} // 用户等级限制
	ShelfTime      interface{} // 上架时间
	IsRenewable    interface{} // 是否可续费 该字段基本是没用
	SortWeight     interface{} // 排序权重 升序
	SellType       interface{} // 售卖类型 1-售卖商品；2-活动商品
	Quality        interface{} // 品质
	Extra          interface{} // 扩展属性
	CreateTime     interface{} //
	UpdateTime     interface{} //
	Number         interface{} // 号码
}

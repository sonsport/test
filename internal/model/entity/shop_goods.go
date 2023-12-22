// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShopGoods is the golang structure for table shop_goods.
type ShopGoods struct {
	Id             uint   `json:"id"             description:"自增ID"`
	CategoryId     uint   `json:"categoryId"     description:"类型"`
	Name           string `json:"name"           description:"商品名称"`
	Status         int    `json:"status"         description:"商品状态 0下架 1上架"`
	IconUrl        string `json:"iconUrl"        description:"图标地址"`
	ResUrl         string `json:"resUrl"         description:"资源地址"`
	IconSize       string `json:"iconSize"       description:"图标尺寸"`
	Price          uint64 `json:"price"          description:"单价 钻石数"`
	ForeverPrice   uint64 `json:"foreverPrice"   description:"永久售卖价值 后续使用对于有永久售卖的商品"`
	Desc           string `json:"desc"           description:"商品描述"`
	DiscountPrice  string `json:"discountPrice"  description:"折扣 json 配置不同天数的价格"`
	UserLevelLimit int    `json:"userLevelLimit" description:"用户等级限制"`
	ShelfTime      int64  `json:"shelfTime"      description:"上架时间"`
	IsRenewable    int    `json:"isRenewable"    description:"是否可续费 该字段基本是没用"`
	SortWeight     int    `json:"sortWeight"     description:"排序权重 升序"`
	SellType       int    `json:"sellType"       description:"售卖类型 1-售卖商品；2-活动商品"`
	Quality        string `json:"quality"        description:"品质"`
	Extra          string `json:"extra"          description:"扩展属性"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
	Number         int64  `json:"number"         description:"号码"`
}

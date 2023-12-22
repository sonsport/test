// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommodityInfo is the golang structure for table commodity_info.
type CommodityInfo struct {
	Cid          uint64 `json:"cid"          description:"商品Id 唯一约束"`
	PaymentId    uint   `json:"paymentId"    description:"关联payment_info表主键"`
	Name         string `json:"name"         description:"名称 比如80钻"`
	Price        uint64 `json:"price"        description:"价格 分"`
	Value        uint64 `json:"value"        description:"钻石数"`
	Bonus        uint   `json:"bonus"        description:"赠送钻石数"`
	Currency     string `json:"currency"     description:"价格对应的单位 比如：IDR、MYR 前台支付"`
	TotalFeeCent uint64 `json:"totalFeeCent" description:"对应印尼盾价格 单位分 统计使用"`
	IsFirstTopup int    `json:"isFirstTopup" description:"是否首次充值优惠档位，0默认，1首充优惠档位、2非首充档位"`
	Sort         int    `json:"sort"         description:"排序、顺序"`
	State        int    `json:"state"        description:"状态；0下架；1上架"`
	IsDefault    int    `json:"isDefault"    description:"是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付"`
	OriginBonus  int    `json:"originBonus"  description:"折扣前赠送"`
	DiscountType int    `json:"discountType" description:"折扣类型，1.首充折扣"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

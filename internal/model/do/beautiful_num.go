// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BeautifulNum is the golang structure of table beautiful_num for DAO operations like Where/Data.
type BeautifulNum struct {
	g.Meta         `orm:"table:beautiful_num, do:true"`
	Id             interface{} //
	SuperId        interface{} // 靓号
	ClassifyId     interface{} // 分类id
	IconUrl        interface{} // icon地址
	ShopStatus     interface{} // 状态 1可用 2不可用
	SortWeight     interface{} // 排序权重
	SkuListPrice   interface{} // 规格列表价格
	UserLevelLimit interface{} // 用户等级限制
	Price          interface{} // 单价
	UserId         interface{} // 用户id
	IsWear         interface{} // 是否佩戴
	Experience     interface{} // 经验值
	ExpirationTime interface{} // 过期时间
	ReclaimTime    interface{} // 回收时间
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AreaInfo is the golang structure of table area_info for DAO operations like Where/Data.
type AreaInfo struct {
	g.Meta       `orm:"table:area_info, do:true"`
	Id           interface{} //
	AreaCode     interface{} // 地区
	Path         interface{} //
	Title        interface{} //
	TitleCn      interface{} //
	Currency     interface{} // 货币编码
	CurrencyCode interface{} // 国家编码
	Sort         interface{} //
	IsRich       interface{} // 是否富有，0普通，1富有
	IsShow       interface{} //
	CreatedTime  interface{} //
	UpdatedTime  interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NobleInfo is the golang structure of table noble_info for DAO operations like Where/Data.
type NobleInfo struct {
	g.Meta        `orm:"table:noble_info, do:true"`
	Id            interface{} //
	NobleName     interface{} // 贵族名称
	NobleStates   interface{} // 贵族状态 1 上架
	NobleIcon     interface{} // 贵族图标
	CurrentPrice  interface{} // 现价
	OriginalPrice interface{} // 原价
	NobleType     interface{} // 贵族类型 1 vip
	NobleLevel    interface{} // 贵族等级
	NobleDay      interface{} // 贵族天数
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

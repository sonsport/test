// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NobleConfig is the golang structure of table noble_config for DAO operations like Where/Data.
type NobleConfig struct {
	g.Meta        `orm:"table:noble_config, do:true"`
	Id            interface{} //
	NobleId       interface{} // 贵族id
	PrivilegeId   interface{} // 特权id
	PrivilegeLink interface{} // 特权值 类型为加成类型时val为倍数、商品时为商品id
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

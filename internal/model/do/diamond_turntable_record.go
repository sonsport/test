// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableRecord is the golang structure of table diamond_turntable_record for DAO operations like Where/Data.
type DiamondTurntableRecord struct {
	g.Meta         `orm:"table:diamond_turntable_record, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户id
	Cid            interface{} // 配置表id
	Gid            interface{} // 商品表id
	Num            interface{} // 数量
	TurntableCount interface{} // 奖池期数
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

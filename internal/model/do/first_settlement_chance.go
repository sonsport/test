// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FirstSettlementChance is the golang structure of table first_settlement_chance for DAO operations like Where/Data.
type FirstSettlementChance struct {
	g.Meta     `orm:"table:first_settlement_chance, do:true"`
	Id         interface{} // 自增Id
	Week       interface{} // 周，一年中的第几周 例如:202230 2022年的第30周
	UserId     interface{} // 用户ID
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

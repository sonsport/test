// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// IncomeInfo is the golang structure of table income_info for DAO operations like Where/Data.
type IncomeInfo struct {
	g.Meta                     `orm:"table:income_info, do:true"`
	Id                         interface{} // id
	UserId                     interface{} // 用户id
	TotalIncome                interface{} // 总收益 单位分
	BalancesIncome             interface{} // 收益余额 单位分
	TotalCommissionDiamonds    interface{} //
	BalancesCommissionDiamonds interface{} //
	CreateTime                 interface{} // 总佣金收入 扩大10000
	UpdateTime                 interface{} // 佣金收入 扩大10000
}

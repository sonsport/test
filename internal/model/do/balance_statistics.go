// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceStatistics is the golang structure of table balance_statistics for DAO operations like Where/Data.
type BalanceStatistics struct {
	g.Meta        `orm:"table:balance_statistics, do:true"`
	Id            interface{} //
	DepletionType interface{} // 消费类型
	Diamonds      interface{} // 钻石变动数额
	BeginTime     interface{} // 开始时间
	EndTime       interface{} // 结束时间
	TimeLevel     interface{} // 时间维度：1.每小时，2.每天
	BigType       interface{} // 消费类型 1支出 2收入
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

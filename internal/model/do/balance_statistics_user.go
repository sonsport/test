// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceStatisticsUser is the golang structure of table balance_statistics_user for DAO operations like Where/Data.
type BalanceStatisticsUser struct {
	g.Meta        `orm:"table:balance_statistics_user, do:true"`
	Id            interface{} //
	UserId        interface{} //
	Date          interface{} // 日期
	DepletionType interface{} // 消费类型
	Type          interface{} // 消费类型 1支出 2收入
	Diamonds      interface{} // 钻石变动数额
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceRecord is the golang structure of table balance_record for DAO operations like Where/Data.
type BalanceRecord struct {
	g.Meta        `orm:"table:balance_record, do:true"`
	Id            interface{} //
	SourceUserId  interface{} // 用户Id
	TargetUserId  interface{} // 主播Id
	DepletionType interface{} // 消费类型
	Diamonds      interface{} // 钻石变动数额
	BeforeChange  interface{} // 变更之前余额
	AfterChange   interface{} // 变更之后余额
	LinkId        interface{} // 引起余额变更的id
	Type          interface{} // 消费类型 1支出 2收入
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

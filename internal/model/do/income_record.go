// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// IncomeRecord is the golang structure of table income_record for DAO operations like Where/Data.
type IncomeRecord struct {
	g.Meta       `orm:"table:income_record, do:true"`
	Id           interface{} // id
	UserId       interface{} // 用户id
	SourceId     interface{} // 收益来源用户的id
	UserRecordId interface{} // 对应用户扣费记录id
	IncomeType   interface{} // 收益类型
	LinkId       interface{} // 引起变更的id
	Income       interface{} // 收益变动数额
	BeforeChange interface{} // 变更之前余额
	AfterChange  interface{} // 变更之后余额
	CreateTime   interface{} // createdAt
	UpdateTime   interface{} // updatedAt
}

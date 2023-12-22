// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BlackList is the golang structure of table black_list for DAO operations like Where/Data.
type BlackList struct {
	g.Meta     `orm:"table:black_list, do:true"`
	Id         interface{} // ID
	UserId     interface{} // 用户ID
	OperatorId interface{} // 操作者ID
	BlockedId  interface{} // 被拉黑的用户ID
	EndTime    interface{} // 拉黑结束时间,为0就是永久拉黑
	CreateTime interface{} //
	UpdateTime interface{} //
}

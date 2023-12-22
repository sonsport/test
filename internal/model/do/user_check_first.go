// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserCheckFirst is the golang structure of table user_check_first for DAO operations like Where/Data.
type UserCheckFirst struct {
	g.Meta     `orm:"table:user_check_first, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户Id
	CheckType  interface{} // 检查类型
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 最后更新时间
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysAdmin is the golang structure of table sys_admin for DAO operations like Where/Data.
type SysAdmin struct {
	g.Meta     `orm:"table:sys_admin, do:true"`
	Id         interface{} //
	UserName   interface{} // 用户名
	NickName   interface{} // 昵称
	Password   interface{} // 密码
	CreateTime interface{} //
	UpdateTime interface{} //
	Status     interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysAdminRole is the golang structure of table sys_admin_role for DAO operations like Where/Data.
type SysAdminRole struct {
	g.Meta   `orm:"table:sys_admin_role, do:true"`
	Id       interface{} // ID
	AdminId  interface{} // 用户ID
	RoleId   interface{} // 角色ID
	ServerId interface{} // 运营号id
}

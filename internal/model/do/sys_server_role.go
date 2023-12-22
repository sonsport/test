// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysServerRole is the golang structure of table sys_server_role for DAO operations like Where/Data.
type SysServerRole struct {
	g.Meta     `orm:"table:sys_server_role, do:true"`
	Id         interface{} // ID
	AdminId    interface{} // 关联admin主键
	Role       interface{} // 角色0普通，1管理
	Group      interface{} // 管理组，1雅加达，2三宝垄
	CreateTime interface{} // 创建日期
	UpdateTime interface{} // 更新时间
}

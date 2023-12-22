// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NoblePrivilege is the golang structure of table noble_privilege for DAO operations like Where/Data.
type NoblePrivilege struct {
	g.Meta          `orm:"table:noble_privilege, do:true"`
	PrivilegeId     interface{} //
	PrivilegeName   interface{} // 特权名称
	PrivilegeIcon   interface{} // 特权图片
	PrivilegeStates interface{} // 特权状态 1 正常
	IsPrivilege     interface{} // 是否特权 1特权 2商品
	PrivilegeType   interface{} // 特权类型
	PrivilegeDesc   interface{} // 描述 多语言配置 {en:xxx}
	NobleType       interface{} // 贵族类型 1 vip
	Sort            interface{} // 排序
	CreateTime      interface{} //
	UpdateTime      interface{} //
}

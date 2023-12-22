// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	MenuId     interface{} // ID
	Pid        interface{} // 上级菜单ID
	Title      interface{} // 菜单标题
	Name       interface{} // 组件名称
	Component  interface{} // 组件
	MenuSort   interface{} // 排序
	Icon       interface{} // 图标
	Path       interface{} // 链接地址
	CreateTime interface{} // 创建日期
	UpdateTime interface{} // 更新时间
	TitleCn    interface{} // 中文菜单标题
}

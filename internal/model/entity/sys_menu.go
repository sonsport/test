// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	MenuId     int64  `json:"menuId"     description:"ID"`
	Pid        int64  `json:"pid"        description:"上级菜单ID"`
	Title      string `json:"title"      description:"菜单标题"`
	Name       string `json:"name"       description:"组件名称"`
	Component  string `json:"component"  description:"组件"`
	MenuSort   int    `json:"menuSort"   description:"排序"`
	Icon       string `json:"icon"       description:"图标"`
	Path       string `json:"path"       description:"链接地址"`
	CreateTime int64  `json:"createTime" description:"创建日期"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
	TitleCn    string `json:"titleCn"    description:"中文菜单标题"`
}

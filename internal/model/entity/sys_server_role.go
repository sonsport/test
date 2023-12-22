// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysServerRole is the golang structure for table sys_server_role.
type SysServerRole struct {
	Id         int64 `json:"id"         description:"ID"`
	AdminId    int64 `json:"adminId"    description:"关联admin主键"`
	Role       int   `json:"role"       description:"角色0普通，1管理"`
	Group      int   `json:"group"      description:"管理组，1雅加达，2三宝垄"`
	CreateTime int64 `json:"createTime" description:"创建日期"`
	UpdateTime int64 `json:"updateTime" description:"更新时间"`
}

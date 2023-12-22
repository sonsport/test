// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysAdminRole is the golang structure for table sys_admin_role.
type SysAdminRole struct {
	Id       int64 `json:"id"       description:"ID"`
	AdminId  int64 `json:"adminId"  description:"用户ID"`
	RoleId   int64 `json:"roleId"   description:"角色ID"`
	ServerId int64 `json:"serverId" description:"运营号id"`
}

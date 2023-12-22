// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NoblePrivilege is the golang structure for table noble_privilege.
type NoblePrivilege struct {
	PrivilegeId     int    `json:"privilegeId"     description:""`
	PrivilegeName   string `json:"privilegeName"   description:"特权名称"`
	PrivilegeIcon   string `json:"privilegeIcon"   description:"特权图片"`
	PrivilegeStates int    `json:"privilegeStates" description:"特权状态 1 正常"`
	IsPrivilege     int    `json:"isPrivilege"     description:"是否特权 1特权 2商品"`
	PrivilegeType   int    `json:"privilegeType"   description:"特权类型"`
	PrivilegeDesc   string `json:"privilegeDesc"   description:"描述 多语言配置 {en:xxx}"`
	NobleType       int    `json:"nobleType"       description:"贵族类型 1 vip"`
	Sort            int    `json:"sort"            description:"排序"`
	CreateTime      int64  `json:"createTime"      description:""`
	UpdateTime      int64  `json:"updateTime"      description:""`
}

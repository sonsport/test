// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysAdmin is the golang structure for table sys_admin.
type SysAdmin struct {
	Id         uint   `json:"id"         description:""`
	UserName   string `json:"userName"   description:"用户名"`
	NickName   string `json:"nickName"   description:"昵称"`
	Password   string `json:"password"   description:"密码"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
	Status     int    `json:"status"     description:""`
}

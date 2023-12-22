// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// InviterInfo is the golang structure of table inviter_info for DAO operations like Where/Data.
type InviterInfo struct {
	g.Meta          `orm:"table:inviter_info, do:true"`
	Id              interface{} //
	UserId          interface{} //
	Stat            interface{} //
	InviteUserCount interface{} //
	CreateTime      interface{} //
	UpdateTime      interface{} //
}

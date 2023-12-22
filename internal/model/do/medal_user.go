// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MedalUser is the golang structure of table medal_user for DAO operations like Where/Data.
type MedalUser struct {
	g.Meta      `orm:"table:medal_user, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户id
	MedalId     interface{} // 勋章id
	AccessCount interface{} // 获得次数
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

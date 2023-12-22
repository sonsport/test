// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SafeInfo is the golang structure of table safe_info for DAO operations like Where/Data.
type SafeInfo struct {
	g.Meta     `orm:"table:safe_info, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户Id
	Content    interface{} // 内容
	SafeType   interface{} // 类型，1头像，2昵称，3签名，后续扩展
	State      interface{} // 状态，0待审核，1审核通过，2审核不通过
	CreateTime interface{} //
	UpdateTime interface{} //
}

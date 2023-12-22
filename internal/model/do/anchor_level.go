// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorLevel is the golang structure of table anchor_level for DAO operations like Where/Data.
type AnchorLevel struct {
	g.Meta     `orm:"table:anchor_level, do:true"`
	Id         interface{} // 自增Id
	UserId     interface{} // 用户ID
	Level      interface{} // 等级
	Remark     interface{} // 等级
	State      interface{} // 状态1有效，2无效
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

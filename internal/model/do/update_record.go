// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateRecord is the golang structure of table update_record for DAO operations like Where/Data.
type UpdateRecord struct {
	g.Meta     `orm:"table:update_record, do:true"`
	Id         interface{} // 主键id
	UserId     interface{} // 用户id
	AdminId    interface{} // 管理员id
	UpdateType interface{} // 调整类型，1星级，2公会，3标签，4备注
	BeforeBody interface{} // 修改前内容
	AfterBody  interface{} // 修改后内容
	Remark     interface{} // 修改后内容
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

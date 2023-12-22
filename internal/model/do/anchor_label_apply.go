// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorLabelApply is the golang structure of table anchor_label_apply for DAO operations like Where/Data.
type AnchorLabelApply struct {
	g.Meta     `orm:"table:anchor_label_apply, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Label      interface{} // 标签
	States     interface{} // 状态 1 待审核 2 已通过 3 已驳回
	CreateTime interface{} //
	UpdateTime interface{} //
}

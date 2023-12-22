// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorRankRule is the golang structure of table anchor_rank_rule for DAO operations like Where/Data.
type AnchorRankRule struct {
	g.Meta     `orm:"table:anchor_rank_rule, do:true"`
	Id         interface{} // 主键ID
	Level      interface{} // 等级从0开始, 0为等级1
	Min        interface{} // 最小值
	Max        interface{} // 最大值
	CreateTime interface{} //
	UpdateTime interface{} //
}

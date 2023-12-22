// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyBattle is the golang structure of table apply_battle for DAO operations like Where/Data.
type ApplyBattle struct {
	g.Meta           `orm:"table:apply_battle, do:true"`
	Id               interface{} // 自增id
	GuildId          interface{} // 公会Id 关联guild_info自增
	UserId           interface{} // 主播ID
	AdminOperateName interface{} // 审核人id
	Status           interface{} // 状态：0待审核 1通过 2不通过
	ApplyType        interface{} // 报名类型，1公会，2 主播
	CreateTime       interface{} // 开始时间
	UpdateTime       interface{} // 开始时间
}

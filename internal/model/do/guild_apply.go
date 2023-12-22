// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildApply is the golang structure of table guild_apply for DAO operations like Where/Data.
type GuildApply struct {
	g.Meta     `orm:"table:guild_apply, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 申请人用户Id
	GuildName  interface{} // 公会名称
	AnchorNum  interface{} // 主播数量
	Whatsapp   interface{} // 联系WhatsApp
	State      interface{} // 状态：0待审核 1通过 2不通过
	Operator   interface{} // 审核人
	Remark     interface{} // 备注
	CreateTime interface{} //
	UpdateTime interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildInfo is the golang structure of table guild_info for DAO operations like Where/Data.
type GuildInfo struct {
	g.Meta       `orm:"table:guild_info, do:true"`
	Id           interface{} // 公会Id
	UserId       interface{} // 公会长Id
	GuildName    interface{} // 公会名称
	Gcode        interface{} // 公会邀请码
	MemberCount  interface{} // 成员数量
	Portrait     interface{} // 公会头像
	Status       interface{} // 状态
	Notice       interface{} // 公告
	CutRate      interface{} // 公会抽成 设置范围0-100
	IsPayToGuild interface{} // 工资是否发放给公会 0非 1是
	Remark       interface{} // 备注
	CreateTime   interface{} //
	UpdateTime   interface{} //
	ServerId     interface{} // 运营号id
	SourceType   interface{} // 绑定来源0无，1运营号直接邀请，2工会邀请，3后台绑定
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildAnchor is the golang structure of table guild_anchor for DAO operations like Where/Data.
type GuildAnchor struct {
	g.Meta          `orm:"table:guild_anchor, do:true"`
	Id              interface{} // 系统自增
	GuildId         interface{} // 公会Id 关联guild_info自增
	UserId          interface{} // 用户Id
	GuildRole       interface{} // 1公会长 2公会成员
	Status          interface{} // 状态 1 正常；2 失效
	Remark          interface{} // 备注
	JoinTime        interface{} // 加入时间
	StarLevel       interface{} // 主播星级
	LabelType       interface{} // 标签类型
	Label           interface{} // 主播标签
	Score           interface{} // 主播基础分数
	LivePermissions interface{} // 1钻石房，2密码房，3钻石房和密码房
	CreateTime      interface{} // 创建时间
	UpdateTime      interface{} // 更新时间
	CloudRecording  interface{} // 录制状态 0不录制，1录制
}

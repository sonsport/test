// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildAnchorRecord is the golang structure of table guild_anchor_record for DAO operations like Where/Data.
type GuildAnchorRecord struct {
	g.Meta               `orm:"table:guild_anchor_record, do:true"`
	Id                   interface{} // 主键id
	UserId               interface{} // 用户id
	AdminId              interface{} // 管理员id
	ExpirationTime       interface{} // 过期时间
	BackgroundAdjustment interface{} // 是否后台调整
	UpdateType           interface{} // 调整类型，1星级，2公会，3标签
	IsHot                interface{} // 是否在热门列表
	OldContent           interface{} // 修改前内容
	NewContent           interface{} // 修改后内容
	CreateTime           interface{} // 创建时间
	UpdateTime           interface{} // 更新时间
}

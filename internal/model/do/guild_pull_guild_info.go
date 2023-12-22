// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildPullGuildInfo is the golang structure of table guild_pull_guild_info for DAO operations like Where/Data.
type GuildPullGuildInfo struct {
	g.Meta       `orm:"table:guild_pull_guild_info, do:true"`
	Id           interface{} //
	PullerUserId interface{} // A拉B，这里存放A公会长ID
	PulledUserId interface{} // A拉B，这里存放B公会长ID
	ManagerId    interface{} // 操作管理员ID
	CreatedTime  interface{} //
	UpdatedTime  interface{} //
}

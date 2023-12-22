// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ServerWeekStatistics is the golang structure of table server_week_statistics for DAO operations like Where/Data.
type ServerWeekStatistics struct {
	g.Meta            `orm:"table:server_week_statistics, do:true"`
	Id                interface{} // 自增id
	UserId            interface{} // 运营号id
	Week              interface{} // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildNum          interface{} // 公会数量
	EffectiveGuildNum interface{} // 有效公会数量
	NewGuildNum       interface{} // 新增公会数量
	CreateTime        interface{} // 创建时间
	UpdateTime        interface{} // 更新时间
}

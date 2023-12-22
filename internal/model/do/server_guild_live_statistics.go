// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ServerGuildLiveStatistics is the golang structure of table server_guild_live_statistics for DAO operations like Where/Data.
type ServerGuildLiveStatistics struct {
	g.Meta                `orm:"table:server_guild_live_statistics, do:true"`
	Id                    interface{} // 主键ID
	BeginTime             interface{} // 开始时间
	EndTime               interface{} // 结束时间
	GuildId               interface{} // 工会id
	GuildName             interface{} // 工会名称
	ServerId              interface{} // 运营号id
	LiveAnchorCount       interface{} // 主播开播数
	LiveStarAnchorCount   interface{} // 标星主播开播数
	ActiveAnchorCount     interface{} // 主播在线数
	StarActiveAnchorCount interface{} // 标星主播在线数
}

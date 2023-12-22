// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ServerGuildLiveStatistics is the golang structure for table server_guild_live_statistics.
type ServerGuildLiveStatistics struct {
	Id                    uint   `json:"id"                    description:"主键ID"`
	BeginTime             int64  `json:"beginTime"             description:"开始时间"`
	EndTime               int64  `json:"endTime"               description:"结束时间"`
	GuildId               int64  `json:"guildId"               description:"工会id"`
	GuildName             string `json:"guildName"             description:"工会名称"`
	ServerId              int64  `json:"serverId"              description:"运营号id"`
	LiveAnchorCount       int64  `json:"liveAnchorCount"       description:"主播开播数"`
	LiveStarAnchorCount   int64  `json:"liveStarAnchorCount"   description:"标星主播开播数"`
	ActiveAnchorCount     int64  `json:"activeAnchorCount"     description:"主播在线数"`
	StarActiveAnchorCount int64  `json:"starActiveAnchorCount" description:"标星主播在线数"`
}

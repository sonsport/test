// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ServerWeekStatistics is the golang structure for table server_week_statistics.
type ServerWeekStatistics struct {
	Id                uint  `json:"id"                description:"自增id"`
	UserId            int64 `json:"userId"            description:"运营号id"`
	Week              int   `json:"week"              description:"周，一年中的第几周 例如:202230 2022年的第30周"`
	GuildNum          int   `json:"guildNum"          description:"公会数量"`
	EffectiveGuildNum int   `json:"effectiveGuildNum" description:"有效公会数量"`
	NewGuildNum       int   `json:"newGuildNum"       description:"新增公会数量"`
	CreateTime        int64 `json:"createTime"        description:"创建时间"`
	UpdateTime        int64 `json:"updateTime"        description:"更新时间"`
}

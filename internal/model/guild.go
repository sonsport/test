package model

type GuildStatistics struct {
	GuildNum          int `json:"guild_num"`           //公会数量
	EffectiveGuildNum int `json:"effective_guild_num"` //公会数量
	NewGuildNum       int `json:"new_guild_num"`       //新增公会数量
}

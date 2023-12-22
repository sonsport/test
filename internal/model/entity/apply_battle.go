// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ApplyBattle is the golang structure for table apply_battle.
type ApplyBattle struct {
	Id               int64  `json:"id"               description:"自增id"`
	GuildId          int    `json:"guildId"          description:"公会Id 关联guild_info自增"`
	UserId           int64  `json:"userId"           description:"主播ID"`
	AdminOperateName string `json:"adminOperateName" description:"审核人id"`
	Status           int    `json:"status"           description:"状态：0待审核 1通过 2不通过"`
	ApplyType        int    `json:"applyType"        description:"报名类型，1公会，2 主播"`
	CreateTime       int64  `json:"createTime"       description:"开始时间"`
	UpdateTime       int64  `json:"updateTime"       description:"开始时间"`
}

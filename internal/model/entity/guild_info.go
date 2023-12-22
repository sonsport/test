// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildInfo is the golang structure for table guild_info.
type GuildInfo struct {
	Id           uint   `json:"id"           description:"公会Id"`
	UserId       int64  `json:"userId"       description:"公会长Id"`
	GuildName    string `json:"guildName"    description:"公会名称"`
	Gcode        int    `json:"gcode"        description:"公会邀请码"`
	MemberCount  int    `json:"memberCount"  description:"成员数量"`
	Portrait     string `json:"portrait"     description:"公会头像"`
	Status       int    `json:"status"       description:"状态"`
	Notice       string `json:"notice"       description:"公告"`
	CutRate      uint   `json:"cutRate"      description:"公会抽成 设置范围0-100"`
	IsPayToGuild int    `json:"isPayToGuild" description:"工资是否发放给公会 0非 1是"`
	Remark       string `json:"remark"       description:"备注"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
	ServerId     int64  `json:"serverId"     description:"运营号id"`
	SourceType   int    `json:"sourceType"   description:"绑定来源0无，1运营号直接邀请，2工会邀请，3后台绑定"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildAnchor is the golang structure for table guild_anchor.
type GuildAnchor struct {
	Id              uint64 `json:"id"              description:"系统自增"`
	GuildId         int    `json:"guildId"         description:"公会Id 关联guild_info自增"`
	UserId          int64  `json:"userId"          description:"用户Id"`
	GuildRole       int    `json:"guildRole"       description:"1公会长 2公会成员"`
	Status          int    `json:"status"          description:"状态 1 正常；2 失效"`
	Remark          string `json:"remark"          description:"备注"`
	JoinTime        int64  `json:"joinTime"        description:"加入时间"`
	StarLevel       int    `json:"starLevel"       description:"主播星级"`
	LabelType       int    `json:"labelType"       description:"标签类型"`
	Label           string `json:"label"           description:"主播标签"`
	Score           int    `json:"score"           description:"主播基础分数"`
	LivePermissions int    `json:"livePermissions" description:"1钻石房，2密码房，3钻石房和密码房"`
	CreateTime      int64  `json:"createTime"      description:"创建时间"`
	UpdateTime      int64  `json:"updateTime"      description:"更新时间"`
	CloudRecording  int    `json:"cloudRecording"  description:"录制状态 0不录制，1录制"`
}

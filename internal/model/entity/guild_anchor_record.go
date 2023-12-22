// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildAnchorRecord is the golang structure for table guild_anchor_record.
type GuildAnchorRecord struct {
	Id                   int64  `json:"id"                   description:"主键id"`
	UserId               int64  `json:"userId"               description:"用户id"`
	AdminId              int64  `json:"adminId"              description:"管理员id"`
	ExpirationTime       int64  `json:"expirationTime"       description:"过期时间"`
	BackgroundAdjustment int    `json:"backgroundAdjustment" description:"是否后台调整"`
	UpdateType           int    `json:"updateType"           description:"调整类型，1星级，2公会，3标签"`
	IsHot                int    `json:"isHot"                description:"是否在热门列表"`
	OldContent           string `json:"oldContent"           description:"修改前内容"`
	NewContent           string `json:"newContent"           description:"修改后内容"`
	CreateTime           int64  `json:"createTime"           description:"创建时间"`
	UpdateTime           int64  `json:"updateTime"           description:"更新时间"`
}

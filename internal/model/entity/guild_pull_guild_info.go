// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildPullGuildInfo is the golang structure for table guild_pull_guild_info.
type GuildPullGuildInfo struct {
	Id           int    `json:"id"           description:""`
	PullerUserId uint64 `json:"pullerUserId" description:"A拉B，这里存放A公会长ID"`
	PulledUserId int64  `json:"pulledUserId" description:"A拉B，这里存放B公会长ID"`
	ManagerId    int    `json:"managerId"    description:"操作管理员ID"`
	CreatedTime  int64  `json:"createdTime"  description:""`
	UpdatedTime  int64  `json:"updatedTime"  description:""`
}

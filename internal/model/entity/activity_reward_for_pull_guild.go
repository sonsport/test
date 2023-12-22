// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityRewardForPullGuild is the golang structure for table activity_reward_for_pull_guild.
type ActivityRewardForPullGuild struct {
	Id                int64  `json:"id"                description:""`
	UserId            uint64 `json:"userId"            description:"用户Id 主播id"`
	GuildId           int    `json:"guildId"           description:"公会Id 关联guild_info自增"`
	GuildUserId       int64  `json:"guildUserId"       description:"用户所属公会长Id （这里是B公会长ID）"`
	PullerGuildUserId int64  `json:"pullerGuildUserId" description:"A公会拉的B公会 这里存储A公会长的ID"`
	LiveTime          int64  `json:"liveTime"          description:"直播时长 单位秒"`
	IsReward          int    `json:"isReward"          description:"是否已经奖励 0未发放 1已发放"`
	Reward            int64  `json:"reward"            description:"奖励、存储印尼盾分 奖励给B公会长的"`
	RewardDate        int64  `json:"rewardDate"        description:"发放日期"`
	CreatedTime       int64  `json:"createdTime"       description:""`
	UpdatedTime       int64  `json:"updatedTime"       description:""`
}

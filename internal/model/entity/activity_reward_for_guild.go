// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityRewardForGuild is the golang structure for table activity_reward_for_guild.
type ActivityRewardForGuild struct {
	Id             int64  `json:"id"             description:""`
	UserId         uint64 `json:"userId"         description:"用户Id 主播id"`
	GuildId        int    `json:"guildId"        description:"公会Id 关联guild_info自增"`
	ActivityType   int    `json:"activityType"   description:"活动类型  0默认 1首播活动 2四周活动 后续扩展"`
	EffectLiveTime int64  `json:"effectLiveTime" description:"有效直播时长（s） 首播、四周活动存储热门直播时长"`
	EffectDay      int    `json:"effectDay"      description:"有效直播天"`
	IsReward       int    `json:"isReward"       description:"是否已经奖励 0未发放 1已发放"`
	Reward         int64  `json:"reward"         description:"奖励、存储印尼盾分，后续其他使用可以存储别的 根据不同的活动类型"`
	RewardDate     int64  `json:"rewardDate"     description:"发放日期"`
	CreatedTime    int64  `json:"createdTime"    description:""`
	UpdatedTime    int64  `json:"updatedTime"    description:""`
}

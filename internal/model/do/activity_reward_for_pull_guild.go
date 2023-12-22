// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRewardForPullGuild is the golang structure of table activity_reward_for_pull_guild for DAO operations like Where/Data.
type ActivityRewardForPullGuild struct {
	g.Meta            `orm:"table:activity_reward_for_pull_guild, do:true"`
	Id                interface{} //
	UserId            interface{} // 用户Id 主播id
	GuildId           interface{} // 公会Id 关联guild_info自增
	GuildUserId       interface{} // 用户所属公会长Id （这里是B公会长ID）
	PullerGuildUserId interface{} // A公会拉的B公会 这里存储A公会长的ID
	LiveTime          interface{} // 直播时长 单位秒
	IsReward          interface{} // 是否已经奖励 0未发放 1已发放
	Reward            interface{} // 奖励、存储印尼盾分 奖励给B公会长的
	RewardDate        interface{} // 发放日期
	CreatedTime       interface{} //
	UpdatedTime       interface{} //
}

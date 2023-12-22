// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRewardForGuild is the golang structure of table activity_reward_for_guild for DAO operations like Where/Data.
type ActivityRewardForGuild struct {
	g.Meta         `orm:"table:activity_reward_for_guild, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户Id 主播id
	GuildId        interface{} // 公会Id 关联guild_info自增
	ActivityType   interface{} // 活动类型  0默认 1首播活动 2四周活动 后续扩展
	EffectLiveTime interface{} // 有效直播时长（s） 首播、四周活动存储热门直播时长
	EffectDay      interface{} // 有效直播天
	IsReward       interface{} // 是否已经奖励 0未发放 1已发放
	Reward         interface{} // 奖励、存储印尼盾分，后续其他使用可以存储别的 根据不同的活动类型
	RewardDate     interface{} // 发放日期
	CreatedTime    interface{} //
	UpdatedTime    interface{} //
}

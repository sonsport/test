// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AnchorDayStatistics is the golang structure of table anchor_day_statistics for DAO operations like Where/Data.
type AnchorDayStatistics struct {
	g.Meta                  `orm:"table:anchor_day_statistics, do:true"`
	StatisticsId            interface{} // 主键id
	StatisticsTime          *gtime.Time // 统计时间
	UserId                  interface{} // 用户id
	GuildId                 interface{} // 工会id
	AllIncome               interface{} // 总收益
	NormalGiftIncome        interface{} // 普通礼物收益
	LuckyGiftIncome         interface{} // 幸运礼物收益
	LiveTime                interface{} // 直播时长（s）
	EffectLiveTime          interface{} // 有效直播时长（s） 大于30分钟的直播才算有效直播累计
	IsEffectDay             interface{} // 是否有效天 0非 1是
	PkCount                 interface{} // 总pk次数
	EffectivePkCount        interface{} // 有效pk次数
	WatchMemberCount        interface{} // 观看人数
	NewFansCount            interface{} // 新增粉丝数
	SendGiftPerson          interface{} // 新增送礼人数
	IsFirstLive             interface{} // 是否首播
	Role                    interface{} // 角色
	RewardDayForGuild       interface{} // 当天对公会的奖励（印尼盾 单位分），比如主播当天首播满1小时给公会的奖励
	RewardDayForGuildRemark interface{} // 活动奖励说明
	SumHotAward             interface{} // 热门钻石奖励
	HotEffectLive           interface{} // 热门开播
	CreateTime              interface{} // 创建时间
	DayWaterDiamond         interface{} // 直播天流水
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AnchorDayStatistics is the golang structure for table anchor_day_statistics.
type AnchorDayStatistics struct {
	StatisticsId            int64       `json:"statisticsId"            description:"主键id"`
	StatisticsTime          *gtime.Time `json:"statisticsTime"          description:"统计时间"`
	UserId                  int64       `json:"userId"                  description:"用户id"`
	GuildId                 int64       `json:"guildId"                 description:"工会id"`
	AllIncome               int64       `json:"allIncome"               description:"总收益"`
	NormalGiftIncome        int64       `json:"normalGiftIncome"        description:"普通礼物收益"`
	LuckyGiftIncome         int64       `json:"luckyGiftIncome"         description:"幸运礼物收益"`
	LiveTime                int         `json:"liveTime"                description:"直播时长（s）"`
	EffectLiveTime          int         `json:"effectLiveTime"          description:"有效直播时长（s） 大于30分钟的直播才算有效直播累计"`
	IsEffectDay             int         `json:"isEffectDay"             description:"是否有效天 0非 1是"`
	PkCount                 int         `json:"pkCount"                 description:"总pk次数"`
	EffectivePkCount        int         `json:"effectivePkCount"        description:"有效pk次数"`
	WatchMemberCount        int         `json:"watchMemberCount"        description:"观看人数"`
	NewFansCount            int         `json:"newFansCount"            description:"新增粉丝数"`
	SendGiftPerson          int         `json:"sendGiftPerson"          description:"新增送礼人数"`
	IsFirstLive             int         `json:"isFirstLive"             description:"是否首播"`
	Role                    int         `json:"role"                    description:"角色"`
	RewardDayForGuild       int64       `json:"rewardDayForGuild"       description:"当天对公会的奖励（印尼盾 单位分），比如主播当天首播满1小时给公会的奖励"`
	RewardDayForGuildRemark string      `json:"rewardDayForGuildRemark" description:"活动奖励说明"`
	SumHotAward             int         `json:"sumHotAward"             description:"热门钻石奖励"`
	HotEffectLive           int64       `json:"hotEffectLive"           description:"热门开播"`
	CreateTime              int64       `json:"createTime"              description:"创建时间"`
	DayWaterDiamond         int64       `json:"dayWaterDiamond"         description:"直播天流水"`
}

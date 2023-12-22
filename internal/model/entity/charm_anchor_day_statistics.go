// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CharmAnchorDayStatistics is the golang structure for table charm_anchor_day_statistics.
type CharmAnchorDayStatistics struct {
	Id                   int64       `json:"id"                   description:"主键id"`
	StatisticsTime       *gtime.Time `json:"statisticsTime"       description:"统计时间"`
	UserId               int64       `json:"userId"               description:"用户id"`
	GuildId              int64       `json:"guildId"              description:"工会id"`
	LiveTime             int         `json:"liveTime"             description:"直播时长（s）"`
	EffectLiveTime       int         `json:"effectLiveTime"       description:"有效直播时长（s） 大于30分钟的直播才算有效直播累计"`
	IsEffectDay          int         `json:"isEffectDay"          description:"是否有效天 0非 1是"`
	NewFansCount         int         `json:"newFansCount"         description:"新增粉丝数"`
	NewSendGiftPerson    int         `json:"newSendGiftPerson"    description:"新增送礼人数"`
	AllSendGiftPerson    int         `json:"allSendGiftPerson"    description:"总送礼人数"`
	AllActivePersonCount int         `json:"allActivePersonCount" description:"直播间活跃人数，发送过消息"`
	JoinRoomPersonCount  int         `json:"joinRoomPersonCount"  description:"进入房间总人数"`
	Remark               string      `json:"remark"               description:"备注"`
	CreateTime           int64       `json:"createTime"           description:""`
	UpdateTime           int64       `json:"updateTime"           description:""`
}

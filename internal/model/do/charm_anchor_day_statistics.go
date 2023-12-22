// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CharmAnchorDayStatistics is the golang structure of table charm_anchor_day_statistics for DAO operations like Where/Data.
type CharmAnchorDayStatistics struct {
	g.Meta               `orm:"table:charm_anchor_day_statistics, do:true"`
	Id                   interface{} // 主键id
	StatisticsTime       *gtime.Time // 统计时间
	UserId               interface{} // 用户id
	GuildId              interface{} // 工会id
	LiveTime             interface{} // 直播时长（s）
	EffectLiveTime       interface{} // 有效直播时长（s） 大于30分钟的直播才算有效直播累计
	IsEffectDay          interface{} // 是否有效天 0非 1是
	NewFansCount         interface{} // 新增粉丝数
	NewSendGiftPerson    interface{} // 新增送礼人数
	AllSendGiftPerson    interface{} // 总送礼人数
	AllActivePersonCount interface{} // 直播间活跃人数，发送过消息
	JoinRoomPersonCount  interface{} // 进入房间总人数
	Remark               interface{} // 备注
	CreateTime           interface{} //
	UpdateTime           interface{} //
}

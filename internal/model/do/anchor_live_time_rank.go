// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AnchorLiveTimeRank is the golang structure of table anchor_live_time_rank for DAO operations like Where/Data.
type AnchorLiveTimeRank struct {
	g.Meta            `orm:"table:anchor_live_time_rank, do:true"`
	Id                interface{} //
	UserId            interface{} // 用户ID
	RankDateTime      *gtime.Time // 排行日期
	EffectLiveSeconds interface{} // 有效直播时长
	CreateTime        interface{} //
	UpdateTime        interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AnchorLiveTimeRank is the golang structure for table anchor_live_time_rank.
type AnchorLiveTimeRank struct {
	Id                int         `json:"id"                description:""`
	UserId            int64       `json:"userId"            description:"用户ID"`
	RankDateTime      *gtime.Time `json:"rankDateTime"      description:"排行日期"`
	EffectLiveSeconds int         `json:"effectLiveSeconds" description:"有效直播时长"`
	CreateTime        int64       `json:"createTime"        description:""`
	UpdateTime        int64       `json:"updateTime"        description:""`
}

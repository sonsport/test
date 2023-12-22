// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PrizePoolConfig is the golang structure for table prize_pool_config.
type PrizePoolConfig struct {
	Id              int64       `json:"id"              description:"主键id"`
	Gid             int64       `json:"gid"             description:"幸运礼物id"`
	StartTime       *gtime.Time `json:"startTime"       description:"配置生效时间"`
	Type            int         `json:"type"            description:"配置类型 1:默认 2:其他"`
	CashbackRatioId int64       `json:"cashbackRatioId" description:"奖池返现配置id"`
	State           int         `json:"state"           description:"状态 1:生效 2:失效"`
	UseCount        int         `json:"useCount"        description:"使用次数"`
	CreateTime      int64       `json:"createTime"      description:"创建时间"`
	UpdateTime      int64       `json:"updateTime"      description:"更新时间"`
}

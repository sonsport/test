// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceStatisticsUser is the golang structure for table balance_statistics_user.
type BalanceStatisticsUser struct {
	Id            uint64      `json:"id"            description:""`
	UserId        uint64      `json:"userId"        description:""`
	Date          string      `json:"date"          description:"日期"`
	DepletionType uint        `json:"depletionType" description:"消费类型"`
	Type          uint        `json:"type"          description:"消费类型 1支出 2收入"`
	Diamonds      uint64      `json:"diamonds"      description:"钻石变动数额"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PrizeCashbackConfig is the golang structure for table prize_cashback_config.
type PrizeCashbackConfig struct {
	Id                     int64  `json:"id"                     description:"主键id"`
	Name                   string `json:"name"                   description:"名称"`
	CashbackRatio          int    `json:"cashbackRatio"          description:"返现比例"`
	FiveTimesNumber        int    `json:"fiveTimesNumber"        description:"5倍数量"`
	TenTimesNumber         int    `json:"tenTimesNumber"         description:"10倍数量"`
	FiftyTimesNumber       int    `json:"fiftyTimesNumber"       description:"50倍数量"`
	OneHundredTimesNumber  int    `json:"oneHundredTimesNumber"  description:"100倍数量"`
	FiveHundredTimesNumber int    `json:"fiveHundredTimesNumber" description:"500倍"`
	CreateTime             int64  `json:"createTime"             description:"创建时间"`
	UpdateTime             int64  `json:"updateTime"             description:"更新时间"`
}

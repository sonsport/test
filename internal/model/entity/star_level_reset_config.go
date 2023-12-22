// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// StarLevelResetConfig is the golang structure for table star_level_reset_config.
type StarLevelResetConfig struct {
	Id                uint  `json:"id"                description:"自增id"`
	StarLevel         int   `json:"starLevel"         description:"主播星级"`
	RuleType          int   `json:"ruleType"          description:"类别 0保级 1升级 2降级"`
	IncomeMin         int64 `json:"incomeMin"         description:"金币min 不用扩大100"`
	SendGiftPersonMin int   `json:"sendGiftPersonMin" description:"送礼人数min"`
	SelfSendRateMax   int   `json:"selfSendRateMax"   description:"自己送自己金币占比max 0.4 存储40 扩大100倍"`
	CreateTime        int64 `json:"createTime"        description:"创建时间"`
	UpdateTime        int64 `json:"updateTime"        description:"更新时间"`
}

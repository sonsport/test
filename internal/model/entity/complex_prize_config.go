// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ComplexPrizeConfig is the golang structure for table complex_prize_config.
type ComplexPrizeConfig struct {
	Id           int    `json:"id"           description:""`
	ActivityType int    `json:"activityType" description:"抽奖活动类型 1 圣诞 2 元旦"`
	Icon         string `json:"icon"         description:"图标"`
	DrawType     int    `json:"drawType"     description:"抽奖类型 1装扮 2钻石 3优惠券 4 实物"`
	LinkId       string `json:"linkId"       description:"关联商品id  1,2  逗号分割"`
	PrizeName    string `json:"prizeName"    description:"奖品名称"`
	Num          int    `json:"num"          description:"数量"`
	IsFirstPool  int    `json:"isFirstPool"  description:"是否第一个奖池出现"`
	Rate         int    `json:"rate"         description:"概率万分比  *10000的整数"`
	Status       int    `json:"status"       description:"配置状态 0失效 1生效"`
}

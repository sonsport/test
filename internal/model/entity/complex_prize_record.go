// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ComplexPrizeRecord is the golang structure for table complex_prize_record.
type ComplexPrizeRecord struct {
	Id           int64  `json:"id"           description:""`
	UserId       int64  `json:"userId"       description:"用户id"`
	CId          int    `json:"cId"          description:"抽奖配置id"`
	ActivityType int    `json:"activityType" description:"抽奖活动类型 1 圣诞 2 元旦"`
	Icon         string `json:"icon"         description:"图标"`
	DrawType     string `json:"drawType"     description:"抽奖类型 1装扮 2钻石 3优惠券 4 实物"`
	PropsLogId   int64  `json:"propsLogId"   description:"扣除道具流水id"`
	LinkId       string `json:"linkId"       description:"关联商品id  1,2  逗号分割"`
	PrizeName    string `json:"prizeName"    description:"奖品名称"`
	PrizeCount   int    `json:"prizeCount"   description:"奖池期数"`
	CreateTime   int64  `json:"createTime"   description:"创建时间"`
}

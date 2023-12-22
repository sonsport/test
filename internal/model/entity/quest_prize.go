// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QuestPrize is the golang structure for table quest_prize.
type QuestPrize struct {
	Id         uint64 `json:"id"         description:"id"`
	QuestId    int64  `json:"questId"    description:"任务id"`
	PrizeName  string `json:"prizeName"  description:"奖品名称"`
	PrizeIcon  string `json:"prizeIcon"  description:"奖品图标"`
	PrizeType  int    `json:"prizeType"  description:"奖品类型 1装扮 2钻石 3优惠券 4贝壳 5 活跃值 6 用户经验 7 主播经验 8 主播收益"`
	LinkId     string `json:"linkId"     description:"关联商品id  1,2  逗号分割"`
	PrizeNum   int    `json:"prizeNum"   description:"奖品数量"`
	Status     int64  `json:"status"     description:"配置状态 0失效 1生效"`
	CreateTime int64  `json:"createTime" description:"createdAt"`
	UpdateTime int64  `json:"updateTime" description:"updatedAt"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GiftStatistics is the golang structure for table gift_statistics.
type GiftStatistics struct {
	Id            uint64 `json:"id"            description:""`
	GiftId        uint64 `json:"giftId"        description:"礼物id"`
	GiftNum       uint   `json:"giftNum"       description:"礼物个数"`
	TotalDiamonds uint64 `json:"totalDiamonds" description:"礼物总价"`
	BeginTime     int64  `json:"beginTime"     description:"开始时间"`
	EndTime       int64  `json:"endTime"       description:"结束时间"`
	TimeLevel     int    `json:"timeLevel"     description:"时间维度：1.每小时，2.每天"`
	LuckyDiamonds int64  `json:"luckyDiamonds" description:"中奖钻石"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
	UserCount     int64  `json:"userCount"     description:"送礼人数"`
	RoomCount     int64  `json:"roomCount"     description:"房间数"`
}

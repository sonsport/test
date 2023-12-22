// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GiftLuckyLog is the golang structure for table gift_lucky_log.
type GiftLuckyLog struct {
	Id            uint64 `json:"id"            description:""`
	SenderId      uint64 `json:"senderId"      description:"送礼者uid"`
	ReceiverId    uint64 `json:"receiverId"    description:"收礼者uid"`
	GiftId        uint64 `json:"giftId"        description:"礼物id"`
	GiftType      uint   `json:"giftType"      description:"礼物类型"`
	GiftNum       uint   `json:"giftNum"       description:"礼物个数"`
	RoomId        string `json:"roomId"        description:"房间id"`
	Diamonds      uint   `json:"diamonds"      description:"礼物单价"`
	TotalDiamonds uint   `json:"totalDiamonds" description:"礼物总价"`
	LuckyMultiple int64  `json:"luckyMultiple" description:"中奖倍速"`
	LuckyDiamonds int64  `json:"luckyDiamonds" description:"中奖钻石"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

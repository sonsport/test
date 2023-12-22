// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GiftLuckyLog is the golang structure of table gift_lucky_log for DAO operations like Where/Data.
type GiftLuckyLog struct {
	g.Meta        `orm:"table:gift_lucky_log, do:true"`
	Id            interface{} //
	SenderId      interface{} // 送礼者uid
	ReceiverId    interface{} // 收礼者uid
	GiftId        interface{} // 礼物id
	GiftType      interface{} // 礼物类型
	GiftNum       interface{} // 礼物个数
	RoomId        interface{} // 房间id
	Diamonds      interface{} // 礼物单价
	TotalDiamonds interface{} // 礼物总价
	LuckyMultiple interface{} // 中奖倍速
	LuckyDiamonds interface{} // 中奖钻石
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

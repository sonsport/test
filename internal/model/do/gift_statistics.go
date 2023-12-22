// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GiftStatistics is the golang structure of table gift_statistics for DAO operations like Where/Data.
type GiftStatistics struct {
	g.Meta        `orm:"table:gift_statistics, do:true"`
	Id            interface{} //
	GiftId        interface{} // 礼物id
	GiftNum       interface{} // 礼物个数
	TotalDiamonds interface{} // 礼物总价
	BeginTime     interface{} // 开始时间
	EndTime       interface{} // 结束时间
	TimeLevel     interface{} // 时间维度：1.每小时，2.每天
	LuckyDiamonds interface{} // 中奖钻石
	CreateTime    interface{} //
	UpdateTime    interface{} //
	UserCount     interface{} // 送礼人数
	RoomCount     interface{} // 房间数
}

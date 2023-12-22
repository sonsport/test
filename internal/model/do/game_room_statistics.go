// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameRoomStatistics is the golang structure of table game_room_statistics for DAO operations like Where/Data.
type GameRoomStatistics struct {
	g.Meta       `orm:"table:game_room_statistics, do:true"`
	Id           interface{} //
	BeginTime    interface{} // 开始时间
	EndTime      interface{} // 结束时间
	TimeLevel    interface{} // 时间维度：1.每小时，2.每天
	RoomId       interface{} // 直播间id
	Diamonds     interface{} // 总钻石
	GameId       interface{} // 游戏id
	AnchorReward interface{} // 主播收益
	PlayNum      interface{} // 操作数
	PlayPerson   interface{} // 游戏人数
	AnchorId     interface{} // 主播id
	MasterId     interface{} // 公会长id
	ServerId     interface{} // 运营号id
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameStatistics is the golang structure of table game_statistics for DAO operations like Where/Data.
type GameStatistics struct {
	g.Meta        `orm:"table:game_statistics, do:true"`
	Id            interface{} //
	AppChannel    interface{} // app渠道
	GameKey       interface{} // 游戏key
	GameIncome    interface{} // 游戏收益
	GameOutlay    interface{} // 游戏支出
	PlayUserCount interface{} // 游戏人数
	RoomCount     interface{} // 房间数
	PlayCount     interface{} // 操作数
	BeginTime     interface{} // 开始时间
	EndTime       interface{} // 结束时间
	TimeLevel     interface{} // 时间维度：1.每小时，2.每天
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

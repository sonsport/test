// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorGameLog is the golang structure of table anchor_game_log for DAO operations like Where/Data.
type AnchorGameLog struct {
	g.Meta         `orm:"table:anchor_game_log, do:true"`
	Id             interface{} // ID
	AnchorId       interface{} // 主播ID
	UserId         interface{} // 用户ID
	RoomId         interface{} // 第三方直播间ID全局唯一
	GameId         interface{} // 直播间配置的游戏id
	Diamonds       interface{} // 钻石数
	AnchorRate     interface{} // 主播分成  *10000的整数 最小为万分之一
	BeforeDiamonds interface{} // 当前钻石数
	LogType        interface{} // 流水类型 1 增加 2 扣除
	AfterDiamonds  interface{} // 增加之后的钻石数
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

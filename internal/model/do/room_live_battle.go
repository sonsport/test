// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLiveBattle is the golang structure of table room_live_battle for DAO operations like Where/Data.
type RoomLiveBattle struct {
	g.Meta              `orm:"table:room_live_battle, do:true"`
	Id                  interface{} // 直播间Id
	UserId              interface{} // 主播ID
	BattleType          interface{} // pk类型 0 随机 1 指定 2后台拉起
	StreamId            interface{} // 推拉流ID全局唯一
	RoomId              interface{} // 房间ID
	DiamondsCount       interface{} // 打赏钻石数量
	ScoreCount          interface{} // 打赏积分数量
	TargetUserId        interface{} // 对手主播ID
	TargetStreamId      interface{} // 对手stream_id
	TargetDiamondsCount interface{} // 对手打赏钻石数量
	TargetScoreCount    interface{} // 对手打赏积分数量
	TargetRoomId        interface{} // 房间ID
	BattleStatus        interface{} // 0 匹配中 1 pk中 2 邀请人pk取消 3 对方拒绝 4 惩罚中 5 pk正常结束 6 pk异常结束
	EndType             interface{} // 结束原因 0 正常结束 1 超时结束 2 手动结束 3 系统结束
	BeginTime           interface{} // pk开始时间
	EndTime             interface{} // pk结束时间
	UserAnswerTime      interface{} // 用户响应时间
	TargetAnswerTime    interface{} // 对方响应时间
	CreateTime          interface{} //
	UpdateTime          interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserBattleMatch is the golang structure of table user_battle_match for DAO operations like Where/Data.
type UserBattleMatch struct {
	g.Meta             `orm:"table:user_battle_match, do:true"`
	Id                 interface{} // 主键ID
	UserId             interface{} // 用户id
	MatchSeason        interface{} // 比赛赛季 递增
	MatchScore         interface{} // 赛季分数
	MatchMaxScore      interface{} // 赛季最高分数
	WinBattleCount     interface{} // 赢场次
	DogFallBattleCount interface{} // 平局场次
	AllBattleCount     interface{} // 所有pk赛场次
	WinStreakCount     interface{} // 连胜场次（中断归零）
	GradingIsComplete  interface{} // 是否完成定级
	GradingCount       interface{} // 已打定级赛数量
	WinGradingCount    interface{} // 已赢定级赛数量
	GradingScore       interface{} // 定级赛分数
	CreateTime         interface{} //
	UpdateTime         interface{} //
}

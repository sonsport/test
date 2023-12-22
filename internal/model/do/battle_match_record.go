// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BattleMatchRecord is the golang structure of table battle_match_record for DAO operations like Where/Data.
type BattleMatchRecord struct {
	g.Meta               `orm:"table:battle_match_record, do:true"`
	Id                   interface{} // 主键ID
	BattleId             interface{} // pk id
	WinUserId            interface{} // 赢方用户id
	HangUpUserId         interface{} // 挂断人id
	MatchSeason          interface{} // 比赛赛季 递增
	WinMatchScore        interface{} // 赢方分数
	LoseMatchScore       interface{} // 输方分数
	UserInterruptScore   interface{} // 用户中断分数
	WinStreakScore       interface{} // 用户连胜分数
	WinSegmentDiffScore  interface{} // 赢方用户段位差分数
	DoubleScore          interface{} // 翻倍卡分数
	FatiguedScore        interface{} // 疲劳扣除分数
	LoseSegmentDiffScore interface{} // 输方用户段位差分数
	LoseScoreGuard       interface{} // 积分保护
	StreakCountGuard     interface{} // 连胜保护
	RankGuard            interface{} // 段位保护
	WinCurrMatchScore    interface{} // 赢方用户当前分数
	LoseCurrMatchScore   interface{} // 输方用户当前分数
	CreateTime           interface{} //
	UpdateTime           interface{} //
}

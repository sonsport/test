// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserBattleMatch is the golang structure for table user_battle_match.
type UserBattleMatch struct {
	Id                 uint64 `json:"id"                 description:"主键ID"`
	UserId             int64  `json:"userId"             description:"用户id"`
	MatchSeason        int    `json:"matchSeason"        description:"比赛赛季 递增"`
	MatchScore         int    `json:"matchScore"         description:"赛季分数"`
	MatchMaxScore      int    `json:"matchMaxScore"      description:"赛季最高分数"`
	WinBattleCount     int    `json:"winBattleCount"     description:"赢场次"`
	DogFallBattleCount int    `json:"dogFallBattleCount" description:"平局场次"`
	AllBattleCount     int    `json:"allBattleCount"     description:"所有pk赛场次"`
	WinStreakCount     int    `json:"winStreakCount"     description:"连胜场次（中断归零）"`
	GradingIsComplete  int    `json:"gradingIsComplete"  description:"是否完成定级"`
	GradingCount       int    `json:"gradingCount"       description:"已打定级赛数量"`
	WinGradingCount    int    `json:"winGradingCount"    description:"已赢定级赛数量"`
	GradingScore       int    `json:"gradingScore"       description:"定级赛分数"`
	CreateTime         int64  `json:"createTime"         description:""`
	UpdateTime         int64  `json:"updateTime"         description:""`
}

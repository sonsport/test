// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BattleMatchRecord is the golang structure for table battle_match_record.
type BattleMatchRecord struct {
	Id                   uint64 `json:"id"                   description:"主键ID"`
	BattleId             int64  `json:"battleId"             description:"pk id"`
	WinUserId            int64  `json:"winUserId"            description:"赢方用户id"`
	HangUpUserId         int64  `json:"hangUpUserId"         description:"挂断人id"`
	MatchSeason          int    `json:"matchSeason"          description:"比赛赛季 递增"`
	WinMatchScore        int    `json:"winMatchScore"        description:"赢方分数"`
	LoseMatchScore       int    `json:"loseMatchScore"       description:"输方分数"`
	UserInterruptScore   int    `json:"userInterruptScore"   description:"用户中断分数"`
	WinStreakScore       int    `json:"winStreakScore"       description:"用户连胜分数"`
	WinSegmentDiffScore  int    `json:"winSegmentDiffScore"  description:"赢方用户段位差分数"`
	DoubleScore          int    `json:"doubleScore"          description:"翻倍卡分数"`
	FatiguedScore        int    `json:"fatiguedScore"        description:"疲劳扣除分数"`
	LoseSegmentDiffScore int    `json:"loseSegmentDiffScore" description:"输方用户段位差分数"`
	LoseScoreGuard       int    `json:"loseScoreGuard"       description:"积分保护"`
	StreakCountGuard     int    `json:"streakCountGuard"     description:"连胜保护"`
	RankGuard            int    `json:"rankGuard"            description:"段位保护"`
	WinCurrMatchScore    int64  `json:"winCurrMatchScore"    description:"赢方用户当前分数"`
	LoseCurrMatchScore   int64  `json:"loseCurrMatchScore"   description:"输方用户当前分数"`
	CreateTime           int64  `json:"createTime"           description:""`
	UpdateTime           int64  `json:"updateTime"           description:""`
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BattleMatchRecordDao is the data access object for table battle_match_record.
type BattleMatchRecordDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns BattleMatchRecordColumns // columns contains all the column names of Table for convenient usage.
}

// BattleMatchRecordColumns defines and stores column names for table battle_match_record.
type BattleMatchRecordColumns struct {
	Id                   string // 主键ID
	BattleId             string // pk id
	WinUserId            string // 赢方用户id
	HangUpUserId         string // 挂断人id
	MatchSeason          string // 比赛赛季 递增
	WinMatchScore        string // 赢方分数
	LoseMatchScore       string // 输方分数
	UserInterruptScore   string // 用户中断分数
	WinStreakScore       string // 用户连胜分数
	WinSegmentDiffScore  string // 赢方用户段位差分数
	DoubleScore          string // 翻倍卡分数
	FatiguedScore        string // 疲劳扣除分数
	LoseSegmentDiffScore string // 输方用户段位差分数
	LoseScoreGuard       string // 积分保护
	StreakCountGuard     string // 连胜保护
	RankGuard            string // 段位保护
	WinCurrMatchScore    string // 赢方用户当前分数
	LoseCurrMatchScore   string // 输方用户当前分数
	CreateTime           string //
	UpdateTime           string //
}

// battleMatchRecordColumns holds the columns for table battle_match_record.
var battleMatchRecordColumns = BattleMatchRecordColumns{
	Id:                   "id",
	BattleId:             "battle_id",
	WinUserId:            "win_user_id",
	HangUpUserId:         "hang_up_user_id",
	MatchSeason:          "match_season",
	WinMatchScore:        "win_match_score",
	LoseMatchScore:       "lose_match_score",
	UserInterruptScore:   "user_interrupt_score",
	WinStreakScore:       "win_streak_score",
	WinSegmentDiffScore:  "win_segment_diff_score",
	DoubleScore:          "double_score",
	FatiguedScore:        "fatigued_score",
	LoseSegmentDiffScore: "lose_segment_diff_score",
	LoseScoreGuard:       "lose_score_guard",
	StreakCountGuard:     "streak_count_guard",
	RankGuard:            "rank_guard",
	WinCurrMatchScore:    "win_curr_match_score",
	LoseCurrMatchScore:   "lose_curr_match_score",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
}

// NewBattleMatchRecordDao creates and returns a new DAO object for table data access.
func NewBattleMatchRecordDao() *BattleMatchRecordDao {
	return &BattleMatchRecordDao{
		group:   "default",
		table:   "battle_match_record",
		columns: battleMatchRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BattleMatchRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BattleMatchRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BattleMatchRecordDao) Columns() BattleMatchRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BattleMatchRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BattleMatchRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BattleMatchRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBattleMatchDao is the data access object for table user_battle_match.
type UserBattleMatchDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns UserBattleMatchColumns // columns contains all the column names of Table for convenient usage.
}

// UserBattleMatchColumns defines and stores column names for table user_battle_match.
type UserBattleMatchColumns struct {
	Id                 string // 主键ID
	UserId             string // 用户id
	MatchSeason        string // 比赛赛季 递增
	MatchScore         string // 赛季分数
	MatchMaxScore      string // 赛季最高分数
	WinBattleCount     string // 赢场次
	DogFallBattleCount string // 平局场次
	AllBattleCount     string // 所有pk赛场次
	WinStreakCount     string // 连胜场次（中断归零）
	GradingIsComplete  string // 是否完成定级
	GradingCount       string // 已打定级赛数量
	WinGradingCount    string // 已赢定级赛数量
	GradingScore       string // 定级赛分数
	CreateTime         string //
	UpdateTime         string //
}

// userBattleMatchColumns holds the columns for table user_battle_match.
var userBattleMatchColumns = UserBattleMatchColumns{
	Id:                 "id",
	UserId:             "user_id",
	MatchSeason:        "match_season",
	MatchScore:         "match_score",
	MatchMaxScore:      "match_max_score",
	WinBattleCount:     "win_battle_count",
	DogFallBattleCount: "dog_fall_battle_count",
	AllBattleCount:     "all_battle_count",
	WinStreakCount:     "win_streak_count",
	GradingIsComplete:  "grading_is_complete",
	GradingCount:       "grading_count",
	WinGradingCount:    "win_grading_count",
	GradingScore:       "grading_score",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewUserBattleMatchDao creates and returns a new DAO object for table data access.
func NewUserBattleMatchDao() *UserBattleMatchDao {
	return &UserBattleMatchDao{
		group:   "default",
		table:   "user_battle_match",
		columns: userBattleMatchColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserBattleMatchDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserBattleMatchDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserBattleMatchDao) Columns() UserBattleMatchColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserBattleMatchDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserBattleMatchDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserBattleMatchDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

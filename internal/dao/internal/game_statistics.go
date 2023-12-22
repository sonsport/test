// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameStatisticsDao is the data access object for table game_statistics.
type GameStatisticsDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns GameStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// GameStatisticsColumns defines and stores column names for table game_statistics.
type GameStatisticsColumns struct {
	Id            string //
	AppChannel    string // app渠道
	GameKey       string // 游戏key
	GameIncome    string // 游戏收益
	GameOutlay    string // 游戏支出
	PlayUserCount string // 游戏人数
	RoomCount     string // 房间数
	PlayCount     string // 操作数
	BeginTime     string // 开始时间
	EndTime       string // 结束时间
	TimeLevel     string // 时间维度：1.每小时，2.每天
	CreateTime    string //
	UpdateTime    string //
}

// gameStatisticsColumns holds the columns for table game_statistics.
var gameStatisticsColumns = GameStatisticsColumns{
	Id:            "id",
	AppChannel:    "app_channel",
	GameKey:       "game_key",
	GameIncome:    "game_income",
	GameOutlay:    "game_outlay",
	PlayUserCount: "play_user_count",
	RoomCount:     "room_count",
	PlayCount:     "play_count",
	BeginTime:     "begin_time",
	EndTime:       "end_time",
	TimeLevel:     "time_level",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewGameStatisticsDao creates and returns a new DAO object for table data access.
func NewGameStatisticsDao() *GameStatisticsDao {
	return &GameStatisticsDao{
		group:   "default",
		table:   "game_statistics",
		columns: gameStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameStatisticsDao) Columns() GameStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

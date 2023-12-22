// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameRoomStatisticsDao is the data access object for table game_room_statistics.
type GameRoomStatisticsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns GameRoomStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// GameRoomStatisticsColumns defines and stores column names for table game_room_statistics.
type GameRoomStatisticsColumns struct {
	Id           string //
	BeginTime    string // 开始时间
	EndTime      string // 结束时间
	TimeLevel    string // 时间维度：1.每小时，2.每天
	RoomId       string // 直播间id
	Diamonds     string // 总钻石
	GameId       string // 游戏id
	AnchorReward string // 主播收益
	PlayNum      string // 操作数
	PlayPerson   string // 游戏人数
	AnchorId     string // 主播id
	MasterId     string // 公会长id
	ServerId     string // 运营号id
	CreateTime   string //
	UpdateTime   string //
}

// gameRoomStatisticsColumns holds the columns for table game_room_statistics.
var gameRoomStatisticsColumns = GameRoomStatisticsColumns{
	Id:           "id",
	BeginTime:    "begin_time",
	EndTime:      "end_time",
	TimeLevel:    "time_level",
	RoomId:       "room_id",
	Diamonds:     "diamonds",
	GameId:       "game_id",
	AnchorReward: "anchor_reward",
	PlayNum:      "play_num",
	PlayPerson:   "play_person",
	AnchorId:     "anchor_id",
	MasterId:     "master_id",
	ServerId:     "server_id",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewGameRoomStatisticsDao creates and returns a new DAO object for table data access.
func NewGameRoomStatisticsDao() *GameRoomStatisticsDao {
	return &GameRoomStatisticsDao{
		group:   "default",
		table:   "game_room_statistics",
		columns: gameRoomStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameRoomStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameRoomStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameRoomStatisticsDao) Columns() GameRoomStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameRoomStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameRoomStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameRoomStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

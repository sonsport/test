// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLiveBattleDao is the data access object for table room_live_battle.
type RoomLiveBattleDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RoomLiveBattleColumns // columns contains all the column names of Table for convenient usage.
}

// RoomLiveBattleColumns defines and stores column names for table room_live_battle.
type RoomLiveBattleColumns struct {
	Id                  string // 直播间Id
	UserId              string // 主播ID
	BattleType          string // pk类型 0 随机 1 指定 2后台拉起
	StreamId            string // 推拉流ID全局唯一
	RoomId              string // 房间ID
	DiamondsCount       string // 打赏钻石数量
	ScoreCount          string // 打赏积分数量
	TargetUserId        string // 对手主播ID
	TargetStreamId      string // 对手stream_id
	TargetDiamondsCount string // 对手打赏钻石数量
	TargetScoreCount    string // 对手打赏积分数量
	TargetRoomId        string // 房间ID
	BattleStatus        string // 0 匹配中 1 pk中 2 邀请人pk取消 3 对方拒绝 4 惩罚中 5 pk正常结束 6 pk异常结束
	EndType             string // 结束原因 0 正常结束 1 超时结束 2 手动结束 3 系统结束
	BeginTime           string // pk开始时间
	EndTime             string // pk结束时间
	UserAnswerTime      string // 用户响应时间
	TargetAnswerTime    string // 对方响应时间
	CreateTime          string //
	UpdateTime          string //
}

// roomLiveBattleColumns holds the columns for table room_live_battle.
var roomLiveBattleColumns = RoomLiveBattleColumns{
	Id:                  "id",
	UserId:              "user_id",
	BattleType:          "battle_type",
	StreamId:            "stream_id",
	RoomId:              "room_id",
	DiamondsCount:       "diamonds_count",
	ScoreCount:          "score_count",
	TargetUserId:        "target_user_id",
	TargetStreamId:      "target_stream_id",
	TargetDiamondsCount: "target_diamonds_count",
	TargetScoreCount:    "target_score_count",
	TargetRoomId:        "target_room_id",
	BattleStatus:        "battle_status",
	EndType:             "end_type",
	BeginTime:           "begin_time",
	EndTime:             "end_time",
	UserAnswerTime:      "user_answer_time",
	TargetAnswerTime:    "target_answer_time",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewRoomLiveBattleDao creates and returns a new DAO object for table data access.
func NewRoomLiveBattleDao() *RoomLiveBattleDao {
	return &RoomLiveBattleDao{
		group:   "default",
		table:   "room_live_battle",
		columns: roomLiveBattleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomLiveBattleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomLiveBattleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomLiveBattleDao) Columns() RoomLiveBattleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomLiveBattleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomLiveBattleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomLiveBattleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

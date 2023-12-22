// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorGameLogDao is the data access object for table anchor_game_log.
type AnchorGameLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AnchorGameLogColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorGameLogColumns defines and stores column names for table anchor_game_log.
type AnchorGameLogColumns struct {
	Id             string // ID
	AnchorId       string // 主播ID
	UserId         string // 用户ID
	RoomId         string // 第三方直播间ID全局唯一
	GameId         string // 直播间配置的游戏id
	Diamonds       string // 钻石数
	AnchorRate     string // 主播分成  *10000的整数 最小为万分之一
	BeforeDiamonds string // 当前钻石数
	LogType        string // 流水类型 1 增加 2 扣除
	AfterDiamonds  string // 增加之后的钻石数
	CreateTime     string //
	UpdateTime     string //
}

// anchorGameLogColumns holds the columns for table anchor_game_log.
var anchorGameLogColumns = AnchorGameLogColumns{
	Id:             "id",
	AnchorId:       "anchor_id",
	UserId:         "user_id",
	RoomId:         "room_id",
	GameId:         "game_id",
	Diamonds:       "diamonds",
	AnchorRate:     "anchor_rate",
	BeforeDiamonds: "before_diamonds",
	LogType:        "log_type",
	AfterDiamonds:  "after_diamonds",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewAnchorGameLogDao creates and returns a new DAO object for table data access.
func NewAnchorGameLogDao() *AnchorGameLogDao {
	return &AnchorGameLogDao{
		group:   "default",
		table:   "anchor_game_log",
		columns: anchorGameLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorGameLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorGameLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorGameLogDao) Columns() AnchorGameLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorGameLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorGameLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorGameLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

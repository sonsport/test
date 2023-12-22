// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableStatisticsDao is the data access object for table diamond_turntable_statistics.
type DiamondTurntableStatisticsDao struct {
	table   string                            // table is the underlying table name of the DAO.
	group   string                            // group is the database configuration group name of current DAO.
	columns DiamondTurntableStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// DiamondTurntableStatisticsColumns defines and stores column names for table diamond_turntable_statistics.
type DiamondTurntableStatisticsColumns struct {
	Id             string //
	UserId         string // 用户id
	WatchTime      string // 在直播间的有效观看时间 单位s
	RaffleNum      string // 获得的抽奖次数
	UseNum         string // 使用的抽奖次数
	Gid            string // 商品表id
	Num            string // 数量
	TurntableCount string // 奖池期数
	CreateTime     string //
	UpdateTime     string //
	Cid            string // 抽奖配置id
	PrizeName      string // 奖品名称
}

// diamondTurntableStatisticsColumns holds the columns for table diamond_turntable_statistics.
var diamondTurntableStatisticsColumns = DiamondTurntableStatisticsColumns{
	Id:             "id",
	UserId:         "user_id",
	WatchTime:      "watch_time",
	RaffleNum:      "raffle_num",
	UseNum:         "use_num",
	Gid:            "gid",
	Num:            "num",
	TurntableCount: "turntable_count",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	Cid:            "cid",
	PrizeName:      "prize_name",
}

// NewDiamondTurntableStatisticsDao creates and returns a new DAO object for table data access.
func NewDiamondTurntableStatisticsDao() *DiamondTurntableStatisticsDao {
	return &DiamondTurntableStatisticsDao{
		group:   "default",
		table:   "diamond_turntable_statistics",
		columns: diamondTurntableStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DiamondTurntableStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DiamondTurntableStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DiamondTurntableStatisticsDao) Columns() DiamondTurntableStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DiamondTurntableStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DiamondTurntableStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DiamondTurntableStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

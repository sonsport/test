// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PrizePoolConfigDao is the data access object for table prize_pool_config.
type PrizePoolConfigDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns PrizePoolConfigColumns // columns contains all the column names of Table for convenient usage.
}

// PrizePoolConfigColumns defines and stores column names for table prize_pool_config.
type PrizePoolConfigColumns struct {
	Id              string // 主键id
	Gid             string // 幸运礼物id
	StartTime       string // 配置生效时间
	Type            string // 配置类型 1:默认 2:其他
	CashbackRatioId string // 奖池返现配置id
	State           string // 状态 1:生效 2:失效
	UseCount        string // 使用次数
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
}

// prizePoolConfigColumns holds the columns for table prize_pool_config.
var prizePoolConfigColumns = PrizePoolConfigColumns{
	Id:              "id",
	Gid:             "gid",
	StartTime:       "start_time",
	Type:            "type",
	CashbackRatioId: "cashback_ratio_id",
	State:           "state",
	UseCount:        "use_count",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewPrizePoolConfigDao creates and returns a new DAO object for table data access.
func NewPrizePoolConfigDao() *PrizePoolConfigDao {
	return &PrizePoolConfigDao{
		group:   "default",
		table:   "prize_pool_config",
		columns: prizePoolConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PrizePoolConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PrizePoolConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PrizePoolConfigDao) Columns() PrizePoolConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PrizePoolConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PrizePoolConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PrizePoolConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

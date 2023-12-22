// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FirstSettlementChanceDao is the data access object for table first_settlement_chance.
type FirstSettlementChanceDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns FirstSettlementChanceColumns // columns contains all the column names of Table for convenient usage.
}

// FirstSettlementChanceColumns defines and stores column names for table first_settlement_chance.
type FirstSettlementChanceColumns struct {
	Id         string // 自增Id
	Week       string // 周，一年中的第几周 例如:202230 2022年的第30周
	UserId     string // 用户ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// firstSettlementChanceColumns holds the columns for table first_settlement_chance.
var firstSettlementChanceColumns = FirstSettlementChanceColumns{
	Id:         "id",
	Week:       "week",
	UserId:     "user_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewFirstSettlementChanceDao creates and returns a new DAO object for table data access.
func NewFirstSettlementChanceDao() *FirstSettlementChanceDao {
	return &FirstSettlementChanceDao{
		group:   "default",
		table:   "first_settlement_chance",
		columns: firstSettlementChanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FirstSettlementChanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FirstSettlementChanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FirstSettlementChanceDao) Columns() FirstSettlementChanceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FirstSettlementChanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FirstSettlementChanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FirstSettlementChanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

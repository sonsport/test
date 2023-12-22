// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceRecordDao is the data access object for table balance_record.
type BalanceRecordDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns BalanceRecordColumns // columns contains all the column names of Table for convenient usage.
}

// BalanceRecordColumns defines and stores column names for table balance_record.
type BalanceRecordColumns struct {
	Id            string //
	SourceUserId  string // 用户Id
	TargetUserId  string // 主播Id
	DepletionType string // 消费类型
	Diamonds      string // 钻石变动数额
	BeforeChange  string // 变更之前余额
	AfterChange   string // 变更之后余额
	LinkId        string // 引起余额变更的id
	Type          string // 消费类型 1支出 2收入
	CreateTime    string //
	UpdateTime    string //
}

// balanceRecordColumns holds the columns for table balance_record.
var balanceRecordColumns = BalanceRecordColumns{
	Id:            "id",
	SourceUserId:  "source_user_id",
	TargetUserId:  "target_user_id",
	DepletionType: "depletion_type",
	Diamonds:      "diamonds",
	BeforeChange:  "before_change",
	AfterChange:   "after_change",
	LinkId:        "link_id",
	Type:          "type",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewBalanceRecordDao creates and returns a new DAO object for table data access.
func NewBalanceRecordDao() *BalanceRecordDao {
	return &BalanceRecordDao{
		group:   "default",
		table:   "balance_record",
		columns: balanceRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BalanceRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BalanceRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BalanceRecordDao) Columns() BalanceRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BalanceRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BalanceRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BalanceRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

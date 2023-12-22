// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// IncomeRecordDao is the data access object for table income_record.
type IncomeRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns IncomeRecordColumns // columns contains all the column names of Table for convenient usage.
}

// IncomeRecordColumns defines and stores column names for table income_record.
type IncomeRecordColumns struct {
	Id           string // id
	UserId       string // 用户id
	SourceId     string // 收益来源用户的id
	UserRecordId string // 对应用户扣费记录id
	IncomeType   string // 收益类型
	LinkId       string // 引起变更的id
	Income       string // 收益变动数额
	BeforeChange string // 变更之前余额
	AfterChange  string // 变更之后余额
	CreateTime   string // createdAt
	UpdateTime   string // updatedAt
}

// incomeRecordColumns holds the columns for table income_record.
var incomeRecordColumns = IncomeRecordColumns{
	Id:           "id",
	UserId:       "user_id",
	SourceId:     "source_id",
	UserRecordId: "user_record_id",
	IncomeType:   "income_type",
	LinkId:       "link_id",
	Income:       "income",
	BeforeChange: "before_change",
	AfterChange:  "after_change",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewIncomeRecordDao creates and returns a new DAO object for table data access.
func NewIncomeRecordDao() *IncomeRecordDao {
	return &IncomeRecordDao{
		group:   "default",
		table:   "income_record",
		columns: incomeRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *IncomeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *IncomeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *IncomeRecordDao) Columns() IncomeRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *IncomeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *IncomeRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *IncomeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

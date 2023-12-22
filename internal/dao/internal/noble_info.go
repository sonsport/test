// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NobleInfoDao is the data access object for table noble_info.
type NobleInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns NobleInfoColumns // columns contains all the column names of Table for convenient usage.
}

// NobleInfoColumns defines and stores column names for table noble_info.
type NobleInfoColumns struct {
	Id            string //
	NobleName     string // 贵族名称
	NobleStates   string // 贵族状态 1 上架
	NobleIcon     string // 贵族图标
	CurrentPrice  string // 现价
	OriginalPrice string // 原价
	NobleType     string // 贵族类型 1 vip
	NobleLevel    string // 贵族等级
	NobleDay      string // 贵族天数
	CreateTime    string //
	UpdateTime    string //
}

// nobleInfoColumns holds the columns for table noble_info.
var nobleInfoColumns = NobleInfoColumns{
	Id:            "id",
	NobleName:     "noble_name",
	NobleStates:   "noble_states",
	NobleIcon:     "noble_icon",
	CurrentPrice:  "current_price",
	OriginalPrice: "original_price",
	NobleType:     "noble_type",
	NobleLevel:    "noble_level",
	NobleDay:      "noble_day",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewNobleInfoDao creates and returns a new DAO object for table data access.
func NewNobleInfoDao() *NobleInfoDao {
	return &NobleInfoDao{
		group:   "default",
		table:   "noble_info",
		columns: nobleInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NobleInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NobleInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NobleInfoDao) Columns() NobleInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NobleInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NobleInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NobleInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

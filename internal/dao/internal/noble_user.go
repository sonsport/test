// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NobleUserDao is the data access object for table noble_user.
type NobleUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns NobleUserColumns // columns contains all the column names of Table for convenient usage.
}

// NobleUserColumns defines and stores column names for table noble_user.
type NobleUserColumns struct {
	Id                  string //
	UserId              string //
	NobleId             string // 贵族id
	ExpirationTime      string // 过期时间
	NobleStates         string // 贵族状态 1 有效
	PaidThenDiamonds    string // 购买vip时的付费钻石数
	PaidProductDay      string // 购买的规格天
	PaidProductDiamond  string // 购买的钻石数
	RenewInspectionTime string // 自动续费检查时间
	CreateTime          string //
	UpdateTime          string //
}

// nobleUserColumns holds the columns for table noble_user.
var nobleUserColumns = NobleUserColumns{
	Id:                  "id",
	UserId:              "user_id",
	NobleId:             "noble_id",
	ExpirationTime:      "expiration_time",
	NobleStates:         "noble_states",
	PaidThenDiamonds:    "paid_then_diamonds",
	PaidProductDay:      "paid_product_day",
	PaidProductDiamond:  "paid_product_diamond",
	RenewInspectionTime: "renew_inspection_time",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
}

// NewNobleUserDao creates and returns a new DAO object for table data access.
func NewNobleUserDao() *NobleUserDao {
	return &NobleUserDao{
		group:   "default",
		table:   "noble_user",
		columns: nobleUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NobleUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NobleUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NobleUserDao) Columns() NobleUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NobleUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NobleUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NobleUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

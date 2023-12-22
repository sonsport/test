// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdjustMessageDao is the data access object for table adjust_message.
type AdjustMessageDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdjustMessageColumns // columns contains all the column names of Table for convenient usage.
}

// AdjustMessageColumns defines and stores column names for table adjust_message.
type AdjustMessageColumns struct {
	Id          string //
	InstallTime string // install_time
	UserId      string // 用户Id
	UserName    string //
	OrderId     string // 用户Id
	EventToken  string //
	GaId        string // 订单id
	DeviceId    string // 设备id
	AppName     string //
	AppId       string //
	Tracker     string //
	Network     string //
	AdjustId    string //
	Currency    string //
	Revenue     string //
	CreateTime  string //
}

// adjustMessageColumns holds the columns for table adjust_message.
var adjustMessageColumns = AdjustMessageColumns{
	Id:          "id",
	InstallTime: "install_time",
	UserId:      "user_id",
	UserName:    "user_name",
	OrderId:     "order_id",
	EventToken:  "event_token",
	GaId:        "ga_id",
	DeviceId:    "device_id",
	AppName:     "app_name",
	AppId:       "app_id",
	Tracker:     "tracker",
	Network:     "network",
	AdjustId:    "adjust_id",
	Currency:    "currency",
	Revenue:     "revenue",
	CreateTime:  "create_time",
}

// NewAdjustMessageDao creates and returns a new DAO object for table data access.
func NewAdjustMessageDao() *AdjustMessageDao {
	return &AdjustMessageDao{
		group:   "default",
		table:   "adjust_message",
		columns: adjustMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdjustMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdjustMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdjustMessageDao) Columns() AdjustMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdjustMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdjustMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdjustMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

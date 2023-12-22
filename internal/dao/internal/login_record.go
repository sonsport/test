// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LoginRecordDao is the data access object for table login_record.
type LoginRecordDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns LoginRecordColumns // columns contains all the column names of Table for convenient usage.
}

// LoginRecordColumns defines and stores column names for table login_record.
type LoginRecordColumns struct {
	Id         string //
	UserId     string // 用户id
	AppName    string // app名称
	AppChannel string // app渠道
	AppVersion string // app版本
	DeviceId   string // 设备号
	Ip         string // 登陆ip
	Language   string // 语言
	CreateTime string //
	UpdateTime string //
	AppToken   string // app登录token
	LoginType  string // 1 gg 2 fb 3 apple
	AppPhone   string // app手机型号
}

// loginRecordColumns holds the columns for table login_record.
var loginRecordColumns = LoginRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	AppName:    "app_name",
	AppChannel: "app_channel",
	AppVersion: "app_version",
	DeviceId:   "device_id",
	Ip:         "ip",
	Language:   "language",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	AppToken:   "app_token",
	LoginType:  "login_type",
	AppPhone:   "app_phone",
}

// NewLoginRecordDao creates and returns a new DAO object for table data access.
func NewLoginRecordDao() *LoginRecordDao {
	return &LoginRecordDao{
		group:   "default",
		table:   "login_record",
		columns: loginRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LoginRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LoginRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LoginRecordDao) Columns() LoginRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LoginRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LoginRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LoginRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

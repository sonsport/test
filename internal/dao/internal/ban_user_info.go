// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BanUserInfoDao is the data access object for table ban_user_info.
type BanUserInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BanUserInfoColumns // columns contains all the column names of Table for convenient usage.
}

// BanUserInfoColumns defines and stores column names for table ban_user_info.
type BanUserInfoColumns struct {
	Id         string // 自增id
	UserId     string // 封禁用户id
	AdminId    string // 管理员id
	LinkId     string // 关联警告表主键id
	Type       string // 封禁类型，1禁发私信，2禁发公屏聊天，3禁开播，4禁止登录，5封禁设备
	BanTime    string // 封禁时间
	DeviceId   string // 封禁设备
	BanReason  string // 封禁原因
	CreateTime string //
	UpdateTime string //
}

// banUserInfoColumns holds the columns for table ban_user_info.
var banUserInfoColumns = BanUserInfoColumns{
	Id:         "id",
	UserId:     "user_id",
	AdminId:    "admin_id",
	LinkId:     "link_id",
	Type:       "type",
	BanTime:    "ban_time",
	DeviceId:   "device_id",
	BanReason:  "ban_reason",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewBanUserInfoDao creates and returns a new DAO object for table data access.
func NewBanUserInfoDao() *BanUserInfoDao {
	return &BanUserInfoDao{
		group:   "default",
		table:   "ban_user_info",
		columns: banUserInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BanUserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BanUserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BanUserInfoDao) Columns() BanUserInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BanUserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BanUserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BanUserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShareNewUserInfoDao is the data access object for table share_new_user_info.
type ShareNewUserInfoDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns ShareNewUserInfoColumns // columns contains all the column names of Table for convenient usage.
}

// ShareNewUserInfoColumns defines and stores column names for table share_new_user_info.
type ShareNewUserInfoColumns struct {
	Id               string //
	UserId           string //
	ShareId          string // 分享ID
	OwnerUserId      string // 分享者uid
	WatchLiveTime    string // 观看直播时长
	RegisterTime     string // 注册时间
	Ip               string // 登陆地IP
	Smid             string // 设备ID
	DeviceName       string // 设备型号
	Status           string // 0:初始化；1:待加币；2:加币成功
	WatchLiveRewards string // 1:看播奖励有效；2:看播奖励无效
	IsOldUser        string // 1:老用户；0:新用户
	FullReasonCode   string // 失败原因 status -1时有效
	CreateTime       string //
	UpdateTime       string //
}

// shareNewUserInfoColumns holds the columns for table share_new_user_info.
var shareNewUserInfoColumns = ShareNewUserInfoColumns{
	Id:               "id",
	UserId:           "user_id",
	ShareId:          "share_id",
	OwnerUserId:      "owner_user_id",
	WatchLiveTime:    "watch_live_time",
	RegisterTime:     "register_time",
	Ip:               "ip",
	Smid:             "smid",
	DeviceName:       "device_name",
	Status:           "status",
	WatchLiveRewards: "watch_live_rewards",
	IsOldUser:        "is_old_user",
	FullReasonCode:   "full_reason_code",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewShareNewUserInfoDao creates and returns a new DAO object for table data access.
func NewShareNewUserInfoDao() *ShareNewUserInfoDao {
	return &ShareNewUserInfoDao{
		group:   "default",
		table:   "share_new_user_info",
		columns: shareNewUserInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShareNewUserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShareNewUserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShareNewUserInfoDao) Columns() ShareNewUserInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShareNewUserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShareNewUserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShareNewUserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

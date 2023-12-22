// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBaseInfoDao is the data access object for table user_base_info.
type UserBaseInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserBaseInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserBaseInfoColumns defines and stores column names for table user_base_info.
type UserBaseInfoColumns struct {
	Id               string //
	UserId           string // 用户Id
	FirstLiveTime    string // 首播日期
	FirstAuthTime    string // 首次认证日期
	LastAuthTime     string // 最后认证日期
	LiveCallSwitch   string // 连麦开关 1打开
	LiveCallType     string // 连麦类型 1 所有人 2 关注我的人 3我的好友
	ReviewUpUser     string // 多人房上麦审核
	MultipleUserType string // 多人房上麦权限  1 所有人 2 关注我的人 3我的好友
	CreatedTime      string //
	UpdatedTime      string //
	ServerId         string // 运营号id
}

// userBaseInfoColumns holds the columns for table user_base_info.
var userBaseInfoColumns = UserBaseInfoColumns{
	Id:               "id",
	UserId:           "user_id",
	FirstLiveTime:    "first_live_time",
	FirstAuthTime:    "first_auth_time",
	LastAuthTime:     "last_auth_time",
	LiveCallSwitch:   "live_call_switch",
	LiveCallType:     "live_call_type",
	ReviewUpUser:     "review_up_user",
	MultipleUserType: "multiple_user_type",
	CreatedTime:      "created_time",
	UpdatedTime:      "updated_time",
	ServerId:         "server_id",
}

// NewUserBaseInfoDao creates and returns a new DAO object for table data access.
func NewUserBaseInfoDao() *UserBaseInfoDao {
	return &UserBaseInfoDao{
		group:   "default",
		table:   "user_base_info",
		columns: userBaseInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserBaseInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserBaseInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserBaseInfoDao) Columns() UserBaseInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserBaseInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserBaseInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserBaseInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

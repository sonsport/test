// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDressUpLogDao is the data access object for table user_dress_up_log.
type UserDressUpLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserDressUpLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserDressUpLogColumns defines and stores column names for table user_dress_up_log.
type UserDressUpLogColumns struct {
	Id         string //
	UserId     string // 用户Id
	ShopId     string // 商品ID
	CategoryId string // 分类ID
	DressType  string // 0：卸载装扮 1：穿戴装扮
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// userDressUpLogColumns holds the columns for table user_dress_up_log.
var userDressUpLogColumns = UserDressUpLogColumns{
	Id:         "id",
	UserId:     "user_id",
	ShopId:     "shop_id",
	CategoryId: "category_id",
	DressType:  "dress_type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewUserDressUpLogDao creates and returns a new DAO object for table data access.
func NewUserDressUpLogDao() *UserDressUpLogDao {
	return &UserDressUpLogDao{
		group:   "default",
		table:   "user_dress_up_log",
		columns: userDressUpLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDressUpLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDressUpLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDressUpLogDao) Columns() UserDressUpLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDressUpLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDressUpLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDressUpLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

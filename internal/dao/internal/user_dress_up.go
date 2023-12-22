// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDressUpDao is the data access object for table user_dress_up.
type UserDressUpDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns UserDressUpColumns // columns contains all the column names of Table for convenient usage.
}

// UserDressUpColumns defines and stores column names for table user_dress_up.
type UserDressUpColumns struct {
	Id         string //
	UserId     string // 用户Id
	ShopId     string // 商品ID
	CategoryId string // 分类ID
	StartTime  string // 开始时间
	EndTime    string // 结束时间 -1标识永久配置
	Status     string // 状态 0-未佩戴 1-佩戴
	EffectDays string // 有效期 （单位：天）-1标识永久
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// userDressUpColumns holds the columns for table user_dress_up.
var userDressUpColumns = UserDressUpColumns{
	Id:         "id",
	UserId:     "user_id",
	ShopId:     "shop_id",
	CategoryId: "category_id",
	StartTime:  "start_time",
	EndTime:    "end_time",
	Status:     "status",
	EffectDays: "effect_days",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewUserDressUpDao creates and returns a new DAO object for table data access.
func NewUserDressUpDao() *UserDressUpDao {
	return &UserDressUpDao{
		group:   "default",
		table:   "user_dress_up",
		columns: userDressUpColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDressUpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDressUpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDressUpDao) Columns() UserDressUpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDressUpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDressUpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDressUpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

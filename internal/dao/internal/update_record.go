// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateRecordDao is the data access object for table update_record.
type UpdateRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UpdateRecordColumns // columns contains all the column names of Table for convenient usage.
}

// UpdateRecordColumns defines and stores column names for table update_record.
type UpdateRecordColumns struct {
	Id         string // 主键id
	UserId     string // 用户id
	AdminId    string // 管理员id
	UpdateType string // 调整类型，1星级，2公会，3标签，4备注
	BeforeBody string // 修改前内容
	AfterBody  string // 修改后内容
	Remark     string // 修改后内容
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// updateRecordColumns holds the columns for table update_record.
var updateRecordColumns = UpdateRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	AdminId:    "admin_id",
	UpdateType: "update_type",
	BeforeBody: "before_body",
	AfterBody:  "after_body",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewUpdateRecordDao creates and returns a new DAO object for table data access.
func NewUpdateRecordDao() *UpdateRecordDao {
	return &UpdateRecordDao{
		group:   "default",
		table:   "update_record",
		columns: updateRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UpdateRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UpdateRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UpdateRecordDao) Columns() UpdateRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UpdateRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UpdateRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UpdateRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

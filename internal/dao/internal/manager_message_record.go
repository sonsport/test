// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ManagerMessageRecordDao is the data access object for table manager_message_record.
type ManagerMessageRecordDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns ManagerMessageRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ManagerMessageRecordColumns defines and stores column names for table manager_message_record.
type ManagerMessageRecordColumns struct {
	Id         string // 自增id
	UserId     string // 用户id
	CreateTime string //
	UpdateTime string //
	AdminId    string // 管理员id
	Message    string // 发送的消息
	WhatsApp   string // whatsapp号码
	State      string //
}

// managerMessageRecordColumns holds the columns for table manager_message_record.
var managerMessageRecordColumns = ManagerMessageRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	AdminId:    "admin_id",
	Message:    "message",
	WhatsApp:   "whats_app",
	State:      "state",
}

// NewManagerMessageRecordDao creates and returns a new DAO object for table data access.
func NewManagerMessageRecordDao() *ManagerMessageRecordDao {
	return &ManagerMessageRecordDao{
		group:   "default",
		table:   "manager_message_record",
		columns: managerMessageRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ManagerMessageRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ManagerMessageRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ManagerMessageRecordDao) Columns() ManagerMessageRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ManagerMessageRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ManagerMessageRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ManagerMessageRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

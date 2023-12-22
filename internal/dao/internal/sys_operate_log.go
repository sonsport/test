// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOperateLogDao is the data access object for table sys_operate_log.
type SysOperateLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SysOperateLogColumns // columns contains all the column names of Table for convenient usage.
}

// SysOperateLogColumns defines and stores column names for table sys_operate_log.
type SysOperateLogColumns struct {
	Id        string // ID
	AdminId   string // 用户ID
	BeginTime string // 开始时间
	EndTime   string // 结束时间
	Ip        string // 请求地址
	Uri       string // ip
	Params    string // 请求参数
	Method    string // HTTP方法
	Response  string // 返回信息
	Host      string // 请求主机
	Brower    string // 浏览器
}

// sysOperateLogColumns holds the columns for table sys_operate_log.
var sysOperateLogColumns = SysOperateLogColumns{
	Id:        "id",
	AdminId:   "admin_id",
	BeginTime: "begin_time",
	EndTime:   "end_time",
	Ip:        "ip",
	Uri:       "uri",
	Params:    "params",
	Method:    "method",
	Response:  "response",
	Host:      "host",
	Brower:    "brower",
}

// NewSysOperateLogDao creates and returns a new DAO object for table data access.
func NewSysOperateLogDao() *SysOperateLogDao {
	return &SysOperateLogDao{
		group:   "default",
		table:   "sys_operate_log",
		columns: sysOperateLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysOperateLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysOperateLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysOperateLogDao) Columns() SysOperateLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysOperateLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysOperateLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysOperateLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

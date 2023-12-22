// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentOperaRecordDao is the data access object for table agent_opera_record.
type AgentOperaRecordDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns AgentOperaRecordColumns // columns contains all the column names of Table for convenient usage.
}

// AgentOperaRecordColumns defines and stores column names for table agent_opera_record.
type AgentOperaRecordColumns struct {
	Id         string //
	UserId     string //
	OperaIp    string // 操作ip
	LinkId     string // 关联id
	UserAgent  string // 系统类型
	OperaType  string // 操作类型 1 登录 2 转账 3 充值
	CreateTime string //
	UpdateTime string //
}

// agentOperaRecordColumns holds the columns for table agent_opera_record.
var agentOperaRecordColumns = AgentOperaRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	OperaIp:    "opera_ip",
	LinkId:     "link_id",
	UserAgent:  "user_agent",
	OperaType:  "opera_type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAgentOperaRecordDao creates and returns a new DAO object for table data access.
func NewAgentOperaRecordDao() *AgentOperaRecordDao {
	return &AgentOperaRecordDao{
		group:   "default",
		table:   "agent_opera_record",
		columns: agentOperaRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AgentOperaRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AgentOperaRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AgentOperaRecordDao) Columns() AgentOperaRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AgentOperaRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AgentOperaRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AgentOperaRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

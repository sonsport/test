// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentDetailRecordDao is the data access object for table agent_detail_record.
type AgentDetailRecordDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns AgentDetailRecordColumns // columns contains all the column names of Table for convenient usage.
}

// AgentDetailRecordColumns defines and stores column names for table agent_detail_record.
type AgentDetailRecordColumns struct {
	Id            string //
	UserId        string // 用户Id
	OperateUserId string // 操作用户id
	ChangeType    string // 触发类型 1 后台 2 系统
	UpdateType    string // 更新类型 1 代理等级 2 支付密码 3 禁止转账截止时间 4 是否锁定 5 代理状态
	BeforeChange  string // 变更之前值
	AfterChange   string // 变更之后值
	Remarks       string // 更新备注
	CreateTime    string //
	UpdateTime    string //
}

// agentDetailRecordColumns holds the columns for table agent_detail_record.
var agentDetailRecordColumns = AgentDetailRecordColumns{
	Id:            "id",
	UserId:        "user_id",
	OperateUserId: "operate_user_id",
	ChangeType:    "change_type",
	UpdateType:    "update_type",
	BeforeChange:  "before_change",
	AfterChange:   "after_change",
	Remarks:       "remarks",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewAgentDetailRecordDao creates and returns a new DAO object for table data access.
func NewAgentDetailRecordDao() *AgentDetailRecordDao {
	return &AgentDetailRecordDao{
		group:   "default",
		table:   "agent_detail_record",
		columns: agentDetailRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AgentDetailRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AgentDetailRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AgentDetailRecordDao) Columns() AgentDetailRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AgentDetailRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AgentDetailRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AgentDetailRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

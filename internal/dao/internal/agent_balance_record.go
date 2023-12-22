// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentBalanceRecordDao is the data access object for table agent_balance_record.
type AgentBalanceRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AgentBalanceRecordColumns // columns contains all the column names of Table for convenient usage.
}

// AgentBalanceRecordColumns defines and stores column names for table agent_balance_record.
type AgentBalanceRecordColumns struct {
	Id             string //
	UserId         string //
	TargetUserId   string //
	RecordType     string // 消费类型 1支出 2收入
	OnePrice       string // 售出单价
	CostPrice      string // 成本单价
	Diamonds       string // 钻石数额
	BeforeChange   string // 变更之前余额
	AfterChange    string // 变更之后余额
	LinkId         string // 引起余额变更的id
	CertificateUrl string // 凭证url
	PaymentId      string // 支付id
	PaymentName    string // 支付名称
	CreateTime     string //
	UpdateTime     string //
}

// agentBalanceRecordColumns holds the columns for table agent_balance_record.
var agentBalanceRecordColumns = AgentBalanceRecordColumns{
	Id:             "id",
	UserId:         "user_id",
	TargetUserId:   "target_user_id",
	RecordType:     "record_type",
	OnePrice:       "one_price",
	CostPrice:      "cost_price",
	Diamonds:       "diamonds",
	BeforeChange:   "before_change",
	AfterChange:    "after_change",
	LinkId:         "link_id",
	CertificateUrl: "certificate_url",
	PaymentId:      "payment_id",
	PaymentName:    "payment_name",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewAgentBalanceRecordDao creates and returns a new DAO object for table data access.
func NewAgentBalanceRecordDao() *AgentBalanceRecordDao {
	return &AgentBalanceRecordDao{
		group:   "default",
		table:   "agent_balance_record",
		columns: agentBalanceRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AgentBalanceRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AgentBalanceRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AgentBalanceRecordDao) Columns() AgentBalanceRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AgentBalanceRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AgentBalanceRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AgentBalanceRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

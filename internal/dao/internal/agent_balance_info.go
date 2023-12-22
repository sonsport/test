// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentBalanceInfoDao is the data access object for table agent_balance_info.
type AgentBalanceInfoDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns AgentBalanceInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AgentBalanceInfoColumns defines and stores column names for table agent_balance_info.
type AgentBalanceInfoColumns struct {
	UserId            string //
	TotalFee          string // 总充值 印尼盾 单位：分
	TotalSaleFee      string // 总销售 印尼盾 单位：分
	TotalBuyDiamonds  string // 总购买钻石
	TotalSaleDiamonds string // 总销售钻石
	RemainDiamonds    string // 剩余钻石数
	PayPassword       string // 支付密码
	CostPrice         string // 成本单价
	IsLock            string // 是否锁定
	AgentStates       string // 代理状态 1正常
	AgentRank         string // 代理等级 1 2 3
	AgentBanTime      string // 禁止转账截止时间
	CreateTime        string //
	UpdateTime        string //
}

// agentBalanceInfoColumns holds the columns for table agent_balance_info.
var agentBalanceInfoColumns = AgentBalanceInfoColumns{
	UserId:            "user_id",
	TotalFee:          "total_fee",
	TotalSaleFee:      "total_sale_fee",
	TotalBuyDiamonds:  "total_buy_diamonds",
	TotalSaleDiamonds: "total_sale_diamonds",
	RemainDiamonds:    "remain_diamonds",
	PayPassword:       "pay_password",
	CostPrice:         "cost_price",
	IsLock:            "is_lock",
	AgentStates:       "agent_states",
	AgentRank:         "agent_rank",
	AgentBanTime:      "agent_ban_time",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// NewAgentBalanceInfoDao creates and returns a new DAO object for table data access.
func NewAgentBalanceInfoDao() *AgentBalanceInfoDao {
	return &AgentBalanceInfoDao{
		group:   "default",
		table:   "agent_balance_info",
		columns: agentBalanceInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AgentBalanceInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AgentBalanceInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AgentBalanceInfoDao) Columns() AgentBalanceInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AgentBalanceInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AgentBalanceInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AgentBalanceInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

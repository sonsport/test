// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgentGiveawayPropsDao is the data access object for table agent_giveaway_props.
type AgentGiveawayPropsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns AgentGiveawayPropsColumns // columns contains all the column names of Table for convenient usage.
}

// AgentGiveawayPropsColumns defines and stores column names for table agent_giveaway_props.
type AgentGiveawayPropsColumns struct {
	Id              string // 直播间Id
	UserId          string // 主播ID
	PropsName       string // 道具名称
	PropsIcon       string // 道具icon
	PropsDay        string // 道具天数 30
	PropsStates     string // 道具状态 1 待使用 2 已使用
	PropsType       string // 道具类型 1 礼物 2 vip 3 装扮
	PropsLink       string // 道具关联id
	BalanceRecordId string // 流水记录id
	UseTime         string // 使用时间
	TargetUserId    string // 使用目标id
	ExpirationTime  string // 过期时间 -1 不过期
	CreateTime      string //
	UpdateTime      string //
}

// agentGiveawayPropsColumns holds the columns for table agent_giveaway_props.
var agentGiveawayPropsColumns = AgentGiveawayPropsColumns{
	Id:              "id",
	UserId:          "user_id",
	PropsName:       "props_name",
	PropsIcon:       "props_icon",
	PropsDay:        "props_day",
	PropsStates:     "props_states",
	PropsType:       "props_type",
	PropsLink:       "props_link",
	BalanceRecordId: "balance_record_id",
	UseTime:         "use_time",
	TargetUserId:    "target_user_id",
	ExpirationTime:  "expiration_time",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewAgentGiveawayPropsDao creates and returns a new DAO object for table data access.
func NewAgentGiveawayPropsDao() *AgentGiveawayPropsDao {
	return &AgentGiveawayPropsDao{
		group:   "default",
		table:   "agent_giveaway_props",
		columns: agentGiveawayPropsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AgentGiveawayPropsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AgentGiveawayPropsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AgentGiveawayPropsDao) Columns() AgentGiveawayPropsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AgentGiveawayPropsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AgentGiveawayPropsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AgentGiveawayPropsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

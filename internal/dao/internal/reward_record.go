// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RewardRecordDao is the data access object for table reward_record.
type RewardRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns RewardRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RewardRecordColumns defines and stores column names for table reward_record.
type RewardRecordColumns struct {
	Id         string // 自增Id
	UserId     string // 用户ID
	GuildId    string // 公会id
	RewardType string // 奖励类型，  比如type  1：为3周2次结算  2……  后续扩展
	Reward     string // 活动奖励给公会 印尼盾 分   该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	Remark     string // 备注
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// rewardRecordColumns holds the columns for table reward_record.
var rewardRecordColumns = RewardRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	GuildId:    "guild_id",
	RewardType: "reward_type",
	Reward:     "reward",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewRewardRecordDao creates and returns a new DAO object for table data access.
func NewRewardRecordDao() *RewardRecordDao {
	return &RewardRecordDao{
		group:   "default",
		table:   "reward_record",
		columns: rewardRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RewardRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RewardRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RewardRecordDao) Columns() RewardRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RewardRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RewardRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RewardRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

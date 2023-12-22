// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StarLevelResetConfigDao is the data access object for table star_level_reset_config.
type StarLevelResetConfigDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns StarLevelResetConfigColumns // columns contains all the column names of Table for convenient usage.
}

// StarLevelResetConfigColumns defines and stores column names for table star_level_reset_config.
type StarLevelResetConfigColumns struct {
	Id                string // 自增id
	StarLevel         string // 主播星级
	RuleType          string // 类别 0保级 1升级 2降级
	IncomeMin         string // 金币min 不用扩大100
	SendGiftPersonMin string // 送礼人数min
	SelfSendRateMax   string // 自己送自己金币占比max 0.4 存储40 扩大100倍
	CreateTime        string // 创建时间
	UpdateTime        string // 更新时间
}

// starLevelResetConfigColumns holds the columns for table star_level_reset_config.
var starLevelResetConfigColumns = StarLevelResetConfigColumns{
	Id:                "id",
	StarLevel:         "star_level",
	RuleType:          "rule_type",
	IncomeMin:         "income_min",
	SendGiftPersonMin: "send_gift_person_min",
	SelfSendRateMax:   "self_send_rate_max",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// NewStarLevelResetConfigDao creates and returns a new DAO object for table data access.
func NewStarLevelResetConfigDao() *StarLevelResetConfigDao {
	return &StarLevelResetConfigDao{
		group:   "default",
		table:   "star_level_reset_config",
		columns: starLevelResetConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *StarLevelResetConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *StarLevelResetConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *StarLevelResetConfigDao) Columns() StarLevelResetConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *StarLevelResetConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *StarLevelResetConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *StarLevelResetConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyBattleDao is the data access object for table apply_battle.
type ApplyBattleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ApplyBattleColumns // columns contains all the column names of Table for convenient usage.
}

// ApplyBattleColumns defines and stores column names for table apply_battle.
type ApplyBattleColumns struct {
	Id               string // 自增id
	GuildId          string // 公会Id 关联guild_info自增
	UserId           string // 主播ID
	AdminOperateName string // 审核人id
	Status           string // 状态：0待审核 1通过 2不通过
	ApplyType        string // 报名类型，1公会，2 主播
	CreateTime       string // 开始时间
	UpdateTime       string // 开始时间
}

// applyBattleColumns holds the columns for table apply_battle.
var applyBattleColumns = ApplyBattleColumns{
	Id:               "id",
	GuildId:          "guild_id",
	UserId:           "user_id",
	AdminOperateName: "admin_operate_name",
	Status:           "status",
	ApplyType:        "apply_type",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewApplyBattleDao creates and returns a new DAO object for table data access.
func NewApplyBattleDao() *ApplyBattleDao {
	return &ApplyBattleDao{
		group:   "default",
		table:   "apply_battle",
		columns: applyBattleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplyBattleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplyBattleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplyBattleDao) Columns() ApplyBattleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplyBattleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplyBattleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplyBattleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

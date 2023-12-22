// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildApplyDao is the data access object for table guild_apply.
type GuildApplyDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns GuildApplyColumns // columns contains all the column names of Table for convenient usage.
}

// GuildApplyColumns defines and stores column names for table guild_apply.
type GuildApplyColumns struct {
	Id         string // 自增id
	UserId     string // 申请人用户Id
	GuildName  string // 公会名称
	AnchorNum  string // 主播数量
	Whatsapp   string // 联系WhatsApp
	State      string // 状态：0待审核 1通过 2不通过
	Operator   string // 审核人
	Remark     string // 备注
	CreateTime string //
	UpdateTime string //
}

// guildApplyColumns holds the columns for table guild_apply.
var guildApplyColumns = GuildApplyColumns{
	Id:         "id",
	UserId:     "user_id",
	GuildName:  "guild_name",
	AnchorNum:  "anchor_num",
	Whatsapp:   "whatsapp",
	State:      "state",
	Operator:   "operator",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewGuildApplyDao creates and returns a new DAO object for table data access.
func NewGuildApplyDao() *GuildApplyDao {
	return &GuildApplyDao{
		group:   "default",
		table:   "guild_apply",
		columns: guildApplyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildApplyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildApplyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildApplyDao) Columns() GuildApplyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildApplyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildApplyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildApplyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

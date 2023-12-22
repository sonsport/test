// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildAnchorRecordDao is the data access object for table guild_anchor_record.
type GuildAnchorRecordDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns GuildAnchorRecordColumns // columns contains all the column names of Table for convenient usage.
}

// GuildAnchorRecordColumns defines and stores column names for table guild_anchor_record.
type GuildAnchorRecordColumns struct {
	Id                   string // 主键id
	UserId               string // 用户id
	AdminId              string // 管理员id
	ExpirationTime       string // 过期时间
	BackgroundAdjustment string // 是否后台调整
	UpdateType           string // 调整类型，1星级，2公会，3标签
	IsHot                string // 是否在热门列表
	OldContent           string // 修改前内容
	NewContent           string // 修改后内容
	CreateTime           string // 创建时间
	UpdateTime           string // 更新时间
}

// guildAnchorRecordColumns holds the columns for table guild_anchor_record.
var guildAnchorRecordColumns = GuildAnchorRecordColumns{
	Id:                   "id",
	UserId:               "user_id",
	AdminId:              "admin_id",
	ExpirationTime:       "expiration_time",
	BackgroundAdjustment: "background_adjustment",
	UpdateType:           "update_type",
	IsHot:                "is_hot",
	OldContent:           "old_content",
	NewContent:           "new_content",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
}

// NewGuildAnchorRecordDao creates and returns a new DAO object for table data access.
func NewGuildAnchorRecordDao() *GuildAnchorRecordDao {
	return &GuildAnchorRecordDao{
		group:   "default",
		table:   "guild_anchor_record",
		columns: guildAnchorRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildAnchorRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildAnchorRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildAnchorRecordDao) Columns() GuildAnchorRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildAnchorRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildAnchorRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildAnchorRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

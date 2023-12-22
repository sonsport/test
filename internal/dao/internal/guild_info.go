// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildInfoDao is the data access object for table guild_info.
type GuildInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GuildInfoColumns // columns contains all the column names of Table for convenient usage.
}

// GuildInfoColumns defines and stores column names for table guild_info.
type GuildInfoColumns struct {
	Id           string // 公会Id
	UserId       string // 公会长Id
	GuildName    string // 公会名称
	Gcode        string // 公会邀请码
	MemberCount  string // 成员数量
	Portrait     string // 公会头像
	Status       string // 状态
	Notice       string // 公告
	CutRate      string // 公会抽成 设置范围0-100
	IsPayToGuild string // 工资是否发放给公会 0非 1是
	Remark       string // 备注
	CreateTime   string //
	UpdateTime   string //
	ServerId     string // 运营号id
	SourceType   string // 绑定来源0无，1运营号直接邀请，2工会邀请，3后台绑定
}

// guildInfoColumns holds the columns for table guild_info.
var guildInfoColumns = GuildInfoColumns{
	Id:           "id",
	UserId:       "user_id",
	GuildName:    "guild_name",
	Gcode:        "gcode",
	MemberCount:  "member_count",
	Portrait:     "portrait",
	Status:       "status",
	Notice:       "notice",
	CutRate:      "cut_rate",
	IsPayToGuild: "is_pay_to_guild",
	Remark:       "remark",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	ServerId:     "server_id",
	SourceType:   "source_type",
}

// NewGuildInfoDao creates and returns a new DAO object for table data access.
func NewGuildInfoDao() *GuildInfoDao {
	return &GuildInfoDao{
		group:   "default",
		table:   "guild_info",
		columns: guildInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildInfoDao) Columns() GuildInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

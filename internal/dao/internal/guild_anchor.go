// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildAnchorDao is the data access object for table guild_anchor.
type GuildAnchorDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns GuildAnchorColumns // columns contains all the column names of Table for convenient usage.
}

// GuildAnchorColumns defines and stores column names for table guild_anchor.
type GuildAnchorColumns struct {
	Id              string // 系统自增
	GuildId         string // 公会Id 关联guild_info自增
	UserId          string // 用户Id
	GuildRole       string // 1公会长 2公会成员
	Status          string // 状态 1 正常；2 失效
	Remark          string // 备注
	JoinTime        string // 加入时间
	StarLevel       string // 主播星级
	LabelType       string // 标签类型
	Label           string // 主播标签
	Score           string // 主播基础分数
	LivePermissions string // 1钻石房，2密码房，3钻石房和密码房
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
	CloudRecording  string // 录制状态 0不录制，1录制
}

// guildAnchorColumns holds the columns for table guild_anchor.
var guildAnchorColumns = GuildAnchorColumns{
	Id:              "id",
	GuildId:         "guild_id",
	UserId:          "user_id",
	GuildRole:       "guild_role",
	Status:          "status",
	Remark:          "remark",
	JoinTime:        "join_time",
	StarLevel:       "star_level",
	LabelType:       "label_type",
	Label:           "label",
	Score:           "score",
	LivePermissions: "live_permissions",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	CloudRecording:  "cloud_recording",
}

// NewGuildAnchorDao creates and returns a new DAO object for table data access.
func NewGuildAnchorDao() *GuildAnchorDao {
	return &GuildAnchorDao{
		group:   "default",
		table:   "guild_anchor",
		columns: guildAnchorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildAnchorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildAnchorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildAnchorDao) Columns() GuildAnchorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildAnchorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildAnchorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildAnchorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

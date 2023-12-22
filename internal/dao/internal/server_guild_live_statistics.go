// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerGuildLiveStatisticsDao is the data access object for table server_guild_live_statistics.
type ServerGuildLiveStatisticsDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns ServerGuildLiveStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// ServerGuildLiveStatisticsColumns defines and stores column names for table server_guild_live_statistics.
type ServerGuildLiveStatisticsColumns struct {
	Id                    string // 主键ID
	BeginTime             string // 开始时间
	EndTime               string // 结束时间
	GuildId               string // 工会id
	GuildName             string // 工会名称
	ServerId              string // 运营号id
	LiveAnchorCount       string // 主播开播数
	LiveStarAnchorCount   string // 标星主播开播数
	ActiveAnchorCount     string // 主播在线数
	StarActiveAnchorCount string // 标星主播在线数
}

// serverGuildLiveStatisticsColumns holds the columns for table server_guild_live_statistics.
var serverGuildLiveStatisticsColumns = ServerGuildLiveStatisticsColumns{
	Id:                    "id",
	BeginTime:             "begin_time",
	EndTime:               "end_time",
	GuildId:               "guild_id",
	GuildName:             "guild_name",
	ServerId:              "server_id",
	LiveAnchorCount:       "live_anchor_count",
	LiveStarAnchorCount:   "live_star_anchor_count",
	ActiveAnchorCount:     "active_anchor_count",
	StarActiveAnchorCount: "star_active_anchor_count",
}

// NewServerGuildLiveStatisticsDao creates and returns a new DAO object for table data access.
func NewServerGuildLiveStatisticsDao() *ServerGuildLiveStatisticsDao {
	return &ServerGuildLiveStatisticsDao{
		group:   "default",
		table:   "server_guild_live_statistics",
		columns: serverGuildLiveStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerGuildLiveStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerGuildLiveStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerGuildLiveStatisticsDao) Columns() ServerGuildLiveStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerGuildLiveStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerGuildLiveStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerGuildLiveStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

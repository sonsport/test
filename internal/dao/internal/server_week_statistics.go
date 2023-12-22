// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerWeekStatisticsDao is the data access object for table server_week_statistics.
type ServerWeekStatisticsDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns ServerWeekStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// ServerWeekStatisticsColumns defines and stores column names for table server_week_statistics.
type ServerWeekStatisticsColumns struct {
	Id                string // 自增id
	UserId            string // 运营号id
	Week              string // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildNum          string // 公会数量
	EffectiveGuildNum string // 有效公会数量
	NewGuildNum       string // 新增公会数量
	CreateTime        string // 创建时间
	UpdateTime        string // 更新时间
}

// serverWeekStatisticsColumns holds the columns for table server_week_statistics.
var serverWeekStatisticsColumns = ServerWeekStatisticsColumns{
	Id:                "id",
	UserId:            "user_id",
	Week:              "week",
	GuildNum:          "guild_num",
	EffectiveGuildNum: "effective_guild_num",
	NewGuildNum:       "new_guild_num",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
}

// NewServerWeekStatisticsDao creates and returns a new DAO object for table data access.
func NewServerWeekStatisticsDao() *ServerWeekStatisticsDao {
	return &ServerWeekStatisticsDao{
		group:   "default",
		table:   "server_week_statistics",
		columns: serverWeekStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerWeekStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerWeekStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerWeekStatisticsDao) Columns() ServerWeekStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerWeekStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerWeekStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerWeekStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrlsDao is the data access object for table short_urls.
type ShortUrlsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ShortUrlsColumns // columns contains all the column names of Table for convenient usage.
}

// ShortUrlsColumns defines and stores column names for table short_urls.
type ShortUrlsColumns struct {
	Id         string //
	UserId     string //
	RawUrl     string //
	Source     string //
	ShortCode  string //
	CreateTime string //
}

// shortUrlsColumns holds the columns for table short_urls.
var shortUrlsColumns = ShortUrlsColumns{
	Id:         "id",
	UserId:     "user_id",
	RawUrl:     "raw_url",
	Source:     "source",
	ShortCode:  "short_code",
	CreateTime: "create_time",
}

// NewShortUrlsDao creates and returns a new DAO object for table data access.
func NewShortUrlsDao() *ShortUrlsDao {
	return &ShortUrlsDao{
		group:   "default",
		table:   "short_urls",
		columns: shortUrlsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShortUrlsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShortUrlsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShortUrlsDao) Columns() ShortUrlsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShortUrlsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShortUrlsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShortUrlsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

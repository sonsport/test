// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthTokenDao is the data access object for table oauth_token.
type OauthTokenDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns OauthTokenColumns // columns contains all the column names of Table for convenient usage.
}

// OauthTokenColumns defines and stores column names for table oauth_token.
type OauthTokenColumns struct {
	UserId     string //
	OpenId     string //
	LoginType  string //
	CreateTime string //
	UpdateTime string //
}

// oauthTokenColumns holds the columns for table oauth_token.
var oauthTokenColumns = OauthTokenColumns{
	UserId:     "user_id",
	OpenId:     "open_id",
	LoginType:  "login_type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewOauthTokenDao creates and returns a new DAO object for table data access.
func NewOauthTokenDao() *OauthTokenDao {
	return &OauthTokenDao{
		group:   "default",
		table:   "oauth_token",
		columns: oauthTokenColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OauthTokenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OauthTokenDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OauthTokenDao) Columns() OauthTokenColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OauthTokenDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OauthTokenDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OauthTokenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

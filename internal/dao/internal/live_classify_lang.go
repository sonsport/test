// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveClassifyLangDao is the data access object for table live_classify_lang.
type LiveClassifyLangDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns LiveClassifyLangColumns // columns contains all the column names of Table for convenient usage.
}

// LiveClassifyLangColumns defines and stores column names for table live_classify_lang.
type LiveClassifyLangColumns struct {
	Id               string //
	ClassifyId       string // 分类id
	ClassifyLang     string // 分类语言
	ClassifyName     string // 分类名称
	ClassifyIcon     string // 分类图标
	ClassifyPickIcon string //
	CreateTime       string //
	UpdateTime       string //
}

// liveClassifyLangColumns holds the columns for table live_classify_lang.
var liveClassifyLangColumns = LiveClassifyLangColumns{
	Id:               "id",
	ClassifyId:       "classify_id",
	ClassifyLang:     "classify_lang",
	ClassifyName:     "classify_name",
	ClassifyIcon:     "classify_icon",
	ClassifyPickIcon: "classify_pick_icon",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewLiveClassifyLangDao creates and returns a new DAO object for table data access.
func NewLiveClassifyLangDao() *LiveClassifyLangDao {
	return &LiveClassifyLangDao{
		group:   "default",
		table:   "live_classify_lang",
		columns: liveClassifyLangColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveClassifyLangDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveClassifyLangDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveClassifyLangDao) Columns() LiveClassifyLangColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveClassifyLangDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveClassifyLangDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveClassifyLangDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

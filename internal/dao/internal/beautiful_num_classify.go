// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BeautifulNumClassifyDao is the data access object for table beautiful_num_classify.
type BeautifulNumClassifyDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns BeautifulNumClassifyColumns // columns contains all the column names of Table for convenient usage.
}

// BeautifulNumClassifyColumns defines and stores column names for table beautiful_num_classify.
type BeautifulNumClassifyColumns struct {
	ClassifyId   string // 分类id
	ClassifyName string //
	TranslateKey string // 翻译key
	IsShow       string // 是否显示
	CreateTime   string //
	UpdateTime   string //
}

// beautifulNumClassifyColumns holds the columns for table beautiful_num_classify.
var beautifulNumClassifyColumns = BeautifulNumClassifyColumns{
	ClassifyId:   "classify_id",
	ClassifyName: "classify_name",
	TranslateKey: "translate_key",
	IsShow:       "is_show",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewBeautifulNumClassifyDao creates and returns a new DAO object for table data access.
func NewBeautifulNumClassifyDao() *BeautifulNumClassifyDao {
	return &BeautifulNumClassifyDao{
		group:   "default",
		table:   "beautiful_num_classify",
		columns: beautifulNumClassifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BeautifulNumClassifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BeautifulNumClassifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BeautifulNumClassifyDao) Columns() BeautifulNumClassifyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BeautifulNumClassifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BeautifulNumClassifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BeautifulNumClassifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

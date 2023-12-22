// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AreaInfoDao is the data access object for table area_info.
type AreaInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AreaInfoColumns // columns contains all the column names of Table for convenient usage.
}

// AreaInfoColumns defines and stores column names for table area_info.
type AreaInfoColumns struct {
	Id           string //
	AreaCode     string // 地区
	Path         string //
	Title        string //
	TitleCn      string //
	Currency     string // 货币编码
	CurrencyCode string // 国家编码
	Sort         string //
	IsRich       string // 是否富有，0普通，1富有
	IsShow       string //
	CreatedTime  string //
	UpdatedTime  string //
}

// areaInfoColumns holds the columns for table area_info.
var areaInfoColumns = AreaInfoColumns{
	Id:           "id",
	AreaCode:     "area_code",
	Path:         "path",
	Title:        "title",
	TitleCn:      "title_cn",
	Currency:     "currency",
	CurrencyCode: "currency_code",
	Sort:         "sort",
	IsRich:       "is_rich",
	IsShow:       "is_show",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}

// NewAreaInfoDao creates and returns a new DAO object for table data access.
func NewAreaInfoDao() *AreaInfoDao {
	return &AreaInfoDao{
		group:   "default",
		table:   "area_info",
		columns: areaInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AreaInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AreaInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AreaInfoDao) Columns() AreaInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AreaInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AreaInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AreaInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BannerInfoDao is the data access object for table banner_info.
type BannerInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns BannerInfoColumns // columns contains all the column names of Table for convenient usage.
}

// BannerInfoColumns defines and stores column names for table banner_info.
type BannerInfoColumns struct {
	Bid        string // bid
	Ranking    string // 排序
	Title      string // 标题
	ClientOs   string // 0 全部；1 android；2 iOS
	AppName    string // app名称
	Cover      string // 封面URL
	Target     string // 跳转类型，0 网页，1 app内页
	Link       string // 跳转URL
	ShowType   string // 展示类型，0 全部，1 未充值用户 2 充值用户 3 主播
	CreateTime string //
	UpdateTime string //
	State      string // 状态1有效、2失效
}

// bannerInfoColumns holds the columns for table banner_info.
var bannerInfoColumns = BannerInfoColumns{
	Bid:        "bid",
	Ranking:    "ranking",
	Title:      "title",
	ClientOs:   "client_os",
	AppName:    "app_name",
	Cover:      "cover",
	Target:     "target",
	Link:       "link",
	ShowType:   "show_type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	State:      "state",
}

// NewBannerInfoDao creates and returns a new DAO object for table data access.
func NewBannerInfoDao() *BannerInfoDao {
	return &BannerInfoDao{
		group:   "default",
		table:   "banner_info",
		columns: bannerInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BannerInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BannerInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BannerInfoDao) Columns() BannerInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BannerInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BannerInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BannerInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

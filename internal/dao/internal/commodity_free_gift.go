// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommodityFreeGiftDao is the data access object for table commodity_free_gift.
type CommodityFreeGiftDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns CommodityFreeGiftColumns // columns contains all the column names of Table for convenient usage.
}

// CommodityFreeGiftColumns defines and stores column names for table commodity_free_gift.
type CommodityFreeGiftColumns struct {
	Id       string //
	Cid      string //
	State    string // 状态 1 启用
	GiftType string // 赠送类型	1 装扮  2 经验
	GiftLink string // 赠送关联	商品id
	GiftVal  string // 赠送数	装扮为天 经验为经验值
}

// commodityFreeGiftColumns holds the columns for table commodity_free_gift.
var commodityFreeGiftColumns = CommodityFreeGiftColumns{
	Id:       "id",
	Cid:      "cid",
	State:    "state",
	GiftType: "gift_type",
	GiftLink: "gift_link",
	GiftVal:  "gift_val",
}

// NewCommodityFreeGiftDao creates and returns a new DAO object for table data access.
func NewCommodityFreeGiftDao() *CommodityFreeGiftDao {
	return &CommodityFreeGiftDao{
		group:   "default",
		table:   "commodity_free_gift",
		columns: commodityFreeGiftColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommodityFreeGiftDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommodityFreeGiftDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommodityFreeGiftDao) Columns() CommodityFreeGiftColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommodityFreeGiftDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommodityFreeGiftDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommodityFreeGiftDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

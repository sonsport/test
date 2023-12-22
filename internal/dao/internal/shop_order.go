// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrderDao is the data access object for table shop_order.
type ShopOrderDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ShopOrderColumns // columns contains all the column names of Table for convenient usage.
}

// ShopOrderColumns defines and stores column names for table shop_order.
type ShopOrderColumns struct {
	Id         string //
	ShopId     string // 商品ID
	Num        string // 购买数量（月）
	BuyUid     string // 购买人uid
	ReceiveUid string // 接收人
	CoinNumber string // 购买金额 钻石数（不扩大）
	Desc       string // 订单描述
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// shopOrderColumns holds the columns for table shop_order.
var shopOrderColumns = ShopOrderColumns{
	Id:         "id",
	ShopId:     "shop_id",
	Num:        "num",
	BuyUid:     "buy_uid",
	ReceiveUid: "receive_uid",
	CoinNumber: "coin_number",
	Desc:       "desc",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewShopOrderDao creates and returns a new DAO object for table data access.
func NewShopOrderDao() *ShopOrderDao {
	return &ShopOrderDao{
		group:   "default",
		table:   "shop_order",
		columns: shopOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShopOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShopOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShopOrderDao) Columns() ShopOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShopOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShopOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShopOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

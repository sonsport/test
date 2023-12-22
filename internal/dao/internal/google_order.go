// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoogleOrderDao is the data access object for table google_order.
type GoogleOrderDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns GoogleOrderColumns // columns contains all the column names of Table for convenient usage.
}

// GoogleOrderColumns defines and stores column names for table google_order.
type GoogleOrderColumns struct {
	Id            string //
	UserId        string // 用户Id
	Receipt       string // inAppPurchaseData 回调
	Signature     string // google签名
	InternalIp    string // 总购买钻石 只累计购买的
	ProductId     string // 商品Id
	GpOrderId     string // google订单idGPA.0000
	PurchaseToken string //
	AppName       string // app名称
	Package       string // 包名 com.aaa.bb
	OrderId       string // 系统订单ID
	State         string // 状态0待处理 1处理完成
	CreateTime    string //
	UpdateTime    string //
}

// googleOrderColumns holds the columns for table google_order.
var googleOrderColumns = GoogleOrderColumns{
	Id:            "id",
	UserId:        "user_id",
	Receipt:       "receipt",
	Signature:     "signature",
	InternalIp:    "internal_ip",
	ProductId:     "product_id",
	GpOrderId:     "gp_order_id",
	PurchaseToken: "purchase_token",
	AppName:       "app_name",
	Package:       "package",
	OrderId:       "order_id",
	State:         "state",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewGoogleOrderDao creates and returns a new DAO object for table data access.
func NewGoogleOrderDao() *GoogleOrderDao {
	return &GoogleOrderDao{
		group:   "default",
		table:   "google_order",
		columns: googleOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoogleOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoogleOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoogleOrderDao) Columns() GoogleOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoogleOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoogleOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoogleOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

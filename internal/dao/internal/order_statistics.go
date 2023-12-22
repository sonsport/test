// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderStatisticsDao is the data access object for table order_statistics.
type OrderStatisticsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns OrderStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// OrderStatisticsColumns defines and stores column names for table order_statistics.
type OrderStatisticsColumns struct {
	Id                string //
	BeginTime         string // 开始时间
	EndTime           string // 结束时间
	TimeLevel         string // 时间维度：1.每小时，2.每天
	AppChannel        string // app渠道
	PayType           string // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	ChannelCode       string // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	TotalFee          string // 商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段
	AllOrderCount     string // 总订单数
	SuccessOrderCount string // 付费订单数
}

// orderStatisticsColumns holds the columns for table order_statistics.
var orderStatisticsColumns = OrderStatisticsColumns{
	Id:                "id",
	BeginTime:         "begin_time",
	EndTime:           "end_time",
	TimeLevel:         "time_level",
	AppChannel:        "app_channel",
	PayType:           "pay_type",
	ChannelCode:       "channel_code",
	TotalFee:          "total_fee",
	AllOrderCount:     "all_order_count",
	SuccessOrderCount: "success_order_count",
}

// NewOrderStatisticsDao creates and returns a new DAO object for table data access.
func NewOrderStatisticsDao() *OrderStatisticsDao {
	return &OrderStatisticsDao{
		group:   "default",
		table:   "order_statistics",
		columns: orderStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderStatisticsDao) Columns() OrderStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

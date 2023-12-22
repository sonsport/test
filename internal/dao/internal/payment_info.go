// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PaymentInfoDao is the data access object for table payment_info.
type PaymentInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns PaymentInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PaymentInfoColumns defines and stores column names for table payment_info.
type PaymentInfoColumns struct {
	Id           string //
	AppChannel   string // app渠道
	PayType      string // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	PaymentName  string // 支付渠道：比如Google、OVO、DANA等等
	CurrencyCode string // 国家数字码，比如印尼：360；马来：458  0为默认
	PayIcon      string // 支付图标
	ChannelCode  string // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	Outside      string // 是否外部网页打开：0不需要，1需要
	State        string // 状态：0下架、1上架
	SystemOs     string // 所属系统；0所有,1安卓 2ios
	IsDefault    string // 是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付
	Sort         string // 排序，顺序排列
	Remark       string // 说明备注
	CreateTime   string //
	UpdateTime   string //
}

// paymentInfoColumns holds the columns for table payment_info.
var paymentInfoColumns = PaymentInfoColumns{
	Id:           "id",
	AppChannel:   "app_channel",
	PayType:      "pay_type",
	PaymentName:  "payment_name",
	CurrencyCode: "currency_code",
	PayIcon:      "pay_icon",
	ChannelCode:  "channel_code",
	Outside:      "outside",
	State:        "state",
	SystemOs:     "system_os",
	IsDefault:    "is_default",
	Sort:         "sort",
	Remark:       "remark",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewPaymentInfoDao creates and returns a new DAO object for table data access.
func NewPaymentInfoDao() *PaymentInfoDao {
	return &PaymentInfoDao{
		group:   "default",
		table:   "payment_info",
		columns: paymentInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PaymentInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PaymentInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PaymentInfoDao) Columns() PaymentInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PaymentInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PaymentInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PaymentInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrderInfoDao is the data access object for table order_info.
type OrderInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns OrderInfoColumns // columns contains all the column names of Table for convenient usage.
}

// OrderInfoColumns defines and stores column names for table order_info.
type OrderInfoColumns struct {
	Oid                string //
	Tid                string // 票据Id
	UserId             string // 用户Id
	Product            string // 商品Id 对应commodity_info主键
	ProductValue       string // 商品钻石
	ProductBonus       string // 商品赠送钻石
	ProductDescription string // 商品描述
	TotalFee           string // 商品总费用 印尼盾 单位分 对应commodity_info的totalFee_cent字段
	Currency           string // 价格对应的单位 比如：IDR、MYR 前台支付  对应commodity_info的currency字段
	CurrencyPrice      string // 价格 分 对应商品手机支付当地货币价格
	IsFirstTopup       string // 是否首次充值 0非首充 1首充  对应商品表的is_first_topup字段
	State              string // 状态 0为待支付 1为已支付
	SystemOs           string // 所属系统；0未知,1安卓 2ios
	AppChannel         string // 所属渠道
	AppName            string // 所属app名称
	PayType            string // 支付大渠道；比如：1为google支付、2为Apple支付、3为xenditpay、4为codapay
	ChannelCode        string // 支付渠道标识：ovo为5、XL为23等，跟三方支付渠道相关，系统无法做决定
	Remarks            string // 备注说明
	CouponId           string // 优惠券id
	ReferrerId         string // 推荐人id
	CertificateUrl     string // 代理充值凭证
	CreateTime         string //
	UpdateTime         string //
}

// orderInfoColumns holds the columns for table order_info.
var orderInfoColumns = OrderInfoColumns{
	Oid:                "oid",
	Tid:                "tid",
	UserId:             "user_id",
	Product:            "product",
	ProductValue:       "product_value",
	ProductBonus:       "product_bonus",
	ProductDescription: "product_description",
	TotalFee:           "total_fee",
	Currency:           "currency",
	CurrencyPrice:      "currency_price",
	IsFirstTopup:       "is_first_topup",
	State:              "state",
	SystemOs:           "system_os",
	AppChannel:         "app_channel",
	AppName:            "app_name",
	PayType:            "pay_type",
	ChannelCode:        "channel_code",
	Remarks:            "remarks",
	CouponId:           "coupon_id",
	ReferrerId:         "referrer_id",
	CertificateUrl:     "certificate_url",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewOrderInfoDao creates and returns a new DAO object for table data access.
func NewOrderInfoDao() *OrderInfoDao {
	return &OrderInfoDao{
		group:   "default",
		table:   "order_info",
		columns: orderInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrderInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrderInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrderInfoDao) Columns() OrderInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrderInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrderInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrderInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

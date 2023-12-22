// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CouponInfoDao is the data access object for table coupon_info.
type CouponInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns CouponInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CouponInfoColumns defines and stores column names for table coupon_info.
type CouponInfoColumns struct {
	Id          string //
	UserId      string // 用户Id
	BeginTime   string // 开始时间
	EndTime     string // 结束时间
	CouponState string // 状态 1未使用 2已使用
	UseTime     string // 使用时间
	CouponType  string // 优惠券类型 1 钻石赠送券
	CouponRatio string // 赠送比例 30 则赠送订单总额的30%
	MinShift    string // 最小支持充值档位
	MaxShift    string // 最大支持充值档位
	LinkId      string // 优惠券来源
	CreateTime  string //
	UpdateTime  string //
}

// couponInfoColumns holds the columns for table coupon_info.
var couponInfoColumns = CouponInfoColumns{
	Id:          "id",
	UserId:      "user_id",
	BeginTime:   "begin_time",
	EndTime:     "end_time",
	CouponState: "coupon_state",
	UseTime:     "use_time",
	CouponType:  "coupon_type",
	CouponRatio: "coupon_ratio",
	MinShift:    "min_shift",
	MaxShift:    "max_shift",
	LinkId:      "link_id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewCouponInfoDao creates and returns a new DAO object for table data access.
func NewCouponInfoDao() *CouponInfoDao {
	return &CouponInfoDao{
		group:   "default",
		table:   "coupon_info",
		columns: couponInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CouponInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CouponInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CouponInfoDao) Columns() CouponInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CouponInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CouponInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CouponInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

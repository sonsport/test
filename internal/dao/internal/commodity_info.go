// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommodityInfoDao is the data access object for table commodity_info.
type CommodityInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns CommodityInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CommodityInfoColumns defines and stores column names for table commodity_info.
type CommodityInfoColumns struct {
	Cid          string // 商品Id 唯一约束
	PaymentId    string // 关联payment_info表主键
	Name         string // 名称 比如80钻
	Price        string // 价格 分
	Value        string // 钻石数
	Bonus        string // 赠送钻石数
	Currency     string // 价格对应的单位 比如：IDR、MYR 前台支付
	TotalFeeCent string // 对应印尼盾价格 单位分 统计使用
	IsFirstTopup string // 是否首次充值优惠档位，0默认，1首充优惠档位、2非首充档位
	Sort         string // 排序、顺序
	State        string // 状态；0下架；1上架
	IsDefault    string // 是否默认：0非默认、1默认 只能设置一个为默认 默认的渠道显示在快捷支付
	OriginBonus  string // 折扣前赠送
	DiscountType string // 折扣类型，1.首充折扣
	CreateTime   string //
	UpdateTime   string //
}

// commodityInfoColumns holds the columns for table commodity_info.
var commodityInfoColumns = CommodityInfoColumns{
	Cid:          "cid",
	PaymentId:    "payment_id",
	Name:         "name",
	Price:        "price",
	Value:        "value",
	Bonus:        "bonus",
	Currency:     "currency",
	TotalFeeCent: "totalFee_cent",
	IsFirstTopup: "is_first_topup",
	Sort:         "sort",
	State:        "state",
	IsDefault:    "is_default",
	OriginBonus:  "origin_bonus",
	DiscountType: "discount_type",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewCommodityInfoDao creates and returns a new DAO object for table data access.
func NewCommodityInfoDao() *CommodityInfoDao {
	return &CommodityInfoDao{
		group:   "default",
		table:   "commodity_info",
		columns: commodityInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommodityInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommodityInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommodityInfoDao) Columns() CommodityInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommodityInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommodityInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommodityInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

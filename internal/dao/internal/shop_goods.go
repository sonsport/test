// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopGoodsDao is the data access object for table shop_goods.
type ShopGoodsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ShopGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// ShopGoodsColumns defines and stores column names for table shop_goods.
type ShopGoodsColumns struct {
	Id             string // 自增ID
	CategoryId     string // 类型
	Name           string // 商品名称
	Status         string // 商品状态 0下架 1上架
	IconUrl        string // 图标地址
	ResUrl         string // 资源地址
	IconSize       string // 图标尺寸
	Price          string // 单价 钻石数
	ForeverPrice   string // 永久售卖价值 后续使用对于有永久售卖的商品
	Desc           string // 商品描述
	DiscountPrice  string // 折扣 json 配置不同天数的价格
	UserLevelLimit string // 用户等级限制
	ShelfTime      string // 上架时间
	IsRenewable    string // 是否可续费 该字段基本是没用
	SortWeight     string // 排序权重 升序
	SellType       string // 售卖类型 1-售卖商品；2-活动商品
	Quality        string // 品质
	Extra          string // 扩展属性
	CreateTime     string //
	UpdateTime     string //
	Number         string // 号码
}

// shopGoodsColumns holds the columns for table shop_goods.
var shopGoodsColumns = ShopGoodsColumns{
	Id:             "id",
	CategoryId:     "category_id",
	Name:           "name",
	Status:         "status",
	IconUrl:        "icon_url",
	ResUrl:         "res_url",
	IconSize:       "icon_size",
	Price:          "price",
	ForeverPrice:   "forever_price",
	Desc:           "desc",
	DiscountPrice:  "discount_price",
	UserLevelLimit: "user_level_limit",
	ShelfTime:      "shelf_time",
	IsRenewable:    "is_renewable",
	SortWeight:     "sort_weight",
	SellType:       "sell_type",
	Quality:        "quality",
	Extra:          "extra",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
	Number:         "number",
}

// NewShopGoodsDao creates and returns a new DAO object for table data access.
func NewShopGoodsDao() *ShopGoodsDao {
	return &ShopGoodsDao{
		group:   "default",
		table:   "shop_goods",
		columns: shopGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShopGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShopGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShopGoodsDao) Columns() ShopGoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShopGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShopGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShopGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BeautifulNumDao is the data access object for table beautiful_num.
type BeautifulNumDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns BeautifulNumColumns // columns contains all the column names of Table for convenient usage.
}

// BeautifulNumColumns defines and stores column names for table beautiful_num.
type BeautifulNumColumns struct {
	Id             string //
	SuperId        string // 靓号
	ClassifyId     string // 分类id
	IconUrl        string // icon地址
	ShopStatus     string // 状态 1可用 2不可用
	SortWeight     string // 排序权重
	SkuListPrice   string // 规格列表价格
	UserLevelLimit string // 用户等级限制
	Price          string // 单价
	UserId         string // 用户id
	IsWear         string // 是否佩戴
	Experience     string // 经验值
	ExpirationTime string // 过期时间
	ReclaimTime    string // 回收时间
	CreateTime     string //
	UpdateTime     string //
}

// beautifulNumColumns holds the columns for table beautiful_num.
var beautifulNumColumns = BeautifulNumColumns{
	Id:             "id",
	SuperId:        "super_id",
	ClassifyId:     "classify_id",
	IconUrl:        "icon_url",
	ShopStatus:     "shop_status",
	SortWeight:     "sort_weight",
	SkuListPrice:   "sku_list_price",
	UserLevelLimit: "user_level_limit",
	Price:          "price",
	UserId:         "user_id",
	IsWear:         "is_wear",
	Experience:     "experience",
	ExpirationTime: "expiration_time",
	ReclaimTime:    "reclaim_time",
	CreateTime:     "create_time",
	UpdateTime:     "update_time",
}

// NewBeautifulNumDao creates and returns a new DAO object for table data access.
func NewBeautifulNumDao() *BeautifulNumDao {
	return &BeautifulNumDao{
		group:   "default",
		table:   "beautiful_num",
		columns: beautifulNumColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BeautifulNumDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BeautifulNumDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BeautifulNumDao) Columns() BeautifulNumColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BeautifulNumDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BeautifulNumDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BeautifulNumDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

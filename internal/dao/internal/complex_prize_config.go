// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ComplexPrizeConfigDao is the data access object for table complex_prize_config.
type ComplexPrizeConfigDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ComplexPrizeConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ComplexPrizeConfigColumns defines and stores column names for table complex_prize_config.
type ComplexPrizeConfigColumns struct {
	Id           string //
	ActivityType string // 抽奖活动类型 1 圣诞 2 元旦
	Icon         string // 图标
	DrawType     string // 抽奖类型 1装扮 2钻石 3优惠券 4 实物
	LinkId       string // 关联商品id  1,2  逗号分割
	PrizeName    string // 奖品名称
	Num          string // 数量
	IsFirstPool  string // 是否第一个奖池出现
	Rate         string // 概率万分比  *10000的整数
	Status       string // 配置状态 0失效 1生效
}

// complexPrizeConfigColumns holds the columns for table complex_prize_config.
var complexPrizeConfigColumns = ComplexPrizeConfigColumns{
	Id:           "id",
	ActivityType: "activity_type",
	Icon:         "icon",
	DrawType:     "draw_type",
	LinkId:       "link_id",
	PrizeName:    "prize_name",
	Num:          "num",
	IsFirstPool:  "is_first_pool",
	Rate:         "rate",
	Status:       "status",
}

// NewComplexPrizeConfigDao creates and returns a new DAO object for table data access.
func NewComplexPrizeConfigDao() *ComplexPrizeConfigDao {
	return &ComplexPrizeConfigDao{
		group:   "default",
		table:   "complex_prize_config",
		columns: complexPrizeConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ComplexPrizeConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ComplexPrizeConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ComplexPrizeConfigDao) Columns() ComplexPrizeConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ComplexPrizeConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ComplexPrizeConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ComplexPrizeConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

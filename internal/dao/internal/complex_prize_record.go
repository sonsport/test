// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ComplexPrizeRecordDao is the data access object for table complex_prize_record.
type ComplexPrizeRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ComplexPrizeRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ComplexPrizeRecordColumns defines and stores column names for table complex_prize_record.
type ComplexPrizeRecordColumns struct {
	Id           string //
	UserId       string // 用户id
	CId          string // 抽奖配置id
	ActivityType string // 抽奖活动类型 1 圣诞 2 元旦
	Icon         string // 图标
	DrawType     string // 抽奖类型 1装扮 2钻石 3优惠券 4 实物
	PropsLogId   string // 扣除道具流水id
	LinkId       string // 关联商品id  1,2  逗号分割
	PrizeName    string // 奖品名称
	PrizeCount   string // 奖池期数
	CreateTime   string // 创建时间
}

// complexPrizeRecordColumns holds the columns for table complex_prize_record.
var complexPrizeRecordColumns = ComplexPrizeRecordColumns{
	Id:           "id",
	UserId:       "user_id",
	CId:          "c_id",
	ActivityType: "activity_type",
	Icon:         "icon",
	DrawType:     "draw_type",
	PropsLogId:   "props_log_id",
	LinkId:       "link_id",
	PrizeName:    "prize_name",
	PrizeCount:   "prize_count",
	CreateTime:   "create_time",
}

// NewComplexPrizeRecordDao creates and returns a new DAO object for table data access.
func NewComplexPrizeRecordDao() *ComplexPrizeRecordDao {
	return &ComplexPrizeRecordDao{
		group:   "default",
		table:   "complex_prize_record",
		columns: complexPrizeRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ComplexPrizeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ComplexPrizeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ComplexPrizeRecordDao) Columns() ComplexPrizeRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ComplexPrizeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ComplexPrizeRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ComplexPrizeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

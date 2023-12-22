// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// QuestPrizeDao is the data access object for table quest_prize.
type QuestPrizeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns QuestPrizeColumns // columns contains all the column names of Table for convenient usage.
}

// QuestPrizeColumns defines and stores column names for table quest_prize.
type QuestPrizeColumns struct {
	Id         string // id
	QuestId    string // 任务id
	PrizeName  string // 奖品名称
	PrizeIcon  string // 奖品图标
	PrizeType  string // 奖品类型 1装扮 2钻石 3优惠券 4贝壳 5 活跃值 6 用户经验 7 主播经验 8 主播收益
	LinkId     string // 关联商品id  1,2  逗号分割
	PrizeNum   string // 奖品数量
	Status     string // 配置状态 0失效 1生效
	CreateTime string // createdAt
	UpdateTime string // updatedAt
}

// questPrizeColumns holds the columns for table quest_prize.
var questPrizeColumns = QuestPrizeColumns{
	Id:         "id",
	QuestId:    "quest_id",
	PrizeName:  "prize_name",
	PrizeIcon:  "prize_icon",
	PrizeType:  "prize_type",
	LinkId:     "link_id",
	PrizeNum:   "prize_num",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewQuestPrizeDao creates and returns a new DAO object for table data access.
func NewQuestPrizeDao() *QuestPrizeDao {
	return &QuestPrizeDao{
		group:   "default",
		table:   "quest_prize",
		columns: questPrizeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *QuestPrizeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *QuestPrizeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *QuestPrizeDao) Columns() QuestPrizeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *QuestPrizeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *QuestPrizeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *QuestPrizeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

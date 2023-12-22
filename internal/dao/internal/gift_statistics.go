// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GiftStatisticsDao is the data access object for table gift_statistics.
type GiftStatisticsDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns GiftStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// GiftStatisticsColumns defines and stores column names for table gift_statistics.
type GiftStatisticsColumns struct {
	Id            string //
	GiftId        string // 礼物id
	GiftNum       string // 礼物个数
	TotalDiamonds string // 礼物总价
	BeginTime     string // 开始时间
	EndTime       string // 结束时间
	TimeLevel     string // 时间维度：1.每小时，2.每天
	LuckyDiamonds string // 中奖钻石
	CreateTime    string //
	UpdateTime    string //
	UserCount     string // 送礼人数
	RoomCount     string // 房间数
}

// giftStatisticsColumns holds the columns for table gift_statistics.
var giftStatisticsColumns = GiftStatisticsColumns{
	Id:            "id",
	GiftId:        "gift_id",
	GiftNum:       "gift_num",
	TotalDiamonds: "total_diamonds",
	BeginTime:     "begin_time",
	EndTime:       "end_time",
	TimeLevel:     "time_level",
	LuckyDiamonds: "lucky_diamonds",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	UserCount:     "user_count",
	RoomCount:     "room_count",
}

// NewGiftStatisticsDao creates and returns a new DAO object for table data access.
func NewGiftStatisticsDao() *GiftStatisticsDao {
	return &GiftStatisticsDao{
		group:   "default",
		table:   "gift_statistics",
		columns: giftStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GiftStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GiftStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GiftStatisticsDao) Columns() GiftStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GiftStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GiftStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GiftStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

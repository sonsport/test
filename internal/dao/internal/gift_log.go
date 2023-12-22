// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GiftLogDao is the data access object for table gift_log.
type GiftLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns GiftLogColumns // columns contains all the column names of Table for convenient usage.
}

// GiftLogColumns defines and stores column names for table gift_log.
type GiftLogColumns struct {
	Id            string //
	SenderId      string // 送礼者uid
	ReceiverId    string // 收礼者uid
	GiftId        string // 礼物id
	GiftType      string // 礼物类型
	GiftNum       string // 礼物个数
	RoomId        string // 房间id
	Diamonds      string // 礼物单价
	TotalDiamonds string // 礼物总价
	CreateTime    string //
	UpdateTime    string //
}

// giftLogColumns holds the columns for table gift_log.
var giftLogColumns = GiftLogColumns{
	Id:            "id",
	SenderId:      "sender_id",
	ReceiverId:    "receiver_id",
	GiftId:        "gift_id",
	GiftType:      "gift_type",
	GiftNum:       "gift_num",
	RoomId:        "room_id",
	Diamonds:      "diamonds",
	TotalDiamonds: "total_diamonds",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewGiftLogDao creates and returns a new DAO object for table data access.
func NewGiftLogDao() *GiftLogDao {
	return &GiftLogDao{
		group:   "default",
		table:   "gift_log",
		columns: giftLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GiftLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GiftLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GiftLogDao) Columns() GiftLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GiftLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GiftLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GiftLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

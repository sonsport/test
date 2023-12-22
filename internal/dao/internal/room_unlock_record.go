// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomUnlockRecordDao is the data access object for table room_unlock_record.
type RoomUnlockRecordDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns RoomUnlockRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RoomUnlockRecordColumns defines and stores column names for table room_unlock_record.
type RoomUnlockRecordColumns struct {
	Id           string //
	RoomId       string //
	UserId       string // 用户id
	AnchorUserId string // 主播id
	IsUnlock     string // 是否解锁
	CreateTime   string //
	UpdateTime   string //
	UnlockPrice  string //
}

// roomUnlockRecordColumns holds the columns for table room_unlock_record.
var roomUnlockRecordColumns = RoomUnlockRecordColumns{
	Id:           "id",
	RoomId:       "room_id",
	UserId:       "user_id",
	AnchorUserId: "anchor_user_id",
	IsUnlock:     "is_unlock",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	UnlockPrice:  "unlock_price",
}

// NewRoomUnlockRecordDao creates and returns a new DAO object for table data access.
func NewRoomUnlockRecordDao() *RoomUnlockRecordDao {
	return &RoomUnlockRecordDao{
		group:   "default",
		table:   "room_unlock_record",
		columns: roomUnlockRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomUnlockRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomUnlockRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomUnlockRecordDao) Columns() RoomUnlockRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomUnlockRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomUnlockRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomUnlockRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

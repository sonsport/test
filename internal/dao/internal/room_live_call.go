// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLiveCallDao is the data access object for table room_live_call.
type RoomLiveCallDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns RoomLiveCallColumns // columns contains all the column names of Table for convenient usage.
}

// RoomLiveCallColumns defines and stores column names for table room_live_call.
type RoomLiveCallColumns struct {
	Id           string //
	UserId       string // 用户id
	AnchorUserId string // 主播用户id
	RoomId       string // 房间id
	CallStates   string // 连麦状态 0 排队中 1连麦中 2 已结束
	StreamId     string // 连麦流id
	CreateTime   string //
	UpdateTime   string //
}

// roomLiveCallColumns holds the columns for table room_live_call.
var roomLiveCallColumns = RoomLiveCallColumns{
	Id:           "id",
	UserId:       "user_id",
	AnchorUserId: "anchor_user_id",
	RoomId:       "room_id",
	CallStates:   "call_states",
	StreamId:     "stream_id",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewRoomLiveCallDao creates and returns a new DAO object for table data access.
func NewRoomLiveCallDao() *RoomLiveCallDao {
	return &RoomLiveCallDao{
		group:   "default",
		table:   "room_live_call",
		columns: roomLiveCallColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomLiveCallDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomLiveCallDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomLiveCallDao) Columns() RoomLiveCallColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomLiveCallDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomLiveCallDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomLiveCallDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

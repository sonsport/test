// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomRemoveListDao is the data access object for table room_remove_list.
type RoomRemoveListDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RoomRemoveListColumns // columns contains all the column names of Table for convenient usage.
}

// RoomRemoveListColumns defines and stores column names for table room_remove_list.
type RoomRemoveListColumns struct {
	Id         string // ID
	UserId     string // UID
	OperatorId string // 操作者ID
	RoomId     string // 房间ID
	EndTime    string // 移出结束时间,为0就是永久拉黑
	CreateTime string //
	UpdateTime string //
}

// roomRemoveListColumns holds the columns for table room_remove_list.
var roomRemoveListColumns = RoomRemoveListColumns{
	Id:         "id",
	UserId:     "user_id",
	OperatorId: "operator_id",
	RoomId:     "room_id",
	EndTime:    "end_time",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewRoomRemoveListDao creates and returns a new DAO object for table data access.
func NewRoomRemoveListDao() *RoomRemoveListDao {
	return &RoomRemoveListDao{
		group:   "default",
		table:   "room_remove_list",
		columns: roomRemoveListColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomRemoveListDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomRemoveListDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomRemoveListDao) Columns() RoomRemoveListColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomRemoveListDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomRemoveListDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomRemoveListDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

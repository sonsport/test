// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomAdminDao is the data access object for table room_admin.
type RoomAdminDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns RoomAdminColumns // columns contains all the column names of Table for convenient usage.
}

// RoomAdminColumns defines and stores column names for table room_admin.
type RoomAdminColumns struct {
	Id         string // ID
	UserId     string // 用户ID 存储主播的id
	ManagerId  string // 管理员ID 主播指定的管理员
	Valid      string // 是否有效，0无效 1有效
	CreateTime string //
	UpdateTime string //
}

// roomAdminColumns holds the columns for table room_admin.
var roomAdminColumns = RoomAdminColumns{
	Id:         "id",
	UserId:     "user_id",
	ManagerId:  "manager_id",
	Valid:      "valid",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewRoomAdminDao creates and returns a new DAO object for table data access.
func NewRoomAdminDao() *RoomAdminDao {
	return &RoomAdminDao{
		group:   "default",
		table:   "room_admin",
		columns: roomAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomAdminDao) Columns() RoomAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

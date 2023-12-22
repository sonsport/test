// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomUserRecordDao is the data access object for table room_user_record.
type RoomUserRecordDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RoomUserRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RoomUserRecordColumns defines and stores column names for table room_user_record.
type RoomUserRecordColumns struct {
	Id         string // 自增长ID
	UserId     string // 用户ID
	RoomId     string // 房间ID
	ExitType   string // 用户退出房间原因，0 自己退出；1 房间管理员；2 主播提出房间 3 超级管理踢出房间；4 主播拉黑；5 异常退出房间
	Extend     string // 额外信息
	IsChat     string // 是否有聊天
	IsWaitLong string // 是否待满5分钟
	Diamonds   string // 在此房间消费的钻石流水
	BeginTime  string // 进入房间时间
	EndTime    string // 退出房间时间
	CreateTime string //
	UpdateTime string //
}

// roomUserRecordColumns holds the columns for table room_user_record.
var roomUserRecordColumns = RoomUserRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	RoomId:     "room_id",
	ExitType:   "exit_type",
	Extend:     "extend",
	IsChat:     "is_chat",
	IsWaitLong: "is_wait_long",
	Diamonds:   "diamonds",
	BeginTime:  "begin_time",
	EndTime:    "end_time",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewRoomUserRecordDao creates and returns a new DAO object for table data access.
func NewRoomUserRecordDao() *RoomUserRecordDao {
	return &RoomUserRecordDao{
		group:   "default",
		table:   "room_user_record",
		columns: roomUserRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomUserRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomUserRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomUserRecordDao) Columns() RoomUserRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomUserRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomUserRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomUserRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

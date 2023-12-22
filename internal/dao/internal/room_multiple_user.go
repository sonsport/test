// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomMultipleUserDao is the data access object for table room_multiple_user.
type RoomMultipleUserDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns RoomMultipleUserColumns // columns contains all the column names of Table for convenient usage.
}

// RoomMultipleUserColumns defines and stores column names for table room_multiple_user.
type RoomMultipleUserColumns struct {
	Id               string //
	UserId           string // 用户id
	Nickname         string // 昵称
	Portrait         string // 头像
	RoomId           string // 房间id
	AnchorUserId     string // 主播用户id
	LiveControl      string // 直播控制 1 正常 2 禁视频
	SoundControl     string // 直播控制 1 正常 2 静音
	LiveStates       string // 多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略
	StreamId         string // 流id 房主的为流混流拉流id
	MultipleUserRole string // 多人房身份 1 多人房主播 2 上麦者
	MainStatus       string // 主麦状态
	CreateTime       string //
	UpdateTime       string //
}

// roomMultipleUserColumns holds the columns for table room_multiple_user.
var roomMultipleUserColumns = RoomMultipleUserColumns{
	Id:               "id",
	UserId:           "user_id",
	Nickname:         "nickname",
	Portrait:         "portrait",
	RoomId:           "room_id",
	AnchorUserId:     "anchor_user_id",
	LiveControl:      "live_control",
	SoundControl:     "sound_control",
	LiveStates:       "live_states",
	StreamId:         "stream_id",
	MultipleUserRole: "multiple_user_role",
	MainStatus:       "main_status",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewRoomMultipleUserDao creates and returns a new DAO object for table data access.
func NewRoomMultipleUserDao() *RoomMultipleUserDao {
	return &RoomMultipleUserDao{
		group:   "default",
		table:   "room_multiple_user",
		columns: roomMultipleUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomMultipleUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomMultipleUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomMultipleUserDao) Columns() RoomMultipleUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomMultipleUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomMultipleUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomMultipleUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

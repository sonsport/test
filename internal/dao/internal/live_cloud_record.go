// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveCloudRecordDao is the data access object for table live_cloud_record.
type LiveCloudRecordDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns LiveCloudRecordColumns // columns contains all the column names of Table for convenient usage.
}

// LiveCloudRecordColumns defines and stores column names for table live_cloud_record.
type LiveCloudRecordColumns struct {
	Id         string // 系统自增
	TaskId     string // 录制任务 ID，长度固定为 16 个字节的字符串。
	GuildId    string // 公会id
	RoomId     string // 录制房间 ID
	UserId     string // 录制流对应的推流用户 ID（混流时，为 mix_output_stream_id）。
	Status     string // 录制状态，1开始录制，2录制结束
	CreateTime string //
	UpdateTime string //
}

// liveCloudRecordColumns holds the columns for table live_cloud_record.
var liveCloudRecordColumns = LiveCloudRecordColumns{
	Id:         "id",
	TaskId:     "task_id",
	GuildId:    "guild_id",
	RoomId:     "room_id",
	UserId:     "user_id",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewLiveCloudRecordDao creates and returns a new DAO object for table data access.
func NewLiveCloudRecordDao() *LiveCloudRecordDao {
	return &LiveCloudRecordDao{
		group:   "default",
		table:   "live_cloud_record",
		columns: liveCloudRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveCloudRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveCloudRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveCloudRecordDao) Columns() LiveCloudRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveCloudRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveCloudRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveCloudRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

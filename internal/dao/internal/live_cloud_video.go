// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiveCloudVideoDao is the data access object for table live_cloud_video.
type LiveCloudVideoDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns LiveCloudVideoColumns // columns contains all the column names of Table for convenient usage.
}

// LiveCloudVideoColumns defines and stores column names for table live_cloud_video.
type LiveCloudVideoColumns struct {
	Id         string // 系统自增
	RoomId     string // 录制房间 ID。
	Duration   string // 文件时长，单位：ms。
	BeginTime  string // 直播开始时间
	FileUrl    string // 文件访问 URL。第三方存储为七牛云或阿里云 Vod 时不返回。
	FileSize   string // 文件大小，单位：字节。
	CreateTime string //
	UpdateTime string //
}

// liveCloudVideoColumns holds the columns for table live_cloud_video.
var liveCloudVideoColumns = LiveCloudVideoColumns{
	Id:         "id",
	RoomId:     "room_id",
	Duration:   "duration",
	BeginTime:  "begin_time",
	FileUrl:    "file_url",
	FileSize:   "file_size",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewLiveCloudVideoDao creates and returns a new DAO object for table data access.
func NewLiveCloudVideoDao() *LiveCloudVideoDao {
	return &LiveCloudVideoDao{
		group:   "default",
		table:   "live_cloud_video",
		columns: liveCloudVideoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiveCloudVideoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiveCloudVideoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiveCloudVideoDao) Columns() LiveCloudVideoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiveCloudVideoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiveCloudVideoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiveCloudVideoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

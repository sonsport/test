// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomActivityBannerDao is the data access object for table room_activity_banner.
type RoomActivityBannerDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns RoomActivityBannerColumns // columns contains all the column names of Table for convenient usage.
}

// RoomActivityBannerColumns defines and stores column names for table room_activity_banner.
type RoomActivityBannerColumns struct {
	Id           string //
	Icon         string // 图片
	ActivityType string // 活动类型 1 首冲 2 h5
	Link         string // 链接地址
	State        string // 状态1启用
	ShowType     string // 展示类型，0 全部，1 未充值用户 2 充值用户 3 主播
}

// roomActivityBannerColumns holds the columns for table room_activity_banner.
var roomActivityBannerColumns = RoomActivityBannerColumns{
	Id:           "id",
	Icon:         "icon",
	ActivityType: "activity_type",
	Link:         "link",
	State:        "state",
	ShowType:     "show_type",
}

// NewRoomActivityBannerDao creates and returns a new DAO object for table data access.
func NewRoomActivityBannerDao() *RoomActivityBannerDao {
	return &RoomActivityBannerDao{
		group:   "default",
		table:   "room_activity_banner",
		columns: roomActivityBannerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomActivityBannerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomActivityBannerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomActivityBannerDao) Columns() RoomActivityBannerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomActivityBannerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomActivityBannerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomActivityBannerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

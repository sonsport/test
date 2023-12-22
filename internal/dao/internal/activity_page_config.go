// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityPageConfigDao is the data access object for table activity_page_config.
type ActivityPageConfigDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ActivityPageConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityPageConfigColumns defines and stores column names for table activity_page_config.
type ActivityPageConfigColumns struct {
	Id           string //
	AcId         string // 活动Id activity_config 自增id
	ResUrl       string // 资源地址
	Duration     string // 持续时间 单位秒 >0才倒计时
	LinkAddress  string // 链接地址
	ShowStyle    string // 0无配置 1全屏 2半屏 后续其他扩展
	PageType     string // 页面类别：启动页(start_up)、首页(home_page)、直播间(room_live)、后续扩展
	PageLocation string // 页面位置：启动页：只有一个位置不用配置，首页：弹窗（pop_up）、浮窗（float）；直播间：弹窗（pop_up）、浮窗（float）、底部图标（bottom_icon）、直播间-礼物（gift_top）；
	StartTime    string // 活动开始时间
	EndTime      string // 活动结束时间
	CreateTime   string //
	UpdateTime   string //
}

// activityPageConfigColumns holds the columns for table activity_page_config.
var activityPageConfigColumns = ActivityPageConfigColumns{
	Id:           "id",
	AcId:         "ac_id",
	ResUrl:       "res_url",
	Duration:     "duration",
	LinkAddress:  "link_address",
	ShowStyle:    "show_style",
	PageType:     "page_type",
	PageLocation: "page_location",
	StartTime:    "start_time",
	EndTime:      "end_time",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewActivityPageConfigDao creates and returns a new DAO object for table data access.
func NewActivityPageConfigDao() *ActivityPageConfigDao {
	return &ActivityPageConfigDao{
		group:   "default",
		table:   "activity_page_config",
		columns: activityPageConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityPageConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityPageConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityPageConfigDao) Columns() ActivityPageConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityPageConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityPageConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityPageConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

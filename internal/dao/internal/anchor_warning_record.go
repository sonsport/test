// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorWarningRecordDao is the data access object for table anchor_warning_record.
type AnchorWarningRecordDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns AnchorWarningRecordColumns // columns contains all the column names of Table for convenient usage.
}

// AnchorWarningRecordColumns defines and stores column names for table anchor_warning_record.
type AnchorWarningRecordColumns struct {
	Id          string // 自增id
	Icon        string // 警告图片证据
	Appeal      string // 申诉
	State       string // 0待申诉，1申诉中，2成功，3失败
	UserId      string // 主播id
	RoomId      string // 直播间id
	WarnType    string // 警告类型
	Reason      string // 警告原因
	WarnBigType string //
	AdminId     string // 提醒大类
	CreateTime  string //
	UpdateTime  string //
}

// anchorWarningRecordColumns holds the columns for table anchor_warning_record.
var anchorWarningRecordColumns = AnchorWarningRecordColumns{
	Id:          "id",
	Icon:        "icon",
	Appeal:      "appeal",
	State:       "state",
	UserId:      "user_id",
	RoomId:      "room_id",
	WarnType:    "warn_type",
	Reason:      "reason",
	WarnBigType: "warn_big_type",
	AdminId:     "admin_id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewAnchorWarningRecordDao creates and returns a new DAO object for table data access.
func NewAnchorWarningRecordDao() *AnchorWarningRecordDao {
	return &AnchorWarningRecordDao{
		group:   "default",
		table:   "anchor_warning_record",
		columns: anchorWarningRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AnchorWarningRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AnchorWarningRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AnchorWarningRecordDao) Columns() AnchorWarningRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AnchorWarningRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AnchorWarningRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AnchorWarningRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

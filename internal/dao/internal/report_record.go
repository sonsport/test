// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReportRecordDao is the data access object for table report_record.
type ReportRecordDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ReportRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ReportRecordColumns defines and stores column names for table report_record.
type ReportRecordColumns struct {
	Id           string // 自增Id
	UserId       string // 用户Id
	ReportType   string // 举报类型：1直播间举报、2用户、3动态、4、评论，后续扩展
	ReportObject string // 举报对象 1直播间举报存储直播间Id，动态举报存储动态Id
	ContentType  string // 举报内容类型，1侮辱谩骂，2色情低俗，3政治谣言，4其他
	Describe     string // 举报用户描述
	State        string // 状态：0待处理 1已处理 后续扩展
	CreateTime   string //
	UpdateTime   string //
}

// reportRecordColumns holds the columns for table report_record.
var reportRecordColumns = ReportRecordColumns{
	Id:           "id",
	UserId:       "user_id",
	ReportType:   "report_type",
	ReportObject: "report_object",
	ContentType:  "content_type",
	Describe:     "describe",
	State:        "state",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewReportRecordDao creates and returns a new DAO object for table data access.
func NewReportRecordDao() *ReportRecordDao {
	return &ReportRecordDao{
		group:   "default",
		table:   "report_record",
		columns: reportRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReportRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReportRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReportRecordDao) Columns() ReportRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReportRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReportRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReportRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

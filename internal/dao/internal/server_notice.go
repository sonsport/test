// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerNoticeDao is the data access object for table server_notice.
type ServerNoticeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ServerNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// ServerNoticeColumns defines and stores column names for table server_notice.
type ServerNoticeColumns struct {
	Id            string // 系统自增
	ServerId      string // 运营号id
	UserId        string // 用户Id
	CreateAdminId string // 通知创建者Id  0系统创建
	NoticeType    string // 通知类型：1系统星级调整 2人工星级调整  后续其他扩展
	BeforeContent string // 之前内容
	AfterContent  string // 之后内容
	SystemRemark  string // 系统备注说明
	MessageType   string // 消息说明类别：系统星级下：0系统星级保持、1系统星级升级、2系统星级降级    人工星级调整  1人工星级降级 2人工星级降级  后续扩展
	State         string // 状态 0未处理 1已知晓 2已处理
	Operator      string // 最后操作人
	ServerRemark  string // 运营号备注
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// serverNoticeColumns holds the columns for table server_notice.
var serverNoticeColumns = ServerNoticeColumns{
	Id:            "id",
	ServerId:      "server_id",
	UserId:        "user_id",
	CreateAdminId: "create_admin_id",
	NoticeType:    "notice_type",
	BeforeContent: "before_content",
	AfterContent:  "after_content",
	SystemRemark:  "system_remark",
	MessageType:   "message_type",
	State:         "state",
	Operator:      "operator",
	ServerRemark:  "server_remark",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewServerNoticeDao creates and returns a new DAO object for table data access.
func NewServerNoticeDao() *ServerNoticeDao {
	return &ServerNoticeDao{
		group:   "default",
		table:   "server_notice",
		columns: serverNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerNoticeDao) Columns() ServerNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

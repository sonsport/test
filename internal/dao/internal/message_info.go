// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageInfoDao is the data access object for table message_info.
type MessageInfoDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns MessageInfoColumns // columns contains all the column names of Table for convenient usage.
}

// MessageInfoColumns defines and stores column names for table message_info.
type MessageInfoColumns struct {
	Id               string // 自增id
	SenderId         string // 发送用户Id
	ReceiverId       string // 接收人
	ConversationType string // 会话类型（1直播间消息、2私信消息）
	MessageInfo      string // 消息内容
	MessageType      string // 消息类型
	MessageSource    string // 消息来源
	SendTime         string // 发送时间
}

// messageInfoColumns holds the columns for table message_info.
var messageInfoColumns = MessageInfoColumns{
	Id:               "id",
	SenderId:         "sender_id",
	ReceiverId:       "receiver_id",
	ConversationType: "conversation_type",
	MessageInfo:      "message_info",
	MessageType:      "message_type",
	MessageSource:    "message_source",
	SendTime:         "send_time",
}

// NewMessageInfoDao creates and returns a new DAO object for table data access.
func NewMessageInfoDao() *MessageInfoDao {
	return &MessageInfoDao{
		group:   "default",
		table:   "message_info",
		columns: messageInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MessageInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MessageInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MessageInfoDao) Columns() MessageInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MessageInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MessageInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MessageInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

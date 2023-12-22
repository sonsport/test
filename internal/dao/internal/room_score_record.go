// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomScoreRecordDao is the data access object for table room_score_record.
type RoomScoreRecordDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns RoomScoreRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RoomScoreRecordColumns defines and stores column names for table room_score_record.
type RoomScoreRecordColumns struct {
	Id            string //
	UserId        string //
	RoomId        string //
	OperationType string // 操作类型
	OperationId   string // 操作id
	Score         string // 排序分值
	BeforeScore   string // before排序分值
	AfterScore    string // after排序分值
	CreateTime    string //
	UpdateTime    string //
}

// roomScoreRecordColumns holds the columns for table room_score_record.
var roomScoreRecordColumns = RoomScoreRecordColumns{
	Id:            "id",
	UserId:        "user_id",
	RoomId:        "room_id",
	OperationType: "operation_type",
	OperationId:   "operation_id",
	Score:         "score",
	BeforeScore:   "before_score",
	AfterScore:    "after_score",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewRoomScoreRecordDao creates and returns a new DAO object for table data access.
func NewRoomScoreRecordDao() *RoomScoreRecordDao {
	return &RoomScoreRecordDao{
		group:   "default",
		table:   "room_score_record",
		columns: roomScoreRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomScoreRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomScoreRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomScoreRecordDao) Columns() RoomScoreRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomScoreRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomScoreRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomScoreRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

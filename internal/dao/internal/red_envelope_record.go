// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RedEnvelopeRecordDao is the data access object for table red_envelope_record.
type RedEnvelopeRecordDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns RedEnvelopeRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RedEnvelopeRecordColumns defines and stores column names for table red_envelope_record.
type RedEnvelopeRecordColumns struct {
	Id            string //
	RId           string // 红包id
	UserId        string // 用户id
	RoomId        string // 房间id
	Diamond       string // 红包钻石
	BeforeDiamond string // 红包之前余额
	AfterDiamond  string // 红包之后余额
	CreateTime    string //
	UpdateTime    string //
}

// redEnvelopeRecordColumns holds the columns for table red_envelope_record.
var redEnvelopeRecordColumns = RedEnvelopeRecordColumns{
	Id:            "id",
	RId:           "r_id",
	UserId:        "user_id",
	RoomId:        "room_id",
	Diamond:       "diamond",
	BeforeDiamond: "before_diamond",
	AfterDiamond:  "after_diamond",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewRedEnvelopeRecordDao creates and returns a new DAO object for table data access.
func NewRedEnvelopeRecordDao() *RedEnvelopeRecordDao {
	return &RedEnvelopeRecordDao{
		group:   "default",
		table:   "red_envelope_record",
		columns: redEnvelopeRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RedEnvelopeRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RedEnvelopeRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RedEnvelopeRecordDao) Columns() RedEnvelopeRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RedEnvelopeRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RedEnvelopeRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RedEnvelopeRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

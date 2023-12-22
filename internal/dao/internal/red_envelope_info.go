// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RedEnvelopeInfoDao is the data access object for table red_envelope_info.
type RedEnvelopeInfoDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns RedEnvelopeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// RedEnvelopeInfoColumns defines and stores column names for table red_envelope_info.
type RedEnvelopeInfoColumns struct {
	Id                 string //
	UserId             string // 用户id
	AnchorUserId       string // 主播id
	RoomId             string // 房间id
	RedEnvelopeType    string // 红包类型 1 房间红包 2 全站红包
	RedEnvelopeCount   string // 红包个数
	RedEnvelopeDiamond string // 红包钻石
	SpentDiamond       string // 已抢钻石
	BeginTime          string // 开始时间
	EndTime            string // 结束时间
	CreateTime         string //
	UpdateTime         string //
}

// redEnvelopeInfoColumns holds the columns for table red_envelope_info.
var redEnvelopeInfoColumns = RedEnvelopeInfoColumns{
	Id:                 "id",
	UserId:             "user_id",
	AnchorUserId:       "anchor_user_id",
	RoomId:             "room_id",
	RedEnvelopeType:    "red_envelope_type",
	RedEnvelopeCount:   "red_envelope_count",
	RedEnvelopeDiamond: "red_envelope_diamond",
	SpentDiamond:       "spent_diamond",
	BeginTime:          "begin_time",
	EndTime:            "end_time",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewRedEnvelopeInfoDao creates and returns a new DAO object for table data access.
func NewRedEnvelopeInfoDao() *RedEnvelopeInfoDao {
	return &RedEnvelopeInfoDao{
		group:   "default",
		table:   "red_envelope_info",
		columns: redEnvelopeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RedEnvelopeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RedEnvelopeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RedEnvelopeInfoDao) Columns() RedEnvelopeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RedEnvelopeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RedEnvelopeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RedEnvelopeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

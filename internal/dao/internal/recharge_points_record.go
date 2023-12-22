// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RechargePointsRecordDao is the data access object for table recharge_points_record.
type RechargePointsRecordDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns RechargePointsRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RechargePointsRecordColumns defines and stores column names for table recharge_points_record.
type RechargePointsRecordColumns struct {
	Id           string //
	UserId       string // 用户id
	Points       string // 积分
	PointsType   string // 积分类型
	BeforePoints string // 之前积分
	AfterPoints  string // 之后积分
	BeforeRank   string // 之前排名
	AfterRank    string // 之后排名
	OrderId      string // 订单id
	CreateTime   string //
	UpdateTime   string //
}

// rechargePointsRecordColumns holds the columns for table recharge_points_record.
var rechargePointsRecordColumns = RechargePointsRecordColumns{
	Id:           "id",
	UserId:       "user_id",
	Points:       "points",
	PointsType:   "points_type",
	BeforePoints: "before_points",
	AfterPoints:  "after_points",
	BeforeRank:   "before_rank",
	AfterRank:    "after_rank",
	OrderId:      "order_id",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewRechargePointsRecordDao creates and returns a new DAO object for table data access.
func NewRechargePointsRecordDao() *RechargePointsRecordDao {
	return &RechargePointsRecordDao{
		group:   "default",
		table:   "recharge_points_record",
		columns: rechargePointsRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RechargePointsRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RechargePointsRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RechargePointsRecordDao) Columns() RechargePointsRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RechargePointsRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RechargePointsRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RechargePointsRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

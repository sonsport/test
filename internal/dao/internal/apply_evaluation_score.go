// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyEvaluationScoreDao is the data access object for table apply_evaluation_score.
type ApplyEvaluationScoreDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns ApplyEvaluationScoreColumns // columns contains all the column names of Table for convenient usage.
}

// ApplyEvaluationScoreColumns defines and stores column names for table apply_evaluation_score.
type ApplyEvaluationScoreColumns struct {
	Id         string //
	UserId     string // 用户id
	Type       string // 弹窗类型 1:真 2:假
	Operate    string // 用户操作 1:确认 2:取消
	Score      string // 用户评分
	Action     string // 评分触发动作 1:举报拉黑 2:直播卡顿 3:频繁切换直接间 4:送礼物 5:主播关播
	DeviceId   string // 设备号
	CreateTime string //
	UpdateTime string //
}

// applyEvaluationScoreColumns holds the columns for table apply_evaluation_score.
var applyEvaluationScoreColumns = ApplyEvaluationScoreColumns{
	Id:         "id",
	UserId:     "user_id",
	Type:       "type",
	Operate:    "operate",
	Score:      "score",
	Action:     "action",
	DeviceId:   "device_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewApplyEvaluationScoreDao creates and returns a new DAO object for table data access.
func NewApplyEvaluationScoreDao() *ApplyEvaluationScoreDao {
	return &ApplyEvaluationScoreDao{
		group:   "default",
		table:   "apply_evaluation_score",
		columns: applyEvaluationScoreColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApplyEvaluationScoreDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ApplyEvaluationScoreDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ApplyEvaluationScoreDao) Columns() ApplyEvaluationScoreColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ApplyEvaluationScoreDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApplyEvaluationScoreDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApplyEvaluationScoreDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

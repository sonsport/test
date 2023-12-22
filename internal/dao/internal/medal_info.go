// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MedalInfoDao is the data access object for table medal_info.
type MedalInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns MedalInfoColumns // columns contains all the column names of Table for convenient usage.
}

// MedalInfoColumns defines and stores column names for table medal_info.
type MedalInfoColumns struct {
	Id              string //
	MedalIcon       string // 勋章图片
	MedalGreyIcon   string // 勋章灰色图片
	MedalNameConfig string // 勋章名称配置json
	MedalDescConfig string // 勋章描述配置
	ConditionsType  string // 获取类型
	BeginTime       string // 获取勋章开始时间
	EndTime         string // 获取勋章结束时间
	MedalStates     string // 勋章状态 1
	MedalSort       string // 排序
	Remark          string // 备注
	CreateTime      string //
	UpdateTime      string //
}

// medalInfoColumns holds the columns for table medal_info.
var medalInfoColumns = MedalInfoColumns{
	Id:              "id",
	MedalIcon:       "medal_icon",
	MedalGreyIcon:   "medal_grey_icon",
	MedalNameConfig: "medal_name_config",
	MedalDescConfig: "medal_desc_config",
	ConditionsType:  "conditions_type",
	BeginTime:       "begin_time",
	EndTime:         "end_time",
	MedalStates:     "medal_states",
	MedalSort:       "medal_sort",
	Remark:          "remark",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
}

// NewMedalInfoDao creates and returns a new DAO object for table data access.
func NewMedalInfoDao() *MedalInfoDao {
	return &MedalInfoDao{
		group:   "default",
		table:   "medal_info",
		columns: medalInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MedalInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MedalInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MedalInfoDao) Columns() MedalInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MedalInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MedalInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MedalInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

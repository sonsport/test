// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShareNewUserRecordDao is the data access object for table share_new_user_record.
type ShareNewUserRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ShareNewUserRecordColumns // columns contains all the column names of Table for convenient usage.
}

// ShareNewUserRecordColumns defines and stores column names for table share_new_user_record.
type ShareNewUserRecordColumns struct {
	Id                string //
	UserId            string //
	OwnerUid          string // 分享者uid
	Type              string // 1:新用户注册；2:新用户充值
	Status            string // 0:初始化；1:待加币；2:加币成功
	OwnerOrderId      string // 分享者加币订单号
	OwnerDiamondNum   string // 分享者获得的钻石数
	NewUserOrderId    string // 新用户加币订单
	NewUserDiamondNum string // 新用户获得的钻石数
	CodeId            string // 分享ID
	RegisterTime      string // 注册时间
	RechargeNum       string // 充值的虚拟币数量
	WatchLiveTime     string // 观看直播时长
	Ip                string // 登陆地IP
	Smid              string // 设备ID
	DeviceName        string // 设备型号
	UpdateTime        string //
	CreateTime        string //
}

// shareNewUserRecordColumns holds the columns for table share_new_user_record.
var shareNewUserRecordColumns = ShareNewUserRecordColumns{
	Id:                "id",
	UserId:            "user_id",
	OwnerUid:          "owner_uid",
	Type:              "type",
	Status:            "status",
	OwnerOrderId:      "owner_order_id",
	OwnerDiamondNum:   "owner_diamond_num",
	NewUserOrderId:    "new_user_order_id",
	NewUserDiamondNum: "new_user_diamond_num",
	CodeId:            "code_id",
	RegisterTime:      "register_time",
	RechargeNum:       "recharge_num",
	WatchLiveTime:     "watch_live_time",
	Ip:                "ip",
	Smid:              "smid",
	DeviceName:        "device_name",
	UpdateTime:        "update_time",
	CreateTime:        "create_time",
}

// NewShareNewUserRecordDao creates and returns a new DAO object for table data access.
func NewShareNewUserRecordDao() *ShareNewUserRecordDao {
	return &ShareNewUserRecordDao{
		group:   "default",
		table:   "share_new_user_record",
		columns: shareNewUserRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShareNewUserRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShareNewUserRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShareNewUserRecordDao) Columns() ShareNewUserRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShareNewUserRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShareNewUserRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShareNewUserRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

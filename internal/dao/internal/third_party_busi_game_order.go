// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ThirdPartyBusiGameOrderDao is the data access object for table third_party_busi_game_order.
type ThirdPartyBusiGameOrderDao struct {
	table   string                         // table is the underlying table name of the DAO.
	group   string                         // group is the database configuration group name of current DAO.
	columns ThirdPartyBusiGameOrderColumns // columns contains all the column names of Table for convenient usage.
}

// ThirdPartyBusiGameOrderColumns defines and stores column names for table third_party_busi_game_order.
type ThirdPartyBusiGameOrderColumns struct {
	Id          string //
	OrderType   string // 1   : 增加钻石, 2:扣减钻石
	BusiId      string // 第三方业务唯一标识
	HimiUid     string // uid
	CoinNum     string // 订单虚拟币额度
	BusiUid     string // 第三方业务用户唯一标识
	OrderStatus string // 订单状态 0-下单,1-执行成功
	RoomId      string // 房间id
	CreateTime  string //
	UpdateTime  string //
}

// thirdPartyBusiGameOrderColumns holds the columns for table third_party_busi_game_order.
var thirdPartyBusiGameOrderColumns = ThirdPartyBusiGameOrderColumns{
	Id:          "id",
	OrderType:   "order_type",
	BusiId:      "busi_id",
	HimiUid:     "himi_uid",
	CoinNum:     "coin_num",
	BusiUid:     "busi_uid",
	OrderStatus: "order_status",
	RoomId:      "room_id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewThirdPartyBusiGameOrderDao creates and returns a new DAO object for table data access.
func NewThirdPartyBusiGameOrderDao() *ThirdPartyBusiGameOrderDao {
	return &ThirdPartyBusiGameOrderDao{
		group:   "default",
		table:   "third_party_busi_game_order",
		columns: thirdPartyBusiGameOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ThirdPartyBusiGameOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ThirdPartyBusiGameOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ThirdPartyBusiGameOrderDao) Columns() ThirdPartyBusiGameOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ThirdPartyBusiGameOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ThirdPartyBusiGameOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ThirdPartyBusiGameOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

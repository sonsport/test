// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PayoutInfoDao is the data access object for table payout_info.
type PayoutInfoDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PayoutInfoColumns // columns contains all the column names of Table for convenient usage.
}

// PayoutInfoColumns defines and stores column names for table payout_info.
type PayoutInfoColumns struct {
	PayoutId        string // 代付id
	UserId          string // 用户Id 自增
	PayoutType      string // 转账类型  1 主播转账 2 公会长提现  3公会长主播提现
	OrderNo         string // 系统订单号
	TradeNo         string // 交易订单号
	OperateUserId   string // 操作员用户id
	CurrencyCode    string // 付款币种编码
	CurrencyFee     string // 付款币种总额，单位分
	PaidFee         string // 实际交易金额，单位分
	ProcessFee      string // 手续费，单位分
	BankCode        string // 银行代码
	PayeeAccount    string // 收款人银行账号
	PayeeName       string // 收款人姓名
	PayeeEmail      string // 收款人邮箱
	PayeePhone      string // 收款人电话
	TransferDate    string // 放款时间
	TransferPurpose string // 用途
	PayType         string // 支付类型
	TransferStatus  string // 转账状态，0.未转账，1.转账中，2.已转账，3.转账失败，4.已退款，5.取消转账
	CreateTime      string // 创建时间
	UpdateTime      string // 修改时间
	Remark          string // 备注
}

// payoutInfoColumns holds the columns for table payout_info.
var payoutInfoColumns = PayoutInfoColumns{
	PayoutId:        "payout_id",
	UserId:          "user_id",
	PayoutType:      "payout_type",
	OrderNo:         "order_no",
	TradeNo:         "trade_no",
	OperateUserId:   "operate_user_id",
	CurrencyCode:    "currency_code",
	CurrencyFee:     "currency_fee",
	PaidFee:         "paid_fee",
	ProcessFee:      "process_fee",
	BankCode:        "bank_code",
	PayeeAccount:    "payee_account",
	PayeeName:       "payee_name",
	PayeeEmail:      "payee_email",
	PayeePhone:      "payee_phone",
	TransferDate:    "transfer_date",
	TransferPurpose: "transfer_purpose",
	PayType:         "pay_type",
	TransferStatus:  "transfer_status",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	Remark:          "remark",
}

// NewPayoutInfoDao creates and returns a new DAO object for table data access.
func NewPayoutInfoDao() *PayoutInfoDao {
	return &PayoutInfoDao{
		group:   "default",
		table:   "payout_info",
		columns: payoutInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PayoutInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PayoutInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PayoutInfoDao) Columns() PayoutInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PayoutInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PayoutInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PayoutInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

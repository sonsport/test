// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBankDao is the data access object for table user_bank.
type UserBankDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserBankColumns // columns contains all the column names of Table for convenient usage.
}

// UserBankColumns defines and stores column names for table user_bank.
type UserBankColumns struct {
	Id           string //
	UserId       string // 用户id
	IdCardNum    string // 身份证
	IdCardName   string // 身份证姓名
	IdCardFront  string // 身份证正面
	IdCardBack   string // 身份证背面
	BankName     string // 账户名
	BankAccounts string // 银行账号
	BankPerson   string // 银行卡持有人
	CreatedTime  string //
	UpdatedTime  string //
}

// userBankColumns holds the columns for table user_bank.
var userBankColumns = UserBankColumns{
	Id:           "id",
	UserId:       "user_id",
	IdCardNum:    "id_card_num",
	IdCardName:   "id_card_name",
	IdCardFront:  "id_card_front",
	IdCardBack:   "id_card_back",
	BankName:     "bank_name",
	BankAccounts: "bank_accounts",
	BankPerson:   "bank_person",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
}

// NewUserBankDao creates and returns a new DAO object for table data access.
func NewUserBankDao() *UserBankDao {
	return &UserBankDao{
		group:   "default",
		table:   "user_bank",
		columns: userBankColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserBankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserBankDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserBankDao) Columns() UserBankColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserBankDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserBankDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserBankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

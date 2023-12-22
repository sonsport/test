// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserBank is the golang structure of table user_bank for DAO operations like Where/Data.
type UserBank struct {
	g.Meta       `orm:"table:user_bank, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户id
	IdCardNum    interface{} // 身份证
	IdCardName   interface{} // 身份证姓名
	IdCardFront  interface{} // 身份证正面
	IdCardBack   interface{} // 身份证背面
	BankName     interface{} // 账户名
	BankAccounts interface{} // 银行账号
	BankPerson   interface{} // 银行卡持有人
	CreatedTime  interface{} //
	UpdatedTime  interface{} //
}

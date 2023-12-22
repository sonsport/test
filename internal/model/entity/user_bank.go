// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserBank is the golang structure for table user_bank.
type UserBank struct {
	Id           int64  `json:"id"           description:""`
	UserId       int64  `json:"userId"       description:"用户id"`
	IdCardNum    string `json:"idCardNum"    description:"身份证"`
	IdCardName   string `json:"idCardName"   description:"身份证姓名"`
	IdCardFront  string `json:"idCardFront"  description:"身份证正面"`
	IdCardBack   string `json:"idCardBack"   description:"身份证背面"`
	BankName     string `json:"bankName"     description:"账户名"`
	BankAccounts string `json:"bankAccounts" description:"银行账号"`
	BankPerson   string `json:"bankPerson"   description:"银行卡持有人"`
	CreatedTime  int64  `json:"createdTime"  description:""`
	UpdatedTime  int64  `json:"updatedTime"  description:""`
}

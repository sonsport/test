// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PayoutInfo is the golang structure for table payout_info.
type PayoutInfo struct {
	PayoutId        int64  `json:"payoutId"        description:"代付id"`
	UserId          int64  `json:"userId"          description:"用户Id 自增"`
	PayoutType      int    `json:"payoutType"      description:"转账类型  1 主播转账 2 公会长提现  3公会长主播提现"`
	OrderNo         string `json:"orderNo"         description:"系统订单号"`
	TradeNo         string `json:"tradeNo"         description:"交易订单号"`
	OperateUserId   int64  `json:"operateUserId"   description:"操作员用户id"`
	CurrencyCode    string `json:"currencyCode"    description:"付款币种编码"`
	CurrencyFee     int64  `json:"currencyFee"     description:"付款币种总额，单位分"`
	PaidFee         int64  `json:"paidFee"         description:"实际交易金额，单位分"`
	ProcessFee      int64  `json:"processFee"      description:"手续费，单位分"`
	BankCode        string `json:"bankCode"        description:"银行代码"`
	PayeeAccount    string `json:"payeeAccount"    description:"收款人银行账号"`
	PayeeName       string `json:"payeeName"       description:"收款人姓名"`
	PayeeEmail      string `json:"payeeEmail"      description:"收款人邮箱"`
	PayeePhone      string `json:"payeePhone"      description:"收款人电话"`
	TransferDate    string `json:"transferDate"    description:"放款时间"`
	TransferPurpose string `json:"transferPurpose" description:"用途"`
	PayType         int    `json:"payType"         description:"支付类型"`
	TransferStatus  int    `json:"transferStatus"  description:"转账状态，0.未转账，1.转账中，2.已转账，3.转账失败，4.已退款，5.取消转账"`
	CreateTime      int64  `json:"createTime"      description:"创建时间"`
	UpdateTime      int64  `json:"updateTime"      description:"修改时间"`
	Remark          string `json:"remark"          description:"备注"`
}

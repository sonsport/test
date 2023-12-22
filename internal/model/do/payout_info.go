// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PayoutInfo is the golang structure of table payout_info for DAO operations like Where/Data.
type PayoutInfo struct {
	g.Meta          `orm:"table:payout_info, do:true"`
	PayoutId        interface{} // 代付id
	UserId          interface{} // 用户Id 自增
	PayoutType      interface{} // 转账类型  1 主播转账 2 公会长提现  3公会长主播提现
	OrderNo         interface{} // 系统订单号
	TradeNo         interface{} // 交易订单号
	OperateUserId   interface{} // 操作员用户id
	CurrencyCode    interface{} // 付款币种编码
	CurrencyFee     interface{} // 付款币种总额，单位分
	PaidFee         interface{} // 实际交易金额，单位分
	ProcessFee      interface{} // 手续费，单位分
	BankCode        interface{} // 银行代码
	PayeeAccount    interface{} // 收款人银行账号
	PayeeName       interface{} // 收款人姓名
	PayeeEmail      interface{} // 收款人邮箱
	PayeePhone      interface{} // 收款人电话
	TransferDate    interface{} // 放款时间
	TransferPurpose interface{} // 用途
	PayType         interface{} // 支付类型
	TransferStatus  interface{} // 转账状态，0.未转账，1.转账中，2.已转账，3.转账失败，4.已退款，5.取消转账
	CreateTime      interface{} // 创建时间
	UpdateTime      interface{} // 修改时间
	Remark          interface{} // 备注
}

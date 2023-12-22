// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AgentBalanceRecord is the golang structure of table agent_balance_record for DAO operations like Where/Data.
type AgentBalanceRecord struct {
	g.Meta         `orm:"table:agent_balance_record, do:true"`
	Id             interface{} //
	UserId         interface{} //
	TargetUserId   interface{} //
	RecordType     interface{} // 消费类型 1支出 2收入
	OnePrice       interface{} // 售出单价
	CostPrice      interface{} // 成本单价
	Diamonds       interface{} // 钻石数额
	BeforeChange   interface{} // 变更之前余额
	AfterChange    interface{} // 变更之后余额
	LinkId         interface{} // 引起余额变更的id
	CertificateUrl interface{} // 凭证url
	PaymentId      interface{} // 支付id
	PaymentName    interface{} // 支付名称
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

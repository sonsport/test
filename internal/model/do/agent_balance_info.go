// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AgentBalanceInfo is the golang structure of table agent_balance_info for DAO operations like Where/Data.
type AgentBalanceInfo struct {
	g.Meta            `orm:"table:agent_balance_info, do:true"`
	UserId            interface{} //
	TotalFee          interface{} // 总充值 印尼盾 单位：分
	TotalSaleFee      interface{} // 总销售 印尼盾 单位：分
	TotalBuyDiamonds  interface{} // 总购买钻石
	TotalSaleDiamonds interface{} // 总销售钻石
	RemainDiamonds    interface{} // 剩余钻石数
	PayPassword       interface{} // 支付密码
	CostPrice         interface{} // 成本单价
	IsLock            interface{} // 是否锁定
	AgentStates       interface{} // 代理状态 1正常
	AgentRank         interface{} // 代理等级 1 2 3
	AgentBanTime      interface{} // 禁止转账截止时间
	CreateTime        interface{} //
	UpdateTime        interface{} //
}

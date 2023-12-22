// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AgentBalanceInfo is the golang structure for table agent_balance_info.
type AgentBalanceInfo struct {
	UserId            int64   `json:"userId"            description:""`
	TotalFee          int64   `json:"totalFee"          description:"总充值 印尼盾 单位：分"`
	TotalSaleFee      int64   `json:"totalSaleFee"      description:"总销售 印尼盾 单位：分"`
	TotalBuyDiamonds  int64   `json:"totalBuyDiamonds"  description:"总购买钻石"`
	TotalSaleDiamonds int64   `json:"totalSaleDiamonds" description:"总销售钻石"`
	RemainDiamonds    int64   `json:"remainDiamonds"    description:"剩余钻石数"`
	PayPassword       string  `json:"payPassword"       description:"支付密码"`
	CostPrice         float64 `json:"costPrice"         description:"成本单价"`
	IsLock            int     `json:"isLock"            description:"是否锁定"`
	AgentStates       int     `json:"agentStates"       description:"代理状态 1正常"`
	AgentRank         int     `json:"agentRank"         description:"代理等级 1 2 3"`
	AgentBanTime      int64   `json:"agentBanTime"      description:"禁止转账截止时间"`
	CreateTime        int64   `json:"createTime"        description:""`
	UpdateTime        int64   `json:"updateTime"        description:""`
}

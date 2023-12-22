// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AgentBalanceRecord is the golang structure for table agent_balance_record.
type AgentBalanceRecord struct {
	Id             int64   `json:"id"             description:""`
	UserId         int64   `json:"userId"         description:""`
	TargetUserId   int64   `json:"targetUserId"   description:""`
	RecordType     int     `json:"recordType"     description:"消费类型 1支出 2收入"`
	OnePrice       float64 `json:"onePrice"       description:"售出单价"`
	CostPrice      float64 `json:"costPrice"      description:"成本单价"`
	Diamonds       int     `json:"diamonds"       description:"钻石数额"`
	BeforeChange   uint64  `json:"beforeChange"   description:"变更之前余额"`
	AfterChange    uint64  `json:"afterChange"    description:"变更之后余额"`
	LinkId         uint64  `json:"linkId"         description:"引起余额变更的id"`
	CertificateUrl string  `json:"certificateUrl" description:"凭证url"`
	PaymentId      int     `json:"paymentId"      description:"支付id"`
	PaymentName    string  `json:"paymentName"    description:"支付名称"`
	CreateTime     int64   `json:"createTime"     description:""`
	UpdateTime     int64   `json:"updateTime"     description:""`
}

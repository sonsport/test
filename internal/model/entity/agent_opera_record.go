// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AgentOperaRecord is the golang structure for table agent_opera_record.
type AgentOperaRecord struct {
	Id         int64  `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:""`
	OperaIp    string `json:"operaIp"    description:"操作ip"`
	LinkId     string `json:"linkId"     description:"关联id"`
	UserAgent  string `json:"userAgent"  description:"系统类型"`
	OperaType  int    `json:"operaType"  description:"操作类型 1 登录 2 转账 3 充值"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

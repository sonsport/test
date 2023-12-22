// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AgentDetailRecord is the golang structure for table agent_detail_record.
type AgentDetailRecord struct {
	Id            int64  `json:"id"            description:""`
	UserId        int64  `json:"userId"        description:"用户Id"`
	OperateUserId int    `json:"operateUserId" description:"操作用户id"`
	ChangeType    int    `json:"changeType"    description:"触发类型 1 后台 2 系统"`
	UpdateType    int    `json:"updateType"    description:"更新类型 1 代理等级 2 支付密码 3 禁止转账截止时间 4 是否锁定 5 代理状态"`
	BeforeChange  string `json:"beforeChange"  description:"变更之前值"`
	AfterChange   string `json:"afterChange"   description:"变更之后值"`
	Remarks       string `json:"remarks"       description:"更新备注"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

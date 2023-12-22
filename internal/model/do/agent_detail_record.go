// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AgentDetailRecord is the golang structure of table agent_detail_record for DAO operations like Where/Data.
type AgentDetailRecord struct {
	g.Meta        `orm:"table:agent_detail_record, do:true"`
	Id            interface{} //
	UserId        interface{} // 用户Id
	OperateUserId interface{} // 操作用户id
	ChangeType    interface{} // 触发类型 1 后台 2 系统
	UpdateType    interface{} // 更新类型 1 代理等级 2 支付密码 3 禁止转账截止时间 4 是否锁定 5 代理状态
	BeforeChange  interface{} // 变更之前值
	AfterChange   interface{} // 变更之后值
	Remarks       interface{} // 更新备注
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

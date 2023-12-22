// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AgentOperaRecord is the golang structure of table agent_opera_record for DAO operations like Where/Data.
type AgentOperaRecord struct {
	g.Meta     `orm:"table:agent_opera_record, do:true"`
	Id         interface{} //
	UserId     interface{} //
	OperaIp    interface{} // 操作ip
	LinkId     interface{} // 关联id
	UserAgent  interface{} // 系统类型
	OperaType  interface{} // 操作类型 1 登录 2 转账 3 充值
	CreateTime interface{} //
	UpdateTime interface{} //
}

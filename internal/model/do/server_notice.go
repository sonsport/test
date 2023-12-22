// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ServerNotice is the golang structure of table server_notice for DAO operations like Where/Data.
type ServerNotice struct {
	g.Meta        `orm:"table:server_notice, do:true"`
	Id            interface{} // 系统自增
	ServerId      interface{} // 运营号id
	UserId        interface{} // 用户Id
	CreateAdminId interface{} // 通知创建者Id  0系统创建
	NoticeType    interface{} // 通知类型：1系统星级调整 2人工星级调整  后续其他扩展
	BeforeContent interface{} // 之前内容
	AfterContent  interface{} // 之后内容
	SystemRemark  interface{} // 系统备注说明
	MessageType   interface{} // 消息说明类别：系统星级下：0系统星级保持、1系统星级升级、2系统星级降级    人工星级调整  1人工星级降级 2人工星级降级  后续扩展
	State         interface{} // 状态 0未处理 1已知晓 2已处理
	Operator      interface{} // 最后操作人
	ServerRemark  interface{} // 运营号备注
	CreateTime    interface{} // 创建时间
	UpdateTime    interface{} // 更新时间
}

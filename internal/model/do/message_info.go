// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MessageInfo is the golang structure of table message_info for DAO operations like Where/Data.
type MessageInfo struct {
	g.Meta           `orm:"table:message_info, do:true"`
	Id               interface{} // 自增id
	SenderId         interface{} // 发送用户Id
	ReceiverId       interface{} // 接收人
	ConversationType interface{} // 会话类型（1直播间消息、2私信消息）
	MessageInfo      interface{} // 消息内容
	MessageType      interface{} // 消息类型
	MessageSource    interface{} // 消息来源
	SendTime         interface{} // 发送时间
}

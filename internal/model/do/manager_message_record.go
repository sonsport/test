// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ManagerMessageRecord is the golang structure of table manager_message_record for DAO operations like Where/Data.
type ManagerMessageRecord struct {
	g.Meta     `orm:"table:manager_message_record, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 用户id
	CreateTime interface{} //
	UpdateTime interface{} //
	AdminId    interface{} // 管理员id
	Message    interface{} // 发送的消息
	WhatsApp   interface{} // whatsapp号码
	State      interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ManagerMessageRecord is the golang structure for table manager_message_record.
type ManagerMessageRecord struct {
	Id         int64  `json:"id"         description:"自增id"`
	UserId     int64  `json:"userId"     description:"用户id"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
	AdminId    int64  `json:"adminId"    description:"管理员id"`
	Message    string `json:"message"    description:"发送的消息"`
	WhatsApp   string `json:"whatsApp"   description:"whatsapp号码"`
	State      int    `json:"state"      description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MessageInfo is the golang structure for table message_info.
type MessageInfo struct {
	Id               int64  `json:"id"               description:"自增id"`
	SenderId         int64  `json:"senderId"         description:"发送用户Id"`
	ReceiverId       string `json:"receiverId"       description:"接收人"`
	ConversationType int    `json:"conversationType" description:"会话类型（1直播间消息、2私信消息）"`
	MessageInfo      string `json:"messageInfo"      description:"消息内容"`
	MessageType      int    `json:"messageType"      description:"消息类型"`
	MessageSource    string `json:"messageSource"    description:"消息来源"`
	SendTime         int64  `json:"sendTime"         description:"发送时间"`
}

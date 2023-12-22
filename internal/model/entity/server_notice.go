// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ServerNotice is the golang structure for table server_notice.
type ServerNotice struct {
	Id            uint64 `json:"id"            description:"系统自增"`
	ServerId      int64  `json:"serverId"      description:"运营号id"`
	UserId        int64  `json:"userId"        description:"用户Id"`
	CreateAdminId int64  `json:"createAdminId" description:"通知创建者Id  0系统创建"`
	NoticeType    int    `json:"noticeType"    description:"通知类型：1系统星级调整 2人工星级调整  后续其他扩展"`
	BeforeContent string `json:"beforeContent" description:"之前内容"`
	AfterContent  string `json:"afterContent"  description:"之后内容"`
	SystemRemark  string `json:"systemRemark"  description:"系统备注说明"`
	MessageType   int    `json:"messageType"   description:"消息说明类别：系统星级下：0系统星级保持、1系统星级升级、2系统星级降级    人工星级调整  1人工星级降级 2人工星级降级  后续扩展"`
	State         int    `json:"state"         description:"状态 0未处理 1已知晓 2已处理"`
	Operator      int64  `json:"operator"      description:"最后操作人"`
	ServerRemark  string `json:"serverRemark"  description:"运营号备注"`
	CreateTime    int64  `json:"createTime"    description:"创建时间"`
	UpdateTime    int64  `json:"updateTime"    description:"更新时间"`
}

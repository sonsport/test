// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UpdateRecord is the golang structure for table update_record.
type UpdateRecord struct {
	Id         int64  `json:"id"         description:"主键id"`
	UserId     int64  `json:"userId"     description:"用户id"`
	AdminId    int64  `json:"adminId"    description:"管理员id"`
	UpdateType int    `json:"updateType" description:"调整类型，1星级，2公会，3标签，4备注"`
	BeforeBody string `json:"beforeBody" description:"修改前内容"`
	AfterBody  string `json:"afterBody"  description:"修改后内容"`
	Remark     string `json:"remark"     description:"修改后内容"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

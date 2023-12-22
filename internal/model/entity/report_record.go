// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ReportRecord is the golang structure for table report_record.
type ReportRecord struct {
	Id           uint64 `json:"id"           description:"自增Id"`
	UserId       int64  `json:"userId"       description:"用户Id"`
	ReportType   int    `json:"reportType"   description:"举报类型：1直播间举报、2用户、3动态、4、评论，后续扩展"`
	ReportObject string `json:"reportObject" description:"举报对象 1直播间举报存储直播间Id，动态举报存储动态Id"`
	ContentType  uint   `json:"contentType"  description:"举报内容类型，1侮辱谩骂，2色情低俗，3政治谣言，4其他"`
	Describe     string `json:"describe"     description:"举报用户描述"`
	State        int    `json:"state"        description:"状态：0待处理 1已处理 后续扩展"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

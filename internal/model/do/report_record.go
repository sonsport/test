// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ReportRecord is the golang structure of table report_record for DAO operations like Where/Data.
type ReportRecord struct {
	g.Meta       `orm:"table:report_record, do:true"`
	Id           interface{} // 自增Id
	UserId       interface{} // 用户Id
	ReportType   interface{} // 举报类型：1直播间举报、2用户、3动态、4、评论，后续扩展
	ReportObject interface{} // 举报对象 1直播间举报存储直播间Id，动态举报存储动态Id
	ContentType  interface{} // 举报内容类型，1侮辱谩骂，2色情低俗，3政治谣言，4其他
	Describe     interface{} // 举报用户描述
	State        interface{} // 状态：0待处理 1已处理 后续扩展
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

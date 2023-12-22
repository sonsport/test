// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeedCommentMark is the golang structure of table feed_comment_mark for DAO operations like Where/Data.
type FeedCommentMark struct {
	g.Meta        `orm:"table:feed_comment_mark, do:true"`
	Id            interface{} // 自增Id
	UserId        interface{} // 用户Id
	FeedCommentId interface{} // 评论Id
	Mark          interface{} // 标记类型：1喜欢、 后续扩展
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

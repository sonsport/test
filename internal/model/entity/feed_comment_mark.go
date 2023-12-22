// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedCommentMark is the golang structure for table feed_comment_mark.
type FeedCommentMark struct {
	Id            uint64 `json:"id"            description:"自增Id"`
	UserId        int64  `json:"userId"        description:"用户Id"`
	FeedCommentId int64  `json:"feedCommentId" description:"评论Id"`
	Mark          int    `json:"mark"          description:"标记类型：1喜欢、 后续扩展"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

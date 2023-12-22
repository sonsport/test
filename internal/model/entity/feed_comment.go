// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedComment is the golang structure for table feed_comment.
type FeedComment struct {
	Id             uint64 `json:"id"             description:"自增Id"`
	UserId         int64  `json:"userId"         description:"评论用户Id"`
	FeedId         uint64 `json:"feedId"         description:"动态Id"`
	ParentId       int64  `json:"parentId"       description:"父级评论ID，  0标识对一级动态回复、 >0标识对评论回复只存储一级回复的Id（自增id）"`
	ToUserId       int64  `json:"toUserId"       description:"A回复了B，这里存储B的用户ID ，一级回复id这里存储0"`
	CommentId      int64  `json:"commentId"      description:"该条评论针对那个评论的回复，一级回复存储0，针对评论的的回复存储回复评论的ID，该数据只是存储"`
	ReplyContext   string `json:"replyContext"   description:"评论/回复内容"`
	ReplyLikeCount uint   `json:"replyLikeCount" description:"评论点赞数"`
	State          int    `json:"state"          description:"0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
}

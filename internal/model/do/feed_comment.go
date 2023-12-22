// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeedComment is the golang structure of table feed_comment for DAO operations like Where/Data.
type FeedComment struct {
	g.Meta         `orm:"table:feed_comment, do:true"`
	Id             interface{} // 自增Id
	UserId         interface{} // 评论用户Id
	FeedId         interface{} // 动态Id
	ParentId       interface{} // 父级评论ID，  0标识对一级动态回复、 >0标识对评论回复只存储一级回复的Id（自增id）
	ToUserId       interface{} // A回复了B，这里存储B的用户ID ，一级回复id这里存储0
	CommentId      interface{} // 该条评论针对那个评论的回复，一级回复存储0，针对评论的的回复存储回复评论的ID，该数据只是存储
	ReplyContext   interface{} // 评论/回复内容
	ReplyLikeCount interface{} // 评论点赞数
	State          interface{} // 0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

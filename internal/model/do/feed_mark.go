// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeedMark is the golang structure of table feed_mark for DAO operations like Where/Data.
type FeedMark struct {
	g.Meta     `orm:"table:feed_mark, do:true"`
	Id         interface{} // 自增Id
	UserId     interface{} // 用户Id
	FeedId     interface{} // 动态Id
	Mark       interface{} // 标记类型：1喜欢、2不感兴趣 后续扩展
	CreateTime interface{} //
	UpdateTime interface{} //
}

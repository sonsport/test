// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeedTopic is the golang structure of table feed_topic for DAO operations like Where/Data.
type FeedTopic struct {
	g.Meta     `orm:"table:feed_topic, do:true"`
	Id         interface{} // 自增Id
	TopicId    interface{} // 话题Id
	FeedId     interface{} // 动态Id
	CreateTime interface{} //
	UpdateTime interface{} //
}

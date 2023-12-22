// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TopicFollow is the golang structure of table topic_follow for DAO operations like Where/Data.
type TopicFollow struct {
	g.Meta     `orm:"table:topic_follow, do:true"`
	Id         interface{} // 自增Id
	TopicId    interface{} // 话题Id
	UserId     interface{} // 用户Id
	CreateTime interface{} //
	UpdateTime interface{} //
}

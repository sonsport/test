// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TopicInfo is the golang structure of table topic_info for DAO operations like Where/Data.
type TopicInfo struct {
	g.Meta         `orm:"table:topic_info, do:true"`
	Id             interface{} // 自增Id
	UserId         interface{} // 用户Id
	Title          interface{} // 话题标题
	CoverImg       interface{} // 封面图
	Score          interface{} // 值越大 热度越高
	TopicType      interface{} // 0普通话题 1官方 2审核模式话题
	RecommendScore interface{} // 推荐 默认是0  排序倒序 值越大越靠前
	State          interface{} // 状态：0待审核 1审核通过 2审核不通过
	CreateTime     interface{} //
	UpdateTime     interface{} //
}

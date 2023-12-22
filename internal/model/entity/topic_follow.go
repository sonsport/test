// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TopicFollow is the golang structure for table topic_follow.
type TopicFollow struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	TopicId    int64  `json:"topicId"    description:"话题Id"`
	UserId     int64  `json:"userId"     description:"用户Id"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

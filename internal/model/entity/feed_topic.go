// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedTopic is the golang structure for table feed_topic.
type FeedTopic struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	TopicId    int64  `json:"topicId"    description:"话题Id"`
	FeedId     int64  `json:"feedId"     description:"动态Id"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

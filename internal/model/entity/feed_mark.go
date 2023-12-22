// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedMark is the golang structure for table feed_mark.
type FeedMark struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	UserId     int64  `json:"userId"     description:"用户Id"`
	FeedId     int64  `json:"feedId"     description:"动态Id"`
	Mark       int    `json:"mark"       description:"标记类型：1喜欢、2不感兴趣 后续扩展"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

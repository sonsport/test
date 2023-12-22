// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TopicInfo is the golang structure for table topic_info.
type TopicInfo struct {
	Id             uint64 `json:"id"             description:"自增Id"`
	UserId         int64  `json:"userId"         description:"用户Id"`
	Title          string `json:"title"          description:"话题标题"`
	CoverImg       string `json:"coverImg"       description:"封面图"`
	Score          int64  `json:"score"          description:"值越大 热度越高"`
	TopicType      int    `json:"topicType"      description:"0普通话题 1官方 2审核模式话题"`
	RecommendScore int    `json:"recommendScore" description:"推荐 默认是0  排序倒序 值越大越靠前"`
	State          int    `json:"state"          description:"状态：0待审核 1审核通过 2审核不通过"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
}

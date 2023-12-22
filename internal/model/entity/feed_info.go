// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FeedInfo is the golang structure for table feed_info.
type FeedInfo struct {
	Id                uint64 `json:"id"                description:"自增Id"`
	UserId            int64  `json:"userId"            description:"用户Id"`
	Text              string `json:"text"              description:"动态内容"`
	Images            string `json:"images"            description:"动态图片组，采用数组存储,目前规定数组长度为9，即最多9张图片"`
	LikeCount         uint   `json:"likeCount"         description:"点赞数"`
	State             int    `json:"state"             description:"0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除"`
	Score             int    `json:"score"             description:"动态排序权重"`
	CreateTime        int64  `json:"createTime"        description:""`
	UpdateTime        int64  `json:"updateTime"        description:""`
	Video             string `json:"video"             description:"动态视频"`
	VideoPicture      string `json:"videoPicture"      description:"视频封面"`
	Type              int    `json:"type"              description:"动态类型，1图片、2视频"`
	UploadVideoWidth  int    `json:"uploadVideoWidth"  description:""`
	UploadVideoHeight int    `json:"uploadVideoHeight" description:""`
	IsVote            int    `json:"isVote"            description:"是否pk动态"`
	OptionOne         string `json:"optionOne"         description:"pk选项1"`
	OptionTwo         string `json:"optionTwo"         description:"pk选项2"`
	SystemOs          int    `json:"systemOs"          description:"发动态的设备类型"`
	AppChannel        string `json:"appChannel"        description:"渠道包"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeedInfo is the golang structure of table feed_info for DAO operations like Where/Data.
type FeedInfo struct {
	g.Meta            `orm:"table:feed_info, do:true"`
	Id                interface{} // 自增Id
	UserId            interface{} // 用户Id
	Text              interface{} // 动态内容
	Images            interface{} // 动态图片组，采用数组存储,目前规定数组长度为9，即最多9张图片
	LikeCount         interface{} // 点赞数
	State             interface{} // 0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除
	Score             interface{} // 动态排序权重
	CreateTime        interface{} //
	UpdateTime        interface{} //
	Video             interface{} // 动态视频
	VideoPicture      interface{} // 视频封面
	Type              interface{} // 动态类型，1图片、2视频
	UploadVideoWidth  interface{} //
	UploadVideoHeight interface{} //
	IsVote            interface{} // 是否pk动态
	OptionOne         interface{} // pk选项1
	OptionTwo         interface{} // pk选项2
	SystemOs          interface{} // 发动态的设备类型
	AppChannel        interface{} // 渠道包
}

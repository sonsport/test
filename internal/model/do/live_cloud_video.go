// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LiveCloudVideo is the golang structure of table live_cloud_video for DAO operations like Where/Data.
type LiveCloudVideo struct {
	g.Meta     `orm:"table:live_cloud_video, do:true"`
	Id         interface{} // 系统自增
	RoomId     interface{} // 录制房间 ID。
	Duration   interface{} // 文件时长，单位：ms。
	BeginTime  interface{} // 直播开始时间
	FileUrl    interface{} // 文件访问 URL。第三方存储为七牛云或阿里云 Vod 时不返回。
	FileSize   interface{} // 文件大小，单位：字节。
	CreateTime interface{} //
	UpdateTime interface{} //
}

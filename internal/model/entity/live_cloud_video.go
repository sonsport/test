// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LiveCloudVideo is the golang structure for table live_cloud_video.
type LiveCloudVideo struct {
	Id         uint64 `json:"id"         description:"系统自增"`
	RoomId     string `json:"roomId"     description:"录制房间 ID。"`
	Duration   int64  `json:"duration"   description:"文件时长，单位：ms。"`
	BeginTime  int64  `json:"beginTime"  description:"直播开始时间"`
	FileUrl    string `json:"fileUrl"    description:"文件访问 URL。第三方存储为七牛云或阿里云 Vod 时不返回。"`
	FileSize   int64  `json:"fileSize"   description:"文件大小，单位：字节。"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

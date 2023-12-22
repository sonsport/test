// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LiveCloudRecord is the golang structure for table live_cloud_record.
type LiveCloudRecord struct {
	Id         uint64 `json:"id"         description:"系统自增"`
	TaskId     string `json:"taskId"     description:"录制任务 ID，长度固定为 16 个字节的字符串。"`
	GuildId    int64  `json:"guildId"    description:"公会id"`
	RoomId     string `json:"roomId"     description:"录制房间 ID"`
	UserId     string `json:"userId"     description:"录制流对应的推流用户 ID（混流时，为 mix_output_stream_id）。"`
	Status     int    `json:"status"     description:"录制状态，1开始录制，2录制结束"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

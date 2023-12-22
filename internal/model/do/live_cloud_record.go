// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LiveCloudRecord is the golang structure of table live_cloud_record for DAO operations like Where/Data.
type LiveCloudRecord struct {
	g.Meta     `orm:"table:live_cloud_record, do:true"`
	Id         interface{} // 系统自增
	TaskId     interface{} // 录制任务 ID，长度固定为 16 个字节的字符串。
	GuildId    interface{} // 公会id
	RoomId     interface{} // 录制房间 ID
	UserId     interface{} // 录制流对应的推流用户 ID（混流时，为 mix_output_stream_id）。
	Status     interface{} // 录制状态，1开始录制，2录制结束
	CreateTime interface{} //
	UpdateTime interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ApplyEvaluationScore is the golang structure for table apply_evaluation_score.
type ApplyEvaluationScore struct {
	Id         uint64 `json:"id"         description:""`
	UserId     uint64 `json:"userId"     description:"用户id"`
	Type       uint   `json:"type"       description:"弹窗类型 1:真 2:假"`
	Operate    uint   `json:"operate"    description:"用户操作 1:确认 2:取消"`
	Score      uint   `json:"score"      description:"用户评分"`
	Action     uint   `json:"action"     description:"评分触发动作 1:举报拉黑 2:直播卡顿 3:频繁切换直接间 4:送礼物 5:主播关播"`
	DeviceId   string `json:"deviceId"   description:"设备号"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

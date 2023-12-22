// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ApplyEvaluationScore is the golang structure of table apply_evaluation_score for DAO operations like Where/Data.
type ApplyEvaluationScore struct {
	g.Meta     `orm:"table:apply_evaluation_score, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户id
	Type       interface{} // 弹窗类型 1:真 2:假
	Operate    interface{} // 用户操作 1:确认 2:取消
	Score      interface{} // 用户评分
	Action     interface{} // 评分触发动作 1:举报拉黑 2:直播卡顿 3:频繁切换直接间 4:送礼物 5:主播关播
	DeviceId   interface{} // 设备号
	CreateTime interface{} //
	UpdateTime interface{} //
}

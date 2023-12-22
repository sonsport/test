// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QuestConfig is the golang structure of table quest_config for DAO operations like Where/Data.
type QuestConfig struct {
	g.Meta           `orm:"table:quest_config, do:true"`
	Id               interface{} // id
	QuestNameJson    interface{} // 任务名称
	QuestIcon        interface{} // 任务图标
	QuestStartTime   interface{} // 任务开始时间
	QuestEndTime     interface{} // 任务结束时间
	QuestKey         interface{} // 任务唯一标识 key唯一则表示套娃任务 区分等级
	QuestParentId    interface{} // 任务上级id -1 为一级任务
	QuestRank        interface{} // 任务等级 套娃任务等级递增
	QuestCycle       interface{} // 任务周期 1 每日 2 每周
	QuestTheme       interface{} // 任务主题 1 日常任务 2 主播任务 3 用户钻石任务 4 用户活跃宝箱
	QuestUserScope   interface{} // 任务用户范围 1 所有用户 2 星级主播 3 公会长
	QuestType        interface{} // 任务类型 common.QuestTypeSignIn
	QuestAchieved    interface{} // 任务达成值 1 次数任务代表1次 计时为1分钟
	IsReward         interface{} // 任务是否有奖励 1 有
	QuestIsMust      interface{} // 任务是否必须 1 是
	QuestSort        interface{} // 任务排序
	QuestStates      interface{} // 任务状态
	CreateTime       interface{} // createdAt
	UpdateTime       interface{} // updatedAt
	AutomaticRelease interface{} // 自动发放 1 是
}

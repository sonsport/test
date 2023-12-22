// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QuestUserDetail is the golang structure of table quest_user_detail for DAO operations like Where/Data.
type QuestUserDetail struct {
	g.Meta         `orm:"table:quest_user_detail, do:true"`
	Id             interface{} // id
	QuestId        interface{} // 任务id
	QuestUserId    interface{} // 参与用户id
	ExpirationTime interface{} // 过期时间
	QuestAchieved  interface{} // 任务达成值 1 次数任务代表1次 计时为1分钟
	IsComplete     interface{} // 是否完成
	CompleteTime   interface{} // 完成时间
	QuestKey       interface{} // 任务唯一标识 key唯一则表示套娃任务 区分等级
	QuestRank      interface{} // 任务等级 套娃任务等级递增
	CreateTime     interface{} // createdAt
	UpdateTime     interface{} // updatedAt
}

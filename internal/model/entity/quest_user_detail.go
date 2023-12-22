// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QuestUserDetail is the golang structure for table quest_user_detail.
type QuestUserDetail struct {
	Id             uint64 `json:"id"             description:"id"`
	QuestId        int64  `json:"questId"        description:"任务id"`
	QuestUserId    int64  `json:"questUserId"    description:"参与用户id"`
	ExpirationTime int64  `json:"expirationTime" description:"过期时间"`
	QuestAchieved  int    `json:"questAchieved"  description:"任务达成值 1 次数任务代表1次 计时为1分钟"`
	IsComplete     int    `json:"isComplete"     description:"是否完成"`
	CompleteTime   int    `json:"completeTime"   description:"完成时间"`
	QuestKey       string `json:"questKey"       description:"任务唯一标识 key唯一则表示套娃任务 区分等级"`
	QuestRank      int    `json:"questRank"      description:"任务等级 套娃任务等级递增"`
	CreateTime     int64  `json:"createTime"     description:"createdAt"`
	UpdateTime     int64  `json:"updateTime"     description:"updatedAt"`
}

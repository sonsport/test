// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// QuestConfig is the golang structure for table quest_config.
type QuestConfig struct {
	Id               uint64 `json:"id"               description:"id"`
	QuestNameJson    string `json:"questNameJson"    description:"任务名称"`
	QuestIcon        string `json:"questIcon"        description:"任务图标"`
	QuestStartTime   int64  `json:"questStartTime"   description:"任务开始时间"`
	QuestEndTime     int64  `json:"questEndTime"     description:"任务结束时间"`
	QuestKey         string `json:"questKey"         description:"任务唯一标识 key唯一则表示套娃任务 区分等级"`
	QuestParentId    int64  `json:"questParentId"    description:"任务上级id -1 为一级任务"`
	QuestRank        int    `json:"questRank"        description:"任务等级 套娃任务等级递增"`
	QuestCycle       int    `json:"questCycle"       description:"任务周期 1 每日 2 每周"`
	QuestTheme       int    `json:"questTheme"       description:"任务主题 1 日常任务 2 主播任务 3 用户钻石任务 4 用户活跃宝箱"`
	QuestUserScope   int    `json:"questUserScope"   description:"任务用户范围 1 所有用户 2 星级主播 3 公会长"`
	QuestType        int    `json:"questType"        description:"任务类型 common.QuestTypeSignIn"`
	QuestAchieved    int    `json:"questAchieved"    description:"任务达成值 1 次数任务代表1次 计时为1分钟"`
	IsReward         int    `json:"isReward"         description:"任务是否有奖励 1 有"`
	QuestIsMust      int    `json:"questIsMust"      description:"任务是否必须 1 是"`
	QuestSort        int    `json:"questSort"        description:"任务排序"`
	QuestStates      int    `json:"questStates"      description:"任务状态"`
	CreateTime       int64  `json:"createTime"       description:"createdAt"`
	UpdateTime       int64  `json:"updateTime"       description:"updatedAt"`
	AutomaticRelease int    `json:"automaticRelease" description:"自动发放 1 是"`
}

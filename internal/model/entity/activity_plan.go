// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityPlan is the golang structure for table activity_plan.
type ActivityPlan struct {
	Id           int    `json:"id"           description:""`
	ActivityType int    `json:"activityType" description:"活动类别,榜单活动"`
	Title        string `json:"title"        description:"活动标题"`
	Icon         string `json:"icon"         description:"图片"`
	Explain      string `json:"explain"      description:"活动说明"`
	StartTime    int64  `json:"startTime"    description:"活动开始时间"`
	EndTime      int64  `json:"endTime"      description:"活动结束时间"`
	TimeLevel    int    `json:"timeLevel"    description:"活动数据类型,（自然周、自然天、时）也可以自定义时间类型（最小颗粒度：小时）、周期内（all）：0"`
	RewardCycle  int    `json:"rewardCycle"  description:"奖励发放周期配置：小时、天、周、默认0不奖励、统一指定时间奖励"`
	Object       string `json:"object"       description:"奖励对象：主播/公会长/用户 或者多个"`
	Source       string `json:"source"       description:"奖励数据源：rank_diamonds、rank_coin、rank_gift(根据礼物id判断)……,以此类推；根据key标记调用对应的封装接口 不配置默认为不调用数据源（特殊活动/非模板活动）"`
	AppChannel   string `json:"appChannel"   description:"能参加活动的渠道"`
	Sort         int    `json:"sort"         description:"排序"`
	State        int    `json:"state"        description:"活动状态是否有效，1有效，2无效"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

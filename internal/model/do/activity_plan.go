// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityPlan is the golang structure of table activity_plan for DAO operations like Where/Data.
type ActivityPlan struct {
	g.Meta       `orm:"table:activity_plan, do:true"`
	Id           interface{} //
	ActivityType interface{} // 活动类别,榜单活动
	Title        interface{} // 活动标题
	Icon         interface{} // 图片
	Explain      interface{} // 活动说明
	StartTime    interface{} // 活动开始时间
	EndTime      interface{} // 活动结束时间
	TimeLevel    interface{} // 活动数据类型,（自然周、自然天、时）也可以自定义时间类型（最小颗粒度：小时）、周期内（all）：0
	RewardCycle  interface{} // 奖励发放周期配置：小时、天、周、默认0不奖励、统一指定时间奖励
	Object       interface{} // 奖励对象：主播/公会长/用户 或者多个
	Source       interface{} // 奖励数据源：rank_diamonds、rank_coin、rank_gift(根据礼物id判断)……,以此类推；根据key标记调用对应的封装接口 不配置默认为不调用数据源（特殊活动/非模板活动）
	AppChannel   interface{} // 能参加活动的渠道
	Sort         interface{} // 排序
	State        interface{} // 活动状态是否有效，1有效，2无效
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

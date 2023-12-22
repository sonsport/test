// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RevenueStatistics is the golang structure of table revenue_statistics for DAO operations like Where/Data.
type RevenueStatistics struct {
	g.Meta                        `orm:"table:revenue_statistics, do:true"`
	StatisticsId                  interface{} //
	BeginTime                     interface{} // 开始时间
	EndTime                       interface{} // 结束时间
	TimeLevel                     interface{} // 时间维度：1.每小时，2.每天
	TotalExpendDiamonds           interface{} // 总消费（流水时间在时间范围内的支出类型的钻石数量）
	TotalRevenueDiamonds          interface{} // 总收入（流水时间在时间范围内的支出类型的钻石数量）
	AllUserCount                  interface{} // 总用户量
	NewUserCount                  interface{} // 每天日活--认证时间在今天
	ActiveUsersCount              interface{} // 每天日活--认证时间在今天
	ActiveAnchorCount             interface{} // 主播日活
	AuthAnchorCount               interface{} // 认证主播数
	AuthStarAnchorCount           interface{} // 认证标星主播数
	LiveAnchorCount               interface{} // 主播开播数
	LiveStarAnchorCount           interface{} // 标星主播开播数
	EffectiveLiveAnchorCount      interface{} // 有效开播数
	NewEffectiveLiveAnchorCount   interface{} // 新增有效开播数
	ReachEffectiveLiveAnchorCount interface{} // 达标开播数
	TotalFee                      interface{} // 总收入(印尼盾)
	IncrementalFee                interface{} // 充值用户数
	AllOrderCount                 interface{} // 总订单数
	SuccessOrderCount             interface{} // 付费订单数
	TotalPaidUserCount            interface{} // 总支付人数
	IncrementalPaidUserCount      interface{} // 新用户支付人数
	YesterdayNewUserCount         interface{} // 昨天新增用户
	YesterdayUserTodayCount       interface{} // 昨天注册用户在今天登录数
	WeekNewUserCount              interface{} // 上周新增用户
	WeekUserTodayCount            interface{} // 上周注册用户在今天登录数
	MonthNewUserCount             interface{} // 上月新增用户
	MonthUserTodayCount           interface{} // 上月注册用户在今天登录数
	YesterdayNewAnchorCount       interface{} // 昨天新增认证主播
	YesterdayAnchorTodayCount     interface{} // 昨天注册主播在今天登录数
	WeekNewAnchorCount            interface{} // 上周新增认证主播
	WeekAnchorTodayCount          interface{} // 上周注册主播在今天登录数
	MonthNewAnchorCount           interface{} // 上月新增认证主播
	MonthAnchorTodayCount         interface{} // 上月注册主播在今天登录数
	YesterdayNewPayUserCount      interface{} // 昨天新增付费用户
	YesterdayPayUserTodayCount    interface{} // 昨天付费在今天登录数
	WeekNewPayUserCount           interface{} // 上周新增付费用户
	WeekPayUserTodayCount         interface{} // 上周付费在今天登录数
	MonthNewPayUserCount          interface{} // 上月新增付费用户
	MonthPayUserTodayCount        interface{} // 上月付费在今天登录数
	RemainDiamonds                interface{} // 剩余钻石数
	CreateTime                    interface{} //
	AppName                       interface{} // app名称
	AppChannel                    interface{} // app渠道
	AdvertisingCosts              interface{} // 广告费
}

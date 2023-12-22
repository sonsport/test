// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RevenueStatistics is the golang structure for table revenue_statistics.
type RevenueStatistics struct {
	StatisticsId                  int64   `json:"statisticsId"                  description:""`
	BeginTime                     int64   `json:"beginTime"                     description:"开始时间"`
	EndTime                       int64   `json:"endTime"                       description:"结束时间"`
	TimeLevel                     int     `json:"timeLevel"                     description:"时间维度：1.每小时，2.每天"`
	TotalExpendDiamonds           int64   `json:"totalExpendDiamonds"           description:"总消费（流水时间在时间范围内的支出类型的钻石数量）"`
	TotalRevenueDiamonds          int64   `json:"totalRevenueDiamonds"          description:"总收入（流水时间在时间范围内的支出类型的钻石数量）"`
	AllUserCount                  int64   `json:"allUserCount"                  description:"总用户量"`
	NewUserCount                  int64   `json:"newUserCount"                  description:"每天日活--认证时间在今天"`
	ActiveUsersCount              int64   `json:"activeUsersCount"              description:"每天日活--认证时间在今天"`
	ActiveAnchorCount             int64   `json:"activeAnchorCount"             description:"主播日活"`
	AuthAnchorCount               int64   `json:"authAnchorCount"               description:"认证主播数"`
	AuthStarAnchorCount           int64   `json:"authStarAnchorCount"           description:"认证标星主播数"`
	LiveAnchorCount               int64   `json:"liveAnchorCount"               description:"主播开播数"`
	LiveStarAnchorCount           int64   `json:"liveStarAnchorCount"           description:"标星主播开播数"`
	EffectiveLiveAnchorCount      int64   `json:"effectiveLiveAnchorCount"      description:"有效开播数"`
	NewEffectiveLiveAnchorCount   int64   `json:"newEffectiveLiveAnchorCount"   description:"新增有效开播数"`
	ReachEffectiveLiveAnchorCount int64   `json:"reachEffectiveLiveAnchorCount" description:"达标开播数"`
	TotalFee                      int64   `json:"totalFee"                      description:"总收入(印尼盾)"`
	IncrementalFee                int64   `json:"incrementalFee"                description:"充值用户数"`
	AllOrderCount                 int64   `json:"allOrderCount"                 description:"总订单数"`
	SuccessOrderCount             int64   `json:"successOrderCount"             description:"付费订单数"`
	TotalPaidUserCount            int64   `json:"totalPaidUserCount"            description:"总支付人数"`
	IncrementalPaidUserCount      int64   `json:"incrementalPaidUserCount"      description:"新用户支付人数"`
	YesterdayNewUserCount         int     `json:"yesterdayNewUserCount"         description:"昨天新增用户"`
	YesterdayUserTodayCount       int     `json:"yesterdayUserTodayCount"       description:"昨天注册用户在今天登录数"`
	WeekNewUserCount              int     `json:"weekNewUserCount"              description:"上周新增用户"`
	WeekUserTodayCount            int     `json:"weekUserTodayCount"            description:"上周注册用户在今天登录数"`
	MonthNewUserCount             int     `json:"monthNewUserCount"             description:"上月新增用户"`
	MonthUserTodayCount           int     `json:"monthUserTodayCount"           description:"上月注册用户在今天登录数"`
	YesterdayNewAnchorCount       int     `json:"yesterdayNewAnchorCount"       description:"昨天新增认证主播"`
	YesterdayAnchorTodayCount     int     `json:"yesterdayAnchorTodayCount"     description:"昨天注册主播在今天登录数"`
	WeekNewAnchorCount            int     `json:"weekNewAnchorCount"            description:"上周新增认证主播"`
	WeekAnchorTodayCount          int     `json:"weekAnchorTodayCount"          description:"上周注册主播在今天登录数"`
	MonthNewAnchorCount           int     `json:"monthNewAnchorCount"           description:"上月新增认证主播"`
	MonthAnchorTodayCount         int     `json:"monthAnchorTodayCount"         description:"上月注册主播在今天登录数"`
	YesterdayNewPayUserCount      int     `json:"yesterdayNewPayUserCount"      description:"昨天新增付费用户"`
	YesterdayPayUserTodayCount    int     `json:"yesterdayPayUserTodayCount"    description:"昨天付费在今天登录数"`
	WeekNewPayUserCount           int     `json:"weekNewPayUserCount"           description:"上周新增付费用户"`
	WeekPayUserTodayCount         int     `json:"weekPayUserTodayCount"         description:"上周付费在今天登录数"`
	MonthNewPayUserCount          int     `json:"monthNewPayUserCount"          description:"上月新增付费用户"`
	MonthPayUserTodayCount        int     `json:"monthPayUserTodayCount"        description:"上月付费在今天登录数"`
	RemainDiamonds                uint64  `json:"remainDiamonds"                description:"剩余钻石数"`
	CreateTime                    int64   `json:"createTime"                    description:""`
	AppName                       string  `json:"appName"                       description:"app名称"`
	AppChannel                    string  `json:"appChannel"                    description:"app渠道"`
	AdvertisingCosts              float64 `json:"advertisingCosts"              description:"广告费"`
}

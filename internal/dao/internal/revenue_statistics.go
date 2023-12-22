// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RevenueStatisticsDao is the data access object for table revenue_statistics.
type RevenueStatisticsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns RevenueStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// RevenueStatisticsColumns defines and stores column names for table revenue_statistics.
type RevenueStatisticsColumns struct {
	StatisticsId                  string //
	BeginTime                     string // 开始时间
	EndTime                       string // 结束时间
	TimeLevel                     string // 时间维度：1.每小时，2.每天
	TotalExpendDiamonds           string // 总消费（流水时间在时间范围内的支出类型的钻石数量）
	TotalRevenueDiamonds          string // 总收入（流水时间在时间范围内的支出类型的钻石数量）
	AllUserCount                  string // 总用户量
	NewUserCount                  string // 每天日活--认证时间在今天
	ActiveUsersCount              string // 每天日活--认证时间在今天
	ActiveAnchorCount             string // 主播日活
	AuthAnchorCount               string // 认证主播数
	AuthStarAnchorCount           string // 认证标星主播数
	LiveAnchorCount               string // 主播开播数
	LiveStarAnchorCount           string // 标星主播开播数
	EffectiveLiveAnchorCount      string // 有效开播数
	NewEffectiveLiveAnchorCount   string // 新增有效开播数
	ReachEffectiveLiveAnchorCount string // 达标开播数
	TotalFee                      string // 总收入(印尼盾)
	IncrementalFee                string // 充值用户数
	AllOrderCount                 string // 总订单数
	SuccessOrderCount             string // 付费订单数
	TotalPaidUserCount            string // 总支付人数
	IncrementalPaidUserCount      string // 新用户支付人数
	YesterdayNewUserCount         string // 昨天新增用户
	YesterdayUserTodayCount       string // 昨天注册用户在今天登录数
	WeekNewUserCount              string // 上周新增用户
	WeekUserTodayCount            string // 上周注册用户在今天登录数
	MonthNewUserCount             string // 上月新增用户
	MonthUserTodayCount           string // 上月注册用户在今天登录数
	YesterdayNewAnchorCount       string // 昨天新增认证主播
	YesterdayAnchorTodayCount     string // 昨天注册主播在今天登录数
	WeekNewAnchorCount            string // 上周新增认证主播
	WeekAnchorTodayCount          string // 上周注册主播在今天登录数
	MonthNewAnchorCount           string // 上月新增认证主播
	MonthAnchorTodayCount         string // 上月注册主播在今天登录数
	YesterdayNewPayUserCount      string // 昨天新增付费用户
	YesterdayPayUserTodayCount    string // 昨天付费在今天登录数
	WeekNewPayUserCount           string // 上周新增付费用户
	WeekPayUserTodayCount         string // 上周付费在今天登录数
	MonthNewPayUserCount          string // 上月新增付费用户
	MonthPayUserTodayCount        string // 上月付费在今天登录数
	RemainDiamonds                string // 剩余钻石数
	CreateTime                    string //
	AppName                       string // app名称
	AppChannel                    string // app渠道
	AdvertisingCosts              string // 广告费
}

// revenueStatisticsColumns holds the columns for table revenue_statistics.
var revenueStatisticsColumns = RevenueStatisticsColumns{
	StatisticsId:                  "statistics_id",
	BeginTime:                     "begin_time",
	EndTime:                       "end_time",
	TimeLevel:                     "time_level",
	TotalExpendDiamonds:           "total_expend_diamonds",
	TotalRevenueDiamonds:          "total_revenue_diamonds",
	AllUserCount:                  "all_user_count",
	NewUserCount:                  "new_user_count",
	ActiveUsersCount:              "active_users_count",
	ActiveAnchorCount:             "active_anchor_count",
	AuthAnchorCount:               "auth_anchor_count",
	AuthStarAnchorCount:           "auth_star_anchor_count",
	LiveAnchorCount:               "live_anchor_count",
	LiveStarAnchorCount:           "live_star_anchor_count",
	EffectiveLiveAnchorCount:      "effective_live_anchor_count",
	NewEffectiveLiveAnchorCount:   "new_effective_live_anchor_count",
	ReachEffectiveLiveAnchorCount: "reach_effective_live_anchor_count",
	TotalFee:                      "total_fee",
	IncrementalFee:                "incremental_fee",
	AllOrderCount:                 "all_order_count",
	SuccessOrderCount:             "success_order_count",
	TotalPaidUserCount:            "total_paid_user_count",
	IncrementalPaidUserCount:      "incremental_paid_user_count",
	YesterdayNewUserCount:         "yesterday_new_user_count",
	YesterdayUserTodayCount:       "yesterday_user_today_count",
	WeekNewUserCount:              "week_new_user_count",
	WeekUserTodayCount:            "week_user_today_count",
	MonthNewUserCount:             "month_new_user_count",
	MonthUserTodayCount:           "month_user_today_count",
	YesterdayNewAnchorCount:       "yesterday_new_anchor_count",
	YesterdayAnchorTodayCount:     "yesterday_anchor_today_count",
	WeekNewAnchorCount:            "week_new_anchor_count",
	WeekAnchorTodayCount:          "week_anchor_today_count",
	MonthNewAnchorCount:           "month_new_anchor_count",
	MonthAnchorTodayCount:         "month_anchor_today_count",
	YesterdayNewPayUserCount:      "yesterday_new_pay_user_count",
	YesterdayPayUserTodayCount:    "yesterday_pay_user_today_count",
	WeekNewPayUserCount:           "week_new_pay_user_count",
	WeekPayUserTodayCount:         "week_pay_user_today_count",
	MonthNewPayUserCount:          "month_new_pay_user_count",
	MonthPayUserTodayCount:        "month_pay_user_today_count",
	RemainDiamonds:                "remain_diamonds",
	CreateTime:                    "create_time",
	AppName:                       "app_name",
	AppChannel:                    "app_channel",
	AdvertisingCosts:              "advertising_costs",
}

// NewRevenueStatisticsDao creates and returns a new DAO object for table data access.
func NewRevenueStatisticsDao() *RevenueStatisticsDao {
	return &RevenueStatisticsDao{
		group:   "default",
		table:   "revenue_statistics",
		columns: revenueStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RevenueStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RevenueStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RevenueStatisticsDao) Columns() RevenueStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RevenueStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RevenueStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RevenueStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

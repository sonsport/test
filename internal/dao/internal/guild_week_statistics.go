// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GuildWeekStatisticsDao is the data access object for table guild_week_statistics.
type GuildWeekStatisticsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns GuildWeekStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// GuildWeekStatisticsColumns defines and stores column names for table guild_week_statistics.
type GuildWeekStatisticsColumns struct {
	Id                         string // 自增id
	ServerId                   string // 运营号id
	Week                       string // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildId                    string // 公会Id 关联guild_info自增
	TotalAnchorNum             string // 总主播数
	TotalLiveTime              string // 总直播时长(秒)
	TotalLivePerson            string // 总开播人数
	NewAnchorNum               string // 新增主播人数
	NewAnchorLiveTime          string // 新增主播直播时长
	NewAnchorLivePerson        string // 新增主播开播人数
	StarAnchorNum              string // 总标星主播数
	StarAnchorLiveTime         string // 总标星主播总开播时长
	StarAnchorLivePerson       string // 总标星主播开播人数
	NewSettlementAnchorNum     string // 新增达标结算主播数
	SettlementAnchorNum        string // 达标结算主播数
	NewStarAnchorNum           string // 新增标星主播人数
	NewStarAnchorLiveTime      string // 新增标星主播总开播时长
	NewStarAnchorLivePerson    string // 新增标星主播开播人数
	SettlementAnchorLiveTime   string // 达标结算主播总直播时长
	SettlementIncome           string // 达标结算收益
	NewLivedUnsettlementPerson string // 新增开播未达标结算主播数
	LivedUnsettlementPerson    string // 已开播未结算主播数
	LivedUnsettlementLiveTime  string // 已开播未结算主播总开播时长
	WeekAwardDiamonds          string // 本周奖励钻石数
	Remarks                    string // 奖励钻石的备注
	InvitePersonNum            string // 邀新总人数
	TotalFee                   string // 拉新总冲值
	InviteNormalUserNum        string // 邀新普通用户数
	InviteNormalUserTotalFee   string // 邀新普通用户充值
	InviteAnchorNum            string // 邀新主播数
	InviteAnchorTotalFee       string // 邀新主播充值
	InviteMasterNum            string // 邀新公会数
	InviteMasterTotalFee       string // 邀新公会长充值
	EffectiveUserNum           string // 有效用户数
	EffectiveUserTotalFee      string // 有效用户金币
	CreateTime                 string // 创建时间
	UpdateTime                 string // 更新时间
}

// guildWeekStatisticsColumns holds the columns for table guild_week_statistics.
var guildWeekStatisticsColumns = GuildWeekStatisticsColumns{
	Id:                         "id",
	ServerId:                   "server_id",
	Week:                       "week",
	GuildId:                    "guild_id",
	TotalAnchorNum:             "total_anchor_num",
	TotalLiveTime:              "total_live_time",
	TotalLivePerson:            "total_live_person",
	NewAnchorNum:               "new_anchor_num",
	NewAnchorLiveTime:          "new_anchor_live_time",
	NewAnchorLivePerson:        "new_anchor_live_person",
	StarAnchorNum:              "star_anchor_num",
	StarAnchorLiveTime:         "star_anchor_live_time",
	StarAnchorLivePerson:       "star_anchor_live_person",
	NewSettlementAnchorNum:     "new_settlement_anchor_num",
	SettlementAnchorNum:        "settlement_anchor_num",
	NewStarAnchorNum:           "new_star_anchor_num",
	NewStarAnchorLiveTime:      "new_star_anchor_live_time",
	NewStarAnchorLivePerson:    "new_star_anchor_live_person",
	SettlementAnchorLiveTime:   "settlement_anchor_live_time",
	SettlementIncome:           "settlement_income",
	NewLivedUnsettlementPerson: "new_lived_unsettlement_person",
	LivedUnsettlementPerson:    "lived_unsettlement_person",
	LivedUnsettlementLiveTime:  "lived_unsettlement_live_time",
	WeekAwardDiamonds:          "week_award_diamonds",
	Remarks:                    "remarks",
	InvitePersonNum:            "invite_person_num",
	TotalFee:                   "total_fee",
	InviteNormalUserNum:        "invite_normal_user_num",
	InviteNormalUserTotalFee:   "invite_normal_user_total_fee",
	InviteAnchorNum:            "invite_anchor_num",
	InviteAnchorTotalFee:       "invite_anchor_total_fee",
	InviteMasterNum:            "invite_master_num",
	InviteMasterTotalFee:       "invite_master_total_fee",
	EffectiveUserNum:           "effective_user_num",
	EffectiveUserTotalFee:      "effective_user_total_fee",
	CreateTime:                 "create_time",
	UpdateTime:                 "update_time",
}

// NewGuildWeekStatisticsDao creates and returns a new DAO object for table data access.
func NewGuildWeekStatisticsDao() *GuildWeekStatisticsDao {
	return &GuildWeekStatisticsDao{
		group:   "default",
		table:   "guild_week_statistics",
		columns: guildWeekStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GuildWeekStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GuildWeekStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GuildWeekStatisticsDao) Columns() GuildWeekStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GuildWeekStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GuildWeekStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GuildWeekStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

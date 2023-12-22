// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CharmAnchorDayStatisticsDao is the data access object for table charm_anchor_day_statistics.
type CharmAnchorDayStatisticsDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns CharmAnchorDayStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// CharmAnchorDayStatisticsColumns defines and stores column names for table charm_anchor_day_statistics.
type CharmAnchorDayStatisticsColumns struct {
	Id                   string // 主键id
	StatisticsTime       string // 统计时间
	UserId               string // 用户id
	GuildId              string // 工会id
	LiveTime             string // 直播时长（s）
	EffectLiveTime       string // 有效直播时长（s） 大于30分钟的直播才算有效直播累计
	IsEffectDay          string // 是否有效天 0非 1是
	NewFansCount         string // 新增粉丝数
	NewSendGiftPerson    string // 新增送礼人数
	AllSendGiftPerson    string // 总送礼人数
	AllActivePersonCount string // 直播间活跃人数，发送过消息
	JoinRoomPersonCount  string // 进入房间总人数
	Remark               string // 备注
	CreateTime           string //
	UpdateTime           string //
}

// charmAnchorDayStatisticsColumns holds the columns for table charm_anchor_day_statistics.
var charmAnchorDayStatisticsColumns = CharmAnchorDayStatisticsColumns{
	Id:                   "id",
	StatisticsTime:       "statistics_time",
	UserId:               "user_id",
	GuildId:              "guild_id",
	LiveTime:             "live_time",
	EffectLiveTime:       "effect_live_time",
	IsEffectDay:          "is_effect_day",
	NewFansCount:         "new_fans_count",
	NewSendGiftPerson:    "new_send_gift_person",
	AllSendGiftPerson:    "all_send_gift_person",
	AllActivePersonCount: "all_active_person_count",
	JoinRoomPersonCount:  "join_room_person_count",
	Remark:               "remark",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
}

// NewCharmAnchorDayStatisticsDao creates and returns a new DAO object for table data access.
func NewCharmAnchorDayStatisticsDao() *CharmAnchorDayStatisticsDao {
	return &CharmAnchorDayStatisticsDao{
		group:   "default",
		table:   "charm_anchor_day_statistics",
		columns: charmAnchorDayStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CharmAnchorDayStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CharmAnchorDayStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CharmAnchorDayStatisticsDao) Columns() CharmAnchorDayStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CharmAnchorDayStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CharmAnchorDayStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CharmAnchorDayStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

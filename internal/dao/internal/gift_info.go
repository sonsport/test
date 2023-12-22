// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GiftInfoDao is the data access object for table gift_info.
type GiftInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns GiftInfoColumns // columns contains all the column names of Table for convenient usage.
}

// GiftInfoColumns defines and stores column names for table gift_info.
type GiftInfoColumns struct {
	Gid              string // 礼物Id 自增
	Name             string // 礼物名称
	Diamonds         string // 价格-钻石
	OriginalDiamonds string // 原价
	Type             string // 礼物大类 0普通礼物  1幸运礼物
	GiftType         string // 0普通小礼物（无特效）1特效礼物  2 房间红包雨 3 全站红包雨 4 全站通知礼物
	Icon             string // 礼物图标
	AnimEffectUrl    string // 特效礼物url
	FileMd5          string // 礼物md5值
	State            string // 0下架  1上架
	Sort             string // 排序 顺序
	ResVoiceUrl      string // 声音
	ResType          string // 1 icon/2 gif/3 文件资源；4为3d；5为面部识别；6为svga；
	IsBatter         string // 是否连送
	MinLevel         string // 最小等级
	LuckPoolType     string // 奖池类型 0 90% 1 95%
	Badge            string // 礼物徽章
	StartTime        string //
	EndTime          string //
	CreatedTime      string //
	UpdatedTime      string //
}

// giftInfoColumns holds the columns for table gift_info.
var giftInfoColumns = GiftInfoColumns{
	Gid:              "gid",
	Name:             "name",
	Diamonds:         "diamonds",
	OriginalDiamonds: "original_diamonds",
	Type:             "type",
	GiftType:         "gift_type",
	Icon:             "icon",
	AnimEffectUrl:    "anim_effect_url",
	FileMd5:          "file_md5",
	State:            "state",
	Sort:             "sort",
	ResVoiceUrl:      "res_voice_url",
	ResType:          "res_type",
	IsBatter:         "is_batter",
	MinLevel:         "min_level",
	LuckPoolType:     "luck_pool_type",
	Badge:            "badge",
	StartTime:        "start_time",
	EndTime:          "end_time",
	CreatedTime:      "created_time",
	UpdatedTime:      "updated_time",
}

// NewGiftInfoDao creates and returns a new DAO object for table data access.
func NewGiftInfoDao() *GiftInfoDao {
	return &GiftInfoDao{
		group:   "default",
		table:   "gift_info",
		columns: giftInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GiftInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GiftInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GiftInfoDao) Columns() GiftInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GiftInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GiftInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GiftInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

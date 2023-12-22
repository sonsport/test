// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameConfigDao is the data access object for table game_config.
type GameConfigDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns GameConfigColumns // columns contains all the column names of Table for convenient usage.
}

// GameConfigColumns defines and stores column names for table game_config.
type GameConfigColumns struct {
	Id                 string //
	GameId             string // 游戏ID
	GameKey            string // 第三方业务唯一标识
	GameName           string // 游戏名称
	IconUrl            string // 图片地址
	MiniIconUrl        string // 小图标
	GameUrl            string // 游戏地址
	GameType           string // 游戏类型 0自研 1joyplay 2cocos
	IsMini             string // 是否半屏 0全屏 1半屏
	IsHot              string // 是否热门 0不是 1是热门
	SystemRate         string // 抽水比例  *10000的整数 最小为万分之一
	AnchorRate         string // 主播分成  *10000的整数 最小为万分之一
	Desc               string // 备注
	GameVersion        string // 游戏版本
	Md5Version         string // MD5版本
	CocosEngineVersion string // cocos引擎版本
	ShowSystem         string // 显示系统 0 全部 1 android 2 ios
	WinRatio           string // 游戏宽高比
	Entrypoint         string // Ws接入点
	Status             string // 状态 1 正常 2 下架
	Sort               string // 排序
	Orientation        string // 游戏方向 横竖屏
	IsFullWin          string // 是否全屏 1全屏
	CreateTime         string //
	UpdateTime         string //
}

// gameConfigColumns holds the columns for table game_config.
var gameConfigColumns = GameConfigColumns{
	Id:                 "id",
	GameId:             "game_id",
	GameKey:            "game_key",
	GameName:           "game_name",
	IconUrl:            "icon_url",
	MiniIconUrl:        "mini_icon_url",
	GameUrl:            "game_url",
	GameType:           "game_type",
	IsMini:             "is_mini",
	IsHot:              "is_hot",
	SystemRate:         "system_rate",
	AnchorRate:         "anchor_rate",
	Desc:               "desc",
	GameVersion:        "game_version",
	Md5Version:         "md5_version",
	CocosEngineVersion: "cocos_engine_version",
	ShowSystem:         "show_system",
	WinRatio:           "win_ratio",
	Entrypoint:         "entrypoint",
	Status:             "status",
	Sort:               "sort",
	Orientation:        "orientation",
	IsFullWin:          "is_full_win",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewGameConfigDao creates and returns a new DAO object for table data access.
func NewGameConfigDao() *GameConfigDao {
	return &GameConfigDao{
		group:   "default",
		table:   "game_config",
		columns: gameConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameConfigDao) Columns() GameConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

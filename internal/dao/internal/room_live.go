// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLiveDao is the data access object for table room_live.
type RoomLiveDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RoomLiveColumns // columns contains all the column names of Table for convenient usage.
}

// RoomLiveColumns defines and stores column names for table room_live.
type RoomLiveColumns struct {
	Id                       string // 直播间Id
	UserId                   string // 主播ID
	RoomName                 string // 直播间名称
	RoomId                   string // 第三方直播间ID全局唯一
	StreamId                 string // 推拉流ID全局唯一
	ImagePath                string // 图片路径
	Score                    string // 排序分值
	MemberCount              string // 累计观看人数
	SendGiftPersonCount      string // 送礼人数
	SendLuckyGiftPersonCount string // 送幸运礼人数
	Status                   string // -1 所有状态 0  预开播 1  开播中 2  直播正常结束 3  由巡管或运营管理员将直播间关闭 4  由服务器轮询出异常退出房间 5  预开播时检测到有未关播的房间，将其退出房间 6  其他异常退出房间 7  正常关闭预开播 8  系统检测关闭预开播 9  暂时离开
	CoinAmount               string // 打赏金币数量 放大100倍
	BeginTime                string // 直播开始时间
	EndTime                  string // 直播结束时间
	LastHeartbeatTime        string // 最后一次心跳时间
	FocusCount               string // 本次直播涨粉数量
	Role                     string // 主播类型 0 用户 1 主播
	BattleStatus             string // pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK
	PrivateStatus            string // 0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
	GameId                   string // 直播间配置的游戏id
	UnlockPrice              string // 钻石房解锁价格
	Password                 string // 密码房密码
	WishGiftId               string // 心愿礼物id
	WishGiftCount            string // 心愿礼物个数
	RoomType                 string // 房间类型 0 普通直播间  2 多人直播间
	MaxMultiple              string // 最大值窗口数
	LayoutSet                string // 布局设置  1 突出 2 平铺
	CreateTime               string //
	UpdateTime               string //
}

// roomLiveColumns holds the columns for table room_live.
var roomLiveColumns = RoomLiveColumns{
	Id:                       "id",
	UserId:                   "user_id",
	RoomName:                 "room_name",
	RoomId:                   "room_id",
	StreamId:                 "stream_id",
	ImagePath:                "image_path",
	Score:                    "score",
	MemberCount:              "member_count",
	SendGiftPersonCount:      "send_gift_person_count",
	SendLuckyGiftPersonCount: "send_lucky_gift_person_count",
	Status:                   "status",
	CoinAmount:               "coin_amount",
	BeginTime:                "begin_time",
	EndTime:                  "end_time",
	LastHeartbeatTime:        "last_heartbeat_time",
	FocusCount:               "focus_count",
	Role:                     "role",
	BattleStatus:             "battle_status",
	PrivateStatus:            "private_status",
	GameId:                   "game_id",
	UnlockPrice:              "unlock_price",
	Password:                 "password",
	WishGiftId:               "wish_gift_id",
	WishGiftCount:            "wish_gift_count",
	RoomType:                 "room_type",
	MaxMultiple:              "max_multiple",
	LayoutSet:                "layout_set",
	CreateTime:               "create_time",
	UpdateTime:               "update_time",
}

// NewRoomLiveDao creates and returns a new DAO object for table data access.
func NewRoomLiveDao() *RoomLiveDao {
	return &RoomLiveDao{
		group:   "default",
		table:   "room_live",
		columns: roomLiveColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomLiveDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomLiveDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomLiveDao) Columns() RoomLiveColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomLiveDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomLiveDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomLiveDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

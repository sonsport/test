// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLive is the golang structure of table room_live for DAO operations like Where/Data.
type RoomLive struct {
	g.Meta                   `orm:"table:room_live, do:true"`
	Id                       interface{} // 直播间Id
	UserId                   interface{} // 主播ID
	RoomName                 interface{} // 直播间名称
	RoomId                   interface{} // 第三方直播间ID全局唯一
	StreamId                 interface{} // 推拉流ID全局唯一
	ImagePath                interface{} // 图片路径
	Score                    interface{} // 排序分值
	MemberCount              interface{} // 累计观看人数
	SendGiftPersonCount      interface{} // 送礼人数
	SendLuckyGiftPersonCount interface{} // 送幸运礼人数
	Status                   interface{} // -1 所有状态 0  预开播 1  开播中 2  直播正常结束 3  由巡管或运营管理员将直播间关闭 4  由服务器轮询出异常退出房间 5  预开播时检测到有未关播的房间，将其退出房间 6  其他异常退出房间 7  正常关闭预开播 8  系统检测关闭预开播 9  暂时离开
	CoinAmount               interface{} // 打赏金币数量 放大100倍
	BeginTime                interface{} // 直播开始时间
	EndTime                  interface{} // 直播结束时间
	LastHeartbeatTime        interface{} // 最后一次心跳时间
	FocusCount               interface{} // 本次直播涨粉数量
	Role                     interface{} // 主播类型 0 用户 1 主播
	BattleStatus             interface{} // pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK
	PrivateStatus            interface{} // 0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
	GameId                   interface{} // 直播间配置的游戏id
	UnlockPrice              interface{} // 钻石房解锁价格
	Password                 interface{} // 密码房密码
	WishGiftId               interface{} // 心愿礼物id
	WishGiftCount            interface{} // 心愿礼物个数
	RoomType                 interface{} // 房间类型 0 普通直播间  2 多人直播间
	MaxMultiple              interface{} // 最大值窗口数
	LayoutSet                interface{} // 布局设置  1 突出 2 平铺
	CreateTime               interface{} //
	UpdateTime               interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomLive is the golang structure for table room_live.
type RoomLive struct {
	Id                       uint64 `json:"id"                       description:"直播间Id"`
	UserId                   int64  `json:"userId"                   description:"主播ID"`
	RoomName                 string `json:"roomName"                 description:"直播间名称"`
	RoomId                   string `json:"roomId"                   description:"第三方直播间ID全局唯一"`
	StreamId                 string `json:"streamId"                 description:"推拉流ID全局唯一"`
	ImagePath                string `json:"imagePath"                description:"图片路径"`
	Score                    int    `json:"score"                    description:"排序分值"`
	MemberCount              int64  `json:"memberCount"              description:"累计观看人数"`
	SendGiftPersonCount      int    `json:"sendGiftPersonCount"      description:"送礼人数"`
	SendLuckyGiftPersonCount int    `json:"sendLuckyGiftPersonCount" description:"送幸运礼人数"`
	Status                   int    `json:"status"                   description:"-1 所有状态 0  预开播 1  开播中 2  直播正常结束 3  由巡管或运营管理员将直播间关闭 4  由服务器轮询出异常退出房间 5  预开播时检测到有未关播的房间，将其退出房间 6  其他异常退出房间 7  正常关闭预开播 8  系统检测关闭预开播 9  暂时离开"`
	CoinAmount               int64  `json:"coinAmount"               description:"打赏金币数量 放大100倍"`
	BeginTime                int64  `json:"beginTime"                description:"直播开始时间"`
	EndTime                  int64  `json:"endTime"                  description:"直播结束时间"`
	LastHeartbeatTime        int64  `json:"lastHeartbeatTime"        description:"最后一次心跳时间"`
	FocusCount               int    `json:"focusCount"               description:"本次直播涨粉数量"`
	Role                     int    `json:"role"                     description:"主播类型 0 用户 1 主播"`
	BattleStatus             int    `json:"battleStatus"             description:"pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK"`
	PrivateStatus            int    `json:"privateStatus"            description:"0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间"`
	GameId                   int    `json:"gameId"                   description:"直播间配置的游戏id"`
	UnlockPrice              int    `json:"unlockPrice"              description:"钻石房解锁价格"`
	Password                 string `json:"password"                 description:"密码房密码"`
	WishGiftId               int    `json:"wishGiftId"               description:"心愿礼物id"`
	WishGiftCount            int    `json:"wishGiftCount"            description:"心愿礼物个数"`
	RoomType                 int    `json:"roomType"                 description:"房间类型 0 普通直播间  2 多人直播间"`
	MaxMultiple              int    `json:"maxMultiple"              description:"最大值窗口数"`
	LayoutSet                int    `json:"layoutSet"                description:"布局设置  1 突出 2 平铺"`
	CreateTime               int64  `json:"createTime"               description:""`
	UpdateTime               int64  `json:"updateTime"               description:""`
}

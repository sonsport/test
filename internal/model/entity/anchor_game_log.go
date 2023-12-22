// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AnchorGameLog is the golang structure for table anchor_game_log.
type AnchorGameLog struct {
	Id             uint64 `json:"id"             description:"ID"`
	AnchorId       int64  `json:"anchorId"       description:"主播ID"`
	UserId         int64  `json:"userId"         description:"用户ID"`
	RoomId         string `json:"roomId"         description:"第三方直播间ID全局唯一"`
	GameId         int    `json:"gameId"         description:"直播间配置的游戏id"`
	Diamonds       uint64 `json:"diamonds"       description:"钻石数"`
	AnchorRate     int    `json:"anchorRate"     description:"主播分成  *10000的整数 最小为万分之一"`
	BeforeDiamonds uint64 `json:"beforeDiamonds" description:"当前钻石数"`
	LogType        int    `json:"logType"        description:"流水类型 1 增加 2 扣除"`
	AfterDiamonds  uint64 `json:"afterDiamonds"  description:"增加之后的钻石数"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
}

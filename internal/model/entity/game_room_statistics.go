// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GameRoomStatistics is the golang structure for table game_room_statistics.
type GameRoomStatistics struct {
	Id           int    `json:"id"           description:""`
	BeginTime    int64  `json:"beginTime"    description:"开始时间"`
	EndTime      int64  `json:"endTime"      description:"结束时间"`
	TimeLevel    int    `json:"timeLevel"    description:"时间维度：1.每小时，2.每天"`
	RoomId       string `json:"roomId"       description:"直播间id"`
	Diamonds     int64  `json:"diamonds"     description:"总钻石"`
	GameId       int64  `json:"gameId"       description:"游戏id"`
	AnchorReward int64  `json:"anchorReward" description:"主播收益"`
	PlayNum      int64  `json:"playNum"      description:"操作数"`
	PlayPerson   int64  `json:"playPerson"   description:"游戏人数"`
	AnchorId     int64  `json:"anchorId"     description:"主播id"`
	MasterId     int64  `json:"masterId"     description:"公会长id"`
	ServerId     int64  `json:"serverId"     description:"运营号id"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

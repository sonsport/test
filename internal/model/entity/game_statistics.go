// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GameStatistics is the golang structure for table game_statistics.
type GameStatistics struct {
	Id            int    `json:"id"            description:""`
	AppChannel    string `json:"appChannel"    description:"app渠道"`
	GameKey       string `json:"gameKey"       description:"游戏key"`
	GameIncome    int64  `json:"gameIncome"    description:"游戏收益"`
	GameOutlay    int64  `json:"gameOutlay"    description:"游戏支出"`
	PlayUserCount int64  `json:"playUserCount" description:"游戏人数"`
	RoomCount     int64  `json:"roomCount"     description:"房间数"`
	PlayCount     int64  `json:"playCount"     description:"操作数"`
	BeginTime     int64  `json:"beginTime"     description:"开始时间"`
	EndTime       int64  `json:"endTime"       description:"结束时间"`
	TimeLevel     int    `json:"timeLevel"     description:"时间维度：1.每小时，2.每天"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

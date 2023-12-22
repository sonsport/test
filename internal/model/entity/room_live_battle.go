// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomLiveBattle is the golang structure for table room_live_battle.
type RoomLiveBattle struct {
	Id                  uint64 `json:"id"                  description:"直播间Id"`
	UserId              int64  `json:"userId"              description:"主播ID"`
	BattleType          int    `json:"battleType"          description:"pk类型 0 随机 1 指定 2后台拉起"`
	StreamId            string `json:"streamId"            description:"推拉流ID全局唯一"`
	RoomId              string `json:"roomId"              description:"房间ID"`
	DiamondsCount       int64  `json:"diamondsCount"       description:"打赏钻石数量"`
	ScoreCount          int64  `json:"scoreCount"          description:"打赏积分数量"`
	TargetUserId        int64  `json:"targetUserId"        description:"对手主播ID"`
	TargetStreamId      string `json:"targetStreamId"      description:"对手stream_id"`
	TargetDiamondsCount int64  `json:"targetDiamondsCount" description:"对手打赏钻石数量"`
	TargetScoreCount    int64  `json:"targetScoreCount"    description:"对手打赏积分数量"`
	TargetRoomId        string `json:"targetRoomId"        description:"房间ID"`
	BattleStatus        int    `json:"battleStatus"        description:"0 匹配中 1 pk中 2 邀请人pk取消 3 对方拒绝 4 惩罚中 5 pk正常结束 6 pk异常结束"`
	EndType             int    `json:"endType"             description:"结束原因 0 正常结束 1 超时结束 2 手动结束 3 系统结束"`
	BeginTime           int64  `json:"beginTime"           description:"pk开始时间"`
	EndTime             int64  `json:"endTime"             description:"pk结束时间"`
	UserAnswerTime      int64  `json:"userAnswerTime"      description:"用户响应时间"`
	TargetAnswerTime    int64  `json:"targetAnswerTime"    description:"对方响应时间"`
	CreateTime          int64  `json:"createTime"          description:""`
	UpdateTime          int64  `json:"updateTime"          description:""`
}

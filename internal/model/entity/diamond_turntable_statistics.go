// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DiamondTurntableStatistics is the golang structure for table diamond_turntable_statistics.
type DiamondTurntableStatistics struct {
	Id             uint64 `json:"id"             description:""`
	UserId         uint64 `json:"userId"         description:"用户id"`
	WatchTime      uint   `json:"watchTime"      description:"在直播间的有效观看时间 单位s"`
	RaffleNum      uint   `json:"raffleNum"      description:"获得的抽奖次数"`
	UseNum         uint   `json:"useNum"         description:"使用的抽奖次数"`
	Gid            uint   `json:"gid"            description:"商品表id"`
	Num            uint   `json:"num"            description:"数量"`
	TurntableCount uint64 `json:"turntableCount" description:"奖池期数"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
	Cid            uint   `json:"cid"            description:"抽奖配置id"`
	PrizeName      string `json:"prizeName"      description:"奖品名称"`
}

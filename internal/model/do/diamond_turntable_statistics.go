// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableStatistics is the golang structure of table diamond_turntable_statistics for DAO operations like Where/Data.
type DiamondTurntableStatistics struct {
	g.Meta         `orm:"table:diamond_turntable_statistics, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户id
	WatchTime      interface{} // 在直播间的有效观看时间 单位s
	RaffleNum      interface{} // 获得的抽奖次数
	UseNum         interface{} // 使用的抽奖次数
	Gid            interface{} // 商品表id
	Num            interface{} // 数量
	TurntableCount interface{} // 奖池期数
	CreateTime     interface{} //
	UpdateTime     interface{} //
	Cid            interface{} // 抽奖配置id
	PrizeName      interface{} // 奖品名称
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DiamondTurntableConfig is the golang structure for table diamond_turntable_config.
type DiamondTurntableConfig struct {
	Id         uint64 `json:"id"         description:""`
	Gid        uint   `json:"gid"        description:"商品表id"`
	Gtype      int    `json:"gtype"      description:"商品类型 1装扮 2钻石"`
	Num        uint   `json:"num"        description:"数量"`
	Rate       int    `json:"rate"       description:"概率百分比  *100的整数"`
	Status     int    `json:"status"     description:"商品状态 0失效 1生效"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

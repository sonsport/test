// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DiamondTurntableRecord is the golang structure for table diamond_turntable_record.
type DiamondTurntableRecord struct {
	Id             uint64 `json:"id"             description:""`
	UserId         uint64 `json:"userId"         description:"用户id"`
	Cid            uint   `json:"cid"            description:"配置表id"`
	Gid            uint   `json:"gid"            description:"商品表id"`
	Num            uint   `json:"num"            description:"数量"`
	TurntableCount uint64 `json:"turntableCount" description:"奖池期数"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
}

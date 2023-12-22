// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DiamondTurntableConfig is the golang structure of table diamond_turntable_config for DAO operations like Where/Data.
type DiamondTurntableConfig struct {
	g.Meta     `orm:"table:diamond_turntable_config, do:true"`
	Id         interface{} //
	Gid        interface{} // 商品表id
	Gtype      interface{} // 商品类型 1装扮 2钻石
	Num        interface{} // 数量
	Rate       interface{} // 概率百分比  *100的整数
	Status     interface{} // 商品状态 0失效 1生效
	CreateTime interface{} //
	UpdateTime interface{} //
}

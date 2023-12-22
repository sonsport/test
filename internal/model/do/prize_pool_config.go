// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PrizePoolConfig is the golang structure of table prize_pool_config for DAO operations like Where/Data.
type PrizePoolConfig struct {
	g.Meta          `orm:"table:prize_pool_config, do:true"`
	Id              interface{} // 主键id
	Gid             interface{} // 幸运礼物id
	StartTime       *gtime.Time // 配置生效时间
	Type            interface{} // 配置类型 1:默认 2:其他
	CashbackRatioId interface{} // 奖池返现配置id
	State           interface{} // 状态 1:生效 2:失效
	UseCount        interface{} // 使用次数
	CreateTime      interface{} // 创建时间
	UpdateTime      interface{} // 更新时间
}

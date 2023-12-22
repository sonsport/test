// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// StarLevelResetConfig is the golang structure of table star_level_reset_config for DAO operations like Where/Data.
type StarLevelResetConfig struct {
	g.Meta            `orm:"table:star_level_reset_config, do:true"`
	Id                interface{} // 自增id
	StarLevel         interface{} // 主播星级
	RuleType          interface{} // 类别 0保级 1升级 2降级
	IncomeMin         interface{} // 金币min 不用扩大100
	SendGiftPersonMin interface{} // 送礼人数min
	SelfSendRateMax   interface{} // 自己送自己金币占比max 0.4 存储40 扩大100倍
	CreateTime        interface{} // 创建时间
	UpdateTime        interface{} // 更新时间
}

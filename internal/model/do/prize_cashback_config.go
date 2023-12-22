// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PrizeCashbackConfig is the golang structure of table prize_cashback_config for DAO operations like Where/Data.
type PrizeCashbackConfig struct {
	g.Meta                 `orm:"table:prize_cashback_config, do:true"`
	Id                     interface{} // 主键id
	Name                   interface{} // 名称
	CashbackRatio          interface{} // 返现比例
	FiveTimesNumber        interface{} // 5倍数量
	TenTimesNumber         interface{} // 10倍数量
	FiftyTimesNumber       interface{} // 50倍数量
	OneHundredTimesNumber  interface{} // 100倍数量
	FiveHundredTimesNumber interface{} // 500倍
	CreateTime             interface{} // 创建时间
	UpdateTime             interface{} // 更新时间
}

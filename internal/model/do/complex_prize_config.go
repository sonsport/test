// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ComplexPrizeConfig is the golang structure of table complex_prize_config for DAO operations like Where/Data.
type ComplexPrizeConfig struct {
	g.Meta       `orm:"table:complex_prize_config, do:true"`
	Id           interface{} //
	ActivityType interface{} // 抽奖活动类型 1 圣诞 2 元旦
	Icon         interface{} // 图标
	DrawType     interface{} // 抽奖类型 1装扮 2钻石 3优惠券 4 实物
	LinkId       interface{} // 关联商品id  1,2  逗号分割
	PrizeName    interface{} // 奖品名称
	Num          interface{} // 数量
	IsFirstPool  interface{} // 是否第一个奖池出现
	Rate         interface{} // 概率万分比  *10000的整数
	Status       interface{} // 配置状态 0失效 1生效
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ComplexPrizeRecord is the golang structure of table complex_prize_record for DAO operations like Where/Data.
type ComplexPrizeRecord struct {
	g.Meta       `orm:"table:complex_prize_record, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户id
	CId          interface{} // 抽奖配置id
	ActivityType interface{} // 抽奖活动类型 1 圣诞 2 元旦
	Icon         interface{} // 图标
	DrawType     interface{} // 抽奖类型 1装扮 2钻石 3优惠券 4 实物
	PropsLogId   interface{} // 扣除道具流水id
	LinkId       interface{} // 关联商品id  1,2  逗号分割
	PrizeName    interface{} // 奖品名称
	PrizeCount   interface{} // 奖池期数
	CreateTime   interface{} // 创建时间
}

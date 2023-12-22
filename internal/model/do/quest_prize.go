// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// QuestPrize is the golang structure of table quest_prize for DAO operations like Where/Data.
type QuestPrize struct {
	g.Meta     `orm:"table:quest_prize, do:true"`
	Id         interface{} // id
	QuestId    interface{} // 任务id
	PrizeName  interface{} // 奖品名称
	PrizeIcon  interface{} // 奖品图标
	PrizeType  interface{} // 奖品类型 1装扮 2钻石 3优惠券 4贝壳 5 活跃值 6 用户经验 7 主播经验 8 主播收益
	LinkId     interface{} // 关联商品id  1,2  逗号分割
	PrizeNum   interface{} // 奖品数量
	Status     interface{} // 配置状态 0失效 1生效
	CreateTime interface{} // createdAt
	UpdateTime interface{} // updatedAt
}

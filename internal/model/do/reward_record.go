// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RewardRecord is the golang structure of table reward_record for DAO operations like Where/Data.
type RewardRecord struct {
	g.Meta     `orm:"table:reward_record, do:true"`
	Id         interface{} // 自增Id
	UserId     interface{} // 用户ID
	GuildId    interface{} // 公会id
	RewardType interface{} // 奖励类型，  比如type  1：为3周2次结算  2……  后续扩展
	Reward     interface{} // 活动奖励给公会 印尼盾 分   该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	Remark     interface{} // 备注
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

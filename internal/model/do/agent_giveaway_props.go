// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AgentGiveawayProps is the golang structure of table agent_giveaway_props for DAO operations like Where/Data.
type AgentGiveawayProps struct {
	g.Meta          `orm:"table:agent_giveaway_props, do:true"`
	Id              interface{} // 直播间Id
	UserId          interface{} // 主播ID
	PropsName       interface{} // 道具名称
	PropsIcon       interface{} // 道具icon
	PropsDay        interface{} // 道具天数 30
	PropsStates     interface{} // 道具状态 1 待使用 2 已使用
	PropsType       interface{} // 道具类型 1 礼物 2 vip 3 装扮
	PropsLink       interface{} // 道具关联id
	BalanceRecordId interface{} // 流水记录id
	UseTime         interface{} // 使用时间
	TargetUserId    interface{} // 使用目标id
	ExpirationTime  interface{} // 过期时间 -1 不过期
	CreateTime      interface{} //
	UpdateTime      interface{} //
}

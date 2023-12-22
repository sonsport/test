// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AgentGiveawayProps is the golang structure for table agent_giveaway_props.
type AgentGiveawayProps struct {
	Id              uint64 `json:"id"              description:"直播间Id"`
	UserId          int64  `json:"userId"          description:"主播ID"`
	PropsName       string `json:"propsName"       description:"道具名称"`
	PropsIcon       string `json:"propsIcon"       description:"道具icon"`
	PropsDay        int    `json:"propsDay"        description:"道具天数 30"`
	PropsStates     int    `json:"propsStates"     description:"道具状态 1 待使用 2 已使用"`
	PropsType       int    `json:"propsType"       description:"道具类型 1 礼物 2 vip 3 装扮"`
	PropsLink       int    `json:"propsLink"       description:"道具关联id"`
	BalanceRecordId int64  `json:"balanceRecordId" description:"流水记录id"`
	UseTime         int64  `json:"useTime"         description:"使用时间"`
	TargetUserId    int    `json:"targetUserId"    description:"使用目标id"`
	ExpirationTime  int64  `json:"expirationTime"  description:"过期时间 -1 不过期"`
	CreateTime      int64  `json:"createTime"      description:""`
	UpdateTime      int64  `json:"updateTime"      description:""`
}

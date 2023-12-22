// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RewardRecord is the golang structure for table reward_record.
type RewardRecord struct {
	Id         uint64 `json:"id"         description:"自增Id"`
	UserId     int64  `json:"userId"     description:"用户ID"`
	GuildId    int64  `json:"guildId"    description:"公会id"`
	RewardType int    `json:"rewardType" description:"奖励类型，  比如type  1：为3周2次结算  2……  后续扩展"`
	Reward     int64  `json:"reward"     description:"活动奖励给公会 印尼盾 分   该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会"`
	Remark     string `json:"remark"     description:"备注"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

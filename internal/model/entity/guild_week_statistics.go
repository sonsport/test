// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildWeekStatistics is the golang structure for table guild_week_statistics.
type GuildWeekStatistics struct {
	Id                         uint   `json:"id"                         description:"自增id"`
	ServerId                   int64  `json:"serverId"                   description:"运营号id"`
	Week                       int    `json:"week"                       description:"周，一年中的第几周 例如:202230 2022年的第30周"`
	GuildId                    int    `json:"guildId"                    description:"公会Id 关联guild_info自增"`
	TotalAnchorNum             int    `json:"totalAnchorNum"             description:"总主播数"`
	TotalLiveTime              int64  `json:"totalLiveTime"              description:"总直播时长(秒)"`
	TotalLivePerson            int    `json:"totalLivePerson"            description:"总开播人数"`
	NewAnchorNum               int    `json:"newAnchorNum"               description:"新增主播人数"`
	NewAnchorLiveTime          int64  `json:"newAnchorLiveTime"          description:"新增主播直播时长"`
	NewAnchorLivePerson        int    `json:"newAnchorLivePerson"        description:"新增主播开播人数"`
	StarAnchorNum              int    `json:"starAnchorNum"              description:"总标星主播数"`
	StarAnchorLiveTime         int64  `json:"starAnchorLiveTime"         description:"总标星主播总开播时长"`
	StarAnchorLivePerson       int    `json:"starAnchorLivePerson"       description:"总标星主播开播人数"`
	NewSettlementAnchorNum     int    `json:"newSettlementAnchorNum"     description:"新增达标结算主播数"`
	SettlementAnchorNum        int    `json:"settlementAnchorNum"        description:"达标结算主播数"`
	NewStarAnchorNum           int    `json:"newStarAnchorNum"           description:"新增标星主播人数"`
	NewStarAnchorLiveTime      int64  `json:"newStarAnchorLiveTime"      description:"新增标星主播总开播时长"`
	NewStarAnchorLivePerson    int    `json:"newStarAnchorLivePerson"    description:"新增标星主播开播人数"`
	SettlementAnchorLiveTime   int64  `json:"settlementAnchorLiveTime"   description:"达标结算主播总直播时长"`
	SettlementIncome           int64  `json:"settlementIncome"           description:"达标结算收益"`
	NewLivedUnsettlementPerson int    `json:"newLivedUnsettlementPerson" description:"新增开播未达标结算主播数"`
	LivedUnsettlementPerson    int    `json:"livedUnsettlementPerson"    description:"已开播未结算主播数"`
	LivedUnsettlementLiveTime  int64  `json:"livedUnsettlementLiveTime"  description:"已开播未结算主播总开播时长"`
	WeekAwardDiamonds          int    `json:"weekAwardDiamonds"          description:"本周奖励钻石数"`
	Remarks                    string `json:"remarks"                    description:"奖励钻石的备注"`
	InvitePersonNum            int    `json:"invitePersonNum"            description:"邀新总人数"`
	TotalFee                   int64  `json:"totalFee"                   description:"拉新总冲值"`
	InviteNormalUserNum        int    `json:"inviteNormalUserNum"        description:"邀新普通用户数"`
	InviteNormalUserTotalFee   int64  `json:"inviteNormalUserTotalFee"   description:"邀新普通用户充值"`
	InviteAnchorNum            int    `json:"inviteAnchorNum"            description:"邀新主播数"`
	InviteAnchorTotalFee       int64  `json:"inviteAnchorTotalFee"       description:"邀新主播充值"`
	InviteMasterNum            int    `json:"inviteMasterNum"            description:"邀新公会数"`
	InviteMasterTotalFee       int64  `json:"inviteMasterTotalFee"       description:"邀新公会长充值"`
	EffectiveUserNum           int    `json:"effectiveUserNum"           description:"有效用户数"`
	EffectiveUserTotalFee      int64  `json:"effectiveUserTotalFee"      description:"有效用户金币"`
	CreateTime                 int64  `json:"createTime"                 description:"创建时间"`
	UpdateTime                 int64  `json:"updateTime"                 description:"更新时间"`
}

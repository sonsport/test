// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GuildWeekStatistics is the golang structure of table guild_week_statistics for DAO operations like Where/Data.
type GuildWeekStatistics struct {
	g.Meta                     `orm:"table:guild_week_statistics, do:true"`
	Id                         interface{} // 自增id
	ServerId                   interface{} // 运营号id
	Week                       interface{} // 周，一年中的第几周 例如:202230 2022年的第30周
	GuildId                    interface{} // 公会Id 关联guild_info自增
	TotalAnchorNum             interface{} // 总主播数
	TotalLiveTime              interface{} // 总直播时长(秒)
	TotalLivePerson            interface{} // 总开播人数
	NewAnchorNum               interface{} // 新增主播人数
	NewAnchorLiveTime          interface{} // 新增主播直播时长
	NewAnchorLivePerson        interface{} // 新增主播开播人数
	StarAnchorNum              interface{} // 总标星主播数
	StarAnchorLiveTime         interface{} // 总标星主播总开播时长
	StarAnchorLivePerson       interface{} // 总标星主播开播人数
	NewSettlementAnchorNum     interface{} // 新增达标结算主播数
	SettlementAnchorNum        interface{} // 达标结算主播数
	NewStarAnchorNum           interface{} // 新增标星主播人数
	NewStarAnchorLiveTime      interface{} // 新增标星主播总开播时长
	NewStarAnchorLivePerson    interface{} // 新增标星主播开播人数
	SettlementAnchorLiveTime   interface{} // 达标结算主播总直播时长
	SettlementIncome           interface{} // 达标结算收益
	NewLivedUnsettlementPerson interface{} // 新增开播未达标结算主播数
	LivedUnsettlementPerson    interface{} // 已开播未结算主播数
	LivedUnsettlementLiveTime  interface{} // 已开播未结算主播总开播时长
	WeekAwardDiamonds          interface{} // 本周奖励钻石数
	Remarks                    interface{} // 奖励钻石的备注
	InvitePersonNum            interface{} // 邀新总人数
	TotalFee                   interface{} // 拉新总冲值
	InviteNormalUserNum        interface{} // 邀新普通用户数
	InviteNormalUserTotalFee   interface{} // 邀新普通用户充值
	InviteAnchorNum            interface{} // 邀新主播数
	InviteAnchorTotalFee       interface{} // 邀新主播充值
	InviteMasterNum            interface{} // 邀新公会数
	InviteMasterTotalFee       interface{} // 邀新公会长充值
	EffectiveUserNum           interface{} // 有效用户数
	EffectiveUserTotalFee      interface{} // 有效用户金币
	CreateTime                 interface{} // 创建时间
	UpdateTime                 interface{} // 更新时间
}

// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalGuildSettlementRewardDao is internal type for wrapping internal DAO implements.
type internalGuildSettlementRewardDao = *internal.GuildSettlementRewardDao

// guildSettlementRewardDao is the data access object for table guild_settlement_reward.
// You can define custom methods on it to extend its functionality as you wish.
type guildSettlementRewardDao struct {
	internalGuildSettlementRewardDao
}

var (
	// GuildSettlementReward is globally public accessible object for table guild_settlement_reward operations.
	GuildSettlementReward = guildSettlementRewardDao{
		internal.NewGuildSettlementRewardDao(),
	}
)

// Fill with you ideas below.
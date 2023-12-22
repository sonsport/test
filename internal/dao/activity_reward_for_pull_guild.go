// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalActivityRewardForPullGuildDao is internal type for wrapping internal DAO implements.
type internalActivityRewardForPullGuildDao = *internal.ActivityRewardForPullGuildDao

// activityRewardForPullGuildDao is the data access object for table activity_reward_for_pull_guild.
// You can define custom methods on it to extend its functionality as you wish.
type activityRewardForPullGuildDao struct {
	internalActivityRewardForPullGuildDao
}

var (
	// ActivityRewardForPullGuild is globally public accessible object for table activity_reward_for_pull_guild operations.
	ActivityRewardForPullGuild = activityRewardForPullGuildDao{
		internal.NewActivityRewardForPullGuildDao(),
	}
)

// Fill with you ideas below.
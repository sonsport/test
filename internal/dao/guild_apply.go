// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalGuildApplyDao is internal type for wrapping internal DAO implements.
type internalGuildApplyDao = *internal.GuildApplyDao

// guildApplyDao is the data access object for table guild_apply.
// You can define custom methods on it to extend its functionality as you wish.
type guildApplyDao struct {
	internalGuildApplyDao
}

var (
	// GuildApply is globally public accessible object for table guild_apply operations.
	GuildApply = guildApplyDao{
		internal.NewGuildApplyDao(),
	}
)

// Fill with you ideas below.
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalGuildAnchorDao is internal type for wrapping internal DAO implements.
type internalGuildAnchorDao = *internal.GuildAnchorDao

// guildAnchorDao is the data access object for table guild_anchor.
// You can define custom methods on it to extend its functionality as you wish.
type guildAnchorDao struct {
	internalGuildAnchorDao
}

var (
	// GuildAnchor is globally public accessible object for table guild_anchor operations.
	GuildAnchor = guildAnchorDao{
		internal.NewGuildAnchorDao(),
	}
)

// Fill with you ideas below.

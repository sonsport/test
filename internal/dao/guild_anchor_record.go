// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalGuildAnchorRecordDao is internal type for wrapping internal DAO implements.
type internalGuildAnchorRecordDao = *internal.GuildAnchorRecordDao

// guildAnchorRecordDao is the data access object for table guild_anchor_record.
// You can define custom methods on it to extend its functionality as you wish.
type guildAnchorRecordDao struct {
	internalGuildAnchorRecordDao
}

var (
	// GuildAnchorRecord is globally public accessible object for table guild_anchor_record operations.
	GuildAnchorRecord = guildAnchorRecordDao{
		internal.NewGuildAnchorRecordDao(),
	}
)

// Fill with you ideas below.
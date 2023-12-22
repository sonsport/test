// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalActivityChannelConfigDao is internal type for wrapping internal DAO implements.
type internalActivityChannelConfigDao = *internal.ActivityChannelConfigDao

// activityChannelConfigDao is the data access object for table activity_channel_config.
// You can define custom methods on it to extend its functionality as you wish.
type activityChannelConfigDao struct {
	internalActivityChannelConfigDao
}

var (
	// ActivityChannelConfig is globally public accessible object for table activity_channel_config operations.
	ActivityChannelConfig = activityChannelConfigDao{
		internal.NewActivityChannelConfigDao(),
	}
)

// Fill with you ideas below.
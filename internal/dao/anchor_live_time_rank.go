// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalAnchorLiveTimeRankDao is internal type for wrapping internal DAO implements.
type internalAnchorLiveTimeRankDao = *internal.AnchorLiveTimeRankDao

// anchorLiveTimeRankDao is the data access object for table anchor_live_time_rank.
// You can define custom methods on it to extend its functionality as you wish.
type anchorLiveTimeRankDao struct {
	internalAnchorLiveTimeRankDao
}

var (
	// AnchorLiveTimeRank is globally public accessible object for table anchor_live_time_rank operations.
	AnchorLiveTimeRank = anchorLiveTimeRankDao{
		internal.NewAnchorLiveTimeRankDao(),
	}
)

// Fill with you ideas below.

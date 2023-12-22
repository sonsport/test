// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalActivePersonStatisticsDao is internal type for wrapping internal DAO implements.
type internalActivePersonStatisticsDao = *internal.ActivePersonStatisticsDao

// activePersonStatisticsDao is the data access object for table active_person_statistics.
// You can define custom methods on it to extend its functionality as you wish.
type activePersonStatisticsDao struct {
	internalActivePersonStatisticsDao
}

var (
	// ActivePersonStatistics is globally public accessible object for table active_person_statistics operations.
	ActivePersonStatistics = activePersonStatisticsDao{
		internal.NewActivePersonStatisticsDao(),
	}
)

// Fill with you ideas below.

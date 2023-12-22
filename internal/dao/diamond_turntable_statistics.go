// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalDiamondTurntableStatisticsDao is internal type for wrapping internal DAO implements.
type internalDiamondTurntableStatisticsDao = *internal.DiamondTurntableStatisticsDao

// diamondTurntableStatisticsDao is the data access object for table diamond_turntable_statistics.
// You can define custom methods on it to extend its functionality as you wish.
type diamondTurntableStatisticsDao struct {
	internalDiamondTurntableStatisticsDao
}

var (
	// DiamondTurntableStatistics is globally public accessible object for table diamond_turntable_statistics operations.
	DiamondTurntableStatistics = diamondTurntableStatisticsDao{
		internal.NewDiamondTurntableStatisticsDao(),
	}
)

// Fill with you ideas below.
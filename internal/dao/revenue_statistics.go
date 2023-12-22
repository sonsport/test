// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalRevenueStatisticsDao is internal type for wrapping internal DAO implements.
type internalRevenueStatisticsDao = *internal.RevenueStatisticsDao

// revenueStatisticsDao is the data access object for table revenue_statistics.
// You can define custom methods on it to extend its functionality as you wish.
type revenueStatisticsDao struct {
	internalRevenueStatisticsDao
}

var (
	// RevenueStatistics is globally public accessible object for table revenue_statistics operations.
	RevenueStatistics = revenueStatisticsDao{
		internal.NewRevenueStatisticsDao(),
	}
)

// Fill with you ideas below.

// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalBalanceStatisticsDao is internal type for wrapping internal DAO implements.
type internalBalanceStatisticsDao = *internal.BalanceStatisticsDao

// balanceStatisticsDao is the data access object for table balance_statistics.
// You can define custom methods on it to extend its functionality as you wish.
type balanceStatisticsDao struct {
	internalBalanceStatisticsDao
}

var (
	// BalanceStatistics is globally public accessible object for table balance_statistics operations.
	BalanceStatistics = balanceStatisticsDao{
		internal.NewBalanceStatisticsDao(),
	}
)

// Fill with you ideas below.

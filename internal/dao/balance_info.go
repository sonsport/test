// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalBalanceInfoDao is internal type for wrapping internal DAO implements.
type internalBalanceInfoDao = *internal.BalanceInfoDao

// balanceInfoDao is the data access object for table balance_info.
// You can define custom methods on it to extend its functionality as you wish.
type balanceInfoDao struct {
	internalBalanceInfoDao
}

var (
	// BalanceInfo is globally public accessible object for table balance_info operations.
	BalanceInfo = balanceInfoDao{
		internal.NewBalanceInfoDao(),
	}
)

// Fill with you ideas below.
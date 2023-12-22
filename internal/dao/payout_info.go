// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalPayoutInfoDao is internal type for wrapping internal DAO implements.
type internalPayoutInfoDao = *internal.PayoutInfoDao

// payoutInfoDao is the data access object for table payout_info.
// You can define custom methods on it to extend its functionality as you wish.
type payoutInfoDao struct {
	internalPayoutInfoDao
}

var (
	// PayoutInfo is globally public accessible object for table payout_info operations.
	PayoutInfo = payoutInfoDao{
		internal.NewPayoutInfoDao(),
	}
)

// Fill with you ideas below.
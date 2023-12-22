// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalGiftInfoDao is internal type for wrapping internal DAO implements.
type internalGiftInfoDao = *internal.GiftInfoDao

// giftInfoDao is the data access object for table gift_info.
// You can define custom methods on it to extend its functionality as you wish.
type giftInfoDao struct {
	internalGiftInfoDao
}

var (
	// GiftInfo is globally public accessible object for table gift_info operations.
	GiftInfo = giftInfoDao{
		internal.NewGiftInfoDao(),
	}
)

// Fill with you ideas below.

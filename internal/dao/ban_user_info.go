// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalBanUserInfoDao is internal type for wrapping internal DAO implements.
type internalBanUserInfoDao = *internal.BanUserInfoDao

// banUserInfoDao is the data access object for table ban_user_info.
// You can define custom methods on it to extend its functionality as you wish.
type banUserInfoDao struct {
	internalBanUserInfoDao
}

var (
	// BanUserInfo is globally public accessible object for table ban_user_info operations.
	BanUserInfo = banUserInfoDao{
		internal.NewBanUserInfoDao(),
	}
)

// Fill with you ideas below.

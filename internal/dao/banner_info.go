// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalBannerInfoDao is internal type for wrapping internal DAO implements.
type internalBannerInfoDao = *internal.BannerInfoDao

// bannerInfoDao is the data access object for table banner_info.
// You can define custom methods on it to extend its functionality as you wish.
type bannerInfoDao struct {
	internalBannerInfoDao
}

var (
	// BannerInfo is globally public accessible object for table banner_info operations.
	BannerInfo = bannerInfoDao{
		internal.NewBannerInfoDao(),
	}
)

// Fill with you ideas below.

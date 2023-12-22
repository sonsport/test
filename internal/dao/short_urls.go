// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalShortUrlsDao is internal type for wrapping internal DAO implements.
type internalShortUrlsDao = *internal.ShortUrlsDao

// shortUrlsDao is the data access object for table short_urls.
// You can define custom methods on it to extend its functionality as you wish.
type shortUrlsDao struct {
	internalShortUrlsDao
}

var (
	// ShortUrls is globally public accessible object for table short_urls operations.
	ShortUrls = shortUrlsDao{
		internal.NewShortUrlsDao(),
	}
)

// Fill with you ideas below.
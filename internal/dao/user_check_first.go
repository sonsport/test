// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalUserCheckFirstDao is internal type for wrapping internal DAO implements.
type internalUserCheckFirstDao = *internal.UserCheckFirstDao

// userCheckFirstDao is the data access object for table user_check_first.
// You can define custom methods on it to extend its functionality as you wish.
type userCheckFirstDao struct {
	internalUserCheckFirstDao
}

var (
	// UserCheckFirst is globally public accessible object for table user_check_first operations.
	UserCheckFirst = userCheckFirstDao{
		internal.NewUserCheckFirstDao(),
	}
)

// Fill with you ideas below.
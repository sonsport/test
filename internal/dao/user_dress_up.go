// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalUserDressUpDao is internal type for wrapping internal DAO implements.
type internalUserDressUpDao = *internal.UserDressUpDao

// userDressUpDao is the data access object for table user_dress_up.
// You can define custom methods on it to extend its functionality as you wish.
type userDressUpDao struct {
	internalUserDressUpDao
}

var (
	// UserDressUp is globally public accessible object for table user_dress_up operations.
	UserDressUp = userDressUpDao{
		internal.NewUserDressUpDao(),
	}
)

// Fill with you ideas below.
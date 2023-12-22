// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalShareNewUserInfoDao is internal type for wrapping internal DAO implements.
type internalShareNewUserInfoDao = *internal.ShareNewUserInfoDao

// shareNewUserInfoDao is the data access object for table share_new_user_info.
// You can define custom methods on it to extend its functionality as you wish.
type shareNewUserInfoDao struct {
	internalShareNewUserInfoDao
}

var (
	// ShareNewUserInfo is globally public accessible object for table share_new_user_info operations.
	ShareNewUserInfo = shareNewUserInfoDao{
		internal.NewShareNewUserInfoDao(),
	}
)

// Fill with you ideas below.
// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalShareNewUserRecordDao is internal type for wrapping internal DAO implements.
type internalShareNewUserRecordDao = *internal.ShareNewUserRecordDao

// shareNewUserRecordDao is the data access object for table share_new_user_record.
// You can define custom methods on it to extend its functionality as you wish.
type shareNewUserRecordDao struct {
	internalShareNewUserRecordDao
}

var (
	// ShareNewUserRecord is globally public accessible object for table share_new_user_record operations.
	ShareNewUserRecord = shareNewUserRecordDao{
		internal.NewShareNewUserRecordDao(),
	}
)

// Fill with you ideas below.

// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalAnchorWarningRecordDao is internal type for wrapping internal DAO implements.
type internalAnchorWarningRecordDao = *internal.AnchorWarningRecordDao

// anchorWarningRecordDao is the data access object for table anchor_warning_record.
// You can define custom methods on it to extend its functionality as you wish.
type anchorWarningRecordDao struct {
	internalAnchorWarningRecordDao
}

var (
	// AnchorWarningRecord is globally public accessible object for table anchor_warning_record operations.
	AnchorWarningRecord = anchorWarningRecordDao{
		internal.NewAnchorWarningRecordDao(),
	}
)

// Fill with you ideas below.

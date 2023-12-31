// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalLiveCloudRecordDao is internal type for wrapping internal DAO implements.
type internalLiveCloudRecordDao = *internal.LiveCloudRecordDao

// liveCloudRecordDao is the data access object for table live_cloud_record.
// You can define custom methods on it to extend its functionality as you wish.
type liveCloudRecordDao struct {
	internalLiveCloudRecordDao
}

var (
	// LiveCloudRecord is globally public accessible object for table live_cloud_record operations.
	LiveCloudRecord = liveCloudRecordDao{
		internal.NewLiveCloudRecordDao(),
	}
)

// Fill with you ideas below.

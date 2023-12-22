// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalRoomScoreRecordDao is internal type for wrapping internal DAO implements.
type internalRoomScoreRecordDao = *internal.RoomScoreRecordDao

// roomScoreRecordDao is the data access object for table room_score_record.
// You can define custom methods on it to extend its functionality as you wish.
type roomScoreRecordDao struct {
	internalRoomScoreRecordDao
}

var (
	// RoomScoreRecord is globally public accessible object for table room_score_record operations.
	RoomScoreRecord = roomScoreRecordDao{
		internal.NewRoomScoreRecordDao(),
	}
)

// Fill with you ideas below.
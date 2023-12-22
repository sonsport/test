// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalIncomeRecordDao is internal type for wrapping internal DAO implements.
type internalIncomeRecordDao = *internal.IncomeRecordDao

// incomeRecordDao is the data access object for table income_record.
// You can define custom methods on it to extend its functionality as you wish.
type incomeRecordDao struct {
	internalIncomeRecordDao
}

var (
	// IncomeRecord is globally public accessible object for table income_record operations.
	IncomeRecord = incomeRecordDao{
		internal.NewIncomeRecordDao(),
	}
)

// Fill with you ideas below.

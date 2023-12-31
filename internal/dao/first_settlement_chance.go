// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalFirstSettlementChanceDao is internal type for wrapping internal DAO implements.
type internalFirstSettlementChanceDao = *internal.FirstSettlementChanceDao

// firstSettlementChanceDao is the data access object for table first_settlement_chance.
// You can define custom methods on it to extend its functionality as you wish.
type firstSettlementChanceDao struct {
	internalFirstSettlementChanceDao
}

var (
	// FirstSettlementChance is globally public accessible object for table first_settlement_chance operations.
	FirstSettlementChance = firstSettlementChanceDao{
		internal.NewFirstSettlementChanceDao(),
	}
)

// Fill with you ideas below.

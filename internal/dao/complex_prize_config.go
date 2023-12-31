// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalComplexPrizeConfigDao is internal type for wrapping internal DAO implements.
type internalComplexPrizeConfigDao = *internal.ComplexPrizeConfigDao

// complexPrizeConfigDao is the data access object for table complex_prize_config.
// You can define custom methods on it to extend its functionality as you wish.
type complexPrizeConfigDao struct {
	internalComplexPrizeConfigDao
}

var (
	// ComplexPrizeConfig is globally public accessible object for table complex_prize_config operations.
	ComplexPrizeConfig = complexPrizeConfigDao{
		internal.NewComplexPrizeConfigDao(),
	}
)

// Fill with you ideas below.

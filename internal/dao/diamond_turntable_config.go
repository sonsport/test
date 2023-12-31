// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalDiamondTurntableConfigDao is internal type for wrapping internal DAO implements.
type internalDiamondTurntableConfigDao = *internal.DiamondTurntableConfigDao

// diamondTurntableConfigDao is the data access object for table diamond_turntable_config.
// You can define custom methods on it to extend its functionality as you wish.
type diamondTurntableConfigDao struct {
	internalDiamondTurntableConfigDao
}

var (
	// DiamondTurntableConfig is globally public accessible object for table diamond_turntable_config operations.
	DiamondTurntableConfig = diamondTurntableConfigDao{
		internal.NewDiamondTurntableConfigDao(),
	}
)

// Fill with you ideas below.

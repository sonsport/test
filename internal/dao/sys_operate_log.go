// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalSysOperateLogDao is internal type for wrapping internal DAO implements.
type internalSysOperateLogDao = *internal.SysOperateLogDao

// sysOperateLogDao is the data access object for table sys_operate_log.
// You can define custom methods on it to extend its functionality as you wish.
type sysOperateLogDao struct {
	internalSysOperateLogDao
}

var (
	// SysOperateLog is globally public accessible object for table sys_operate_log operations.
	SysOperateLog = sysOperateLogDao{
		internal.NewSysOperateLogDao(),
	}
)

// Fill with you ideas below.

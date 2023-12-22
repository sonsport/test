// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalAgentBalanceInfoDao is internal type for wrapping internal DAO implements.
type internalAgentBalanceInfoDao = *internal.AgentBalanceInfoDao

// agentBalanceInfoDao is the data access object for table agent_balance_info.
// You can define custom methods on it to extend its functionality as you wish.
type agentBalanceInfoDao struct {
	internalAgentBalanceInfoDao
}

var (
	// AgentBalanceInfo is globally public accessible object for table agent_balance_info operations.
	AgentBalanceInfo = agentBalanceInfoDao{
		internal.NewAgentBalanceInfoDao(),
	}
)

// Fill with you ideas below.

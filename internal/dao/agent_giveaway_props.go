// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalAgentGiveawayPropsDao is internal type for wrapping internal DAO implements.
type internalAgentGiveawayPropsDao = *internal.AgentGiveawayPropsDao

// agentGiveawayPropsDao is the data access object for table agent_giveaway_props.
// You can define custom methods on it to extend its functionality as you wish.
type agentGiveawayPropsDao struct {
	internalAgentGiveawayPropsDao
}

var (
	// AgentGiveawayProps is globally public accessible object for table agent_giveaway_props operations.
	AgentGiveawayProps = agentGiveawayPropsDao{
		internal.NewAgentGiveawayPropsDao(),
	}
)

// Fill with you ideas below.

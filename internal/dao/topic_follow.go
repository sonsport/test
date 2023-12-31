// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalTopicFollowDao is internal type for wrapping internal DAO implements.
type internalTopicFollowDao = *internal.TopicFollowDao

// topicFollowDao is the data access object for table topic_follow.
// You can define custom methods on it to extend its functionality as you wish.
type topicFollowDao struct {
	internalTopicFollowDao
}

var (
	// TopicFollow is globally public accessible object for table topic_follow operations.
	TopicFollow = topicFollowDao{
		internal.NewTopicFollowDao(),
	}
)

// Fill with you ideas below.

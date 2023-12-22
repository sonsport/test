// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalFeedCommentDao is internal type for wrapping internal DAO implements.
type internalFeedCommentDao = *internal.FeedCommentDao

// feedCommentDao is the data access object for table feed_comment.
// You can define custom methods on it to extend its functionality as you wish.
type feedCommentDao struct {
	internalFeedCommentDao
}

var (
	// FeedComment is globally public accessible object for table feed_comment operations.
	FeedComment = feedCommentDao{
		internal.NewFeedCommentDao(),
	}
)

// Fill with you ideas below.
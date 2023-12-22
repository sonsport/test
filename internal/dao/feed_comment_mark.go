// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"fuya-ark/internal/dao/internal"
)

// internalFeedCommentMarkDao is internal type for wrapping internal DAO implements.
type internalFeedCommentMarkDao = *internal.FeedCommentMarkDao

// feedCommentMarkDao is the data access object for table feed_comment_mark.
// You can define custom methods on it to extend its functionality as you wish.
type feedCommentMarkDao struct {
	internalFeedCommentMarkDao
}

var (
	// FeedCommentMark is globally public accessible object for table feed_comment_mark operations.
	FeedCommentMark = feedCommentMarkDao{
		internal.NewFeedCommentMarkDao(),
	}
)

// Fill with you ideas below.

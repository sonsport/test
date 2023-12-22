// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedInfoDao is the data access object for table feed_info.
type FeedInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FeedInfoColumns // columns contains all the column names of Table for convenient usage.
}

// FeedInfoColumns defines and stores column names for table feed_info.
type FeedInfoColumns struct {
	Id                string // 自增Id
	UserId            string // 用户Id
	Text              string // 动态内容
	Images            string // 动态图片组，采用数组存储,目前规定数组长度为9，即最多9张图片
	LikeCount         string // 点赞数
	State             string // 0未审核、1审核通过、2审核不通过、3系统删除、4用户自己删除
	Score             string // 动态排序权重
	CreateTime        string //
	UpdateTime        string //
	Video             string // 动态视频
	VideoPicture      string // 视频封面
	Type              string // 动态类型，1图片、2视频
	UploadVideoWidth  string //
	UploadVideoHeight string //
	IsVote            string // 是否pk动态
	OptionOne         string // pk选项1
	OptionTwo         string // pk选项2
	SystemOs          string // 发动态的设备类型
	AppChannel        string // 渠道包
}

// feedInfoColumns holds the columns for table feed_info.
var feedInfoColumns = FeedInfoColumns{
	Id:                "id",
	UserId:            "user_id",
	Text:              "text",
	Images:            "images",
	LikeCount:         "like_count",
	State:             "state",
	Score:             "score",
	CreateTime:        "create_time",
	UpdateTime:        "update_time",
	Video:             "video",
	VideoPicture:      "video_picture",
	Type:              "type",
	UploadVideoWidth:  "upload_video_width",
	UploadVideoHeight: "upload_video_height",
	IsVote:            "is_vote",
	OptionOne:         "option_one",
	OptionTwo:         "option_two",
	SystemOs:          "system_os",
	AppChannel:        "app_channel",
}

// NewFeedInfoDao creates and returns a new DAO object for table data access.
func NewFeedInfoDao() *FeedInfoDao {
	return &FeedInfoDao{
		group:   "default",
		table:   "feed_info",
		columns: feedInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeedInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeedInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeedInfoDao) Columns() FeedInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeedInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeedInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeedInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

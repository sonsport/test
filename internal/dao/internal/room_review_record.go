// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomReviewRecordDao is the data access object for table room_review_record.
type RoomReviewRecordDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns RoomReviewRecordColumns // columns contains all the column names of Table for convenient usage.
}

// RoomReviewRecordColumns defines and stores column names for table room_review_record.
type RoomReviewRecordColumns struct {
	Id               string //
	RoomId           string // 房间id
	UserId           string // 用户id
	StatCode         string // 审核状态。0：审核中1：审核结束
	RiskLevel        string // 风险级别  0：pass，正常内容， 1：review，可疑内容 2：reject，违规内容
	ContentType      string // 用来区分音频和图片回调。 1：该回调为图片回调 2：该回调为音频回调
	StreamId         string // 开始视频流审核时输入的流 ID。
	ProcessStartTime string // 开始处理时间戳（毫秒）。
	ProcessEndTime   string // 检测完成时间戳（毫秒）
	RiskType         string // 标识风险类型。 0：正常 100：涉政 200：色情 210：性感 300：广告 310：二维码 320：水印 400：暴恐 500：违规 510：不良场景 530：行业违规
	RiskSource       string // 风险来源。 1000：无风险。 1001：文字有风险（OCR 图片识别文字）。 1002：截帧有风险。
	MatchedDetail    string // 命中所有敏感词名单的详情。 调用 开始视频流审核 时，如果 ImageType[] 传值包括 5（广告识别）、10（OCR），会触发敏感词识别。
	MatchedItem      string // 命中的具体敏感词
	MatchedList      string // 命中敏感词所在的名单名称
	Description      string // 策略规则风险原因描述。
	DetectType       string // 用来区分截帧图片是否过了检测。 1：截帧图片过了检测 2：截帧图片没过检测
	Similarity       string // 与上一张截帧图片的相似概率值 取值范围 [0-1]，数值越接近 1 越相似。
	StillTime        string // 展示静止画面时间（秒）
	ImageUrl         string // 截帧图片 url 地址。
	ImageTime        string // 视频流截帧图片违规发生的时间戳（毫秒）。
	ImageText        string // 视频中画面识别出的文字内容
	CreateTime       string // 创建时间
	UpdateTime       string // 更新时间
}

// roomReviewRecordColumns holds the columns for table room_review_record.
var roomReviewRecordColumns = RoomReviewRecordColumns{
	Id:               "id",
	RoomId:           "room_id",
	UserId:           "user_id",
	StatCode:         "stat_code",
	RiskLevel:        "risk_level",
	ContentType:      "content_type",
	StreamId:         "stream_id",
	ProcessStartTime: "process_start_time",
	ProcessEndTime:   "process_end_time",
	RiskType:         "risk_type",
	RiskSource:       "risk_source",
	MatchedDetail:    "matched_detail",
	MatchedItem:      "matched_item",
	MatchedList:      "matched_list",
	Description:      "description",
	DetectType:       "detect_type",
	Similarity:       "similarity",
	StillTime:        "still_time",
	ImageUrl:         "image_url",
	ImageTime:        "image_time",
	ImageText:        "image_text",
	CreateTime:       "create_time",
	UpdateTime:       "update_time",
}

// NewRoomReviewRecordDao creates and returns a new DAO object for table data access.
func NewRoomReviewRecordDao() *RoomReviewRecordDao {
	return &RoomReviewRecordDao{
		group:   "default",
		table:   "room_review_record",
		columns: roomReviewRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomReviewRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomReviewRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomReviewRecordDao) Columns() RoomReviewRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomReviewRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomReviewRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomReviewRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

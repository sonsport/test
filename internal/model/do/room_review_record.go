// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomReviewRecord is the golang structure of table room_review_record for DAO operations like Where/Data.
type RoomReviewRecord struct {
	g.Meta           `orm:"table:room_review_record, do:true"`
	Id               interface{} //
	RoomId           interface{} // 房间id
	UserId           interface{} // 用户id
	StatCode         interface{} // 审核状态。0：审核中1：审核结束
	RiskLevel        interface{} // 风险级别  0：pass，正常内容， 1：review，可疑内容 2：reject，违规内容
	ContentType      interface{} // 用来区分音频和图片回调。 1：该回调为图片回调 2：该回调为音频回调
	StreamId         interface{} // 开始视频流审核时输入的流 ID。
	ProcessStartTime interface{} // 开始处理时间戳（毫秒）。
	ProcessEndTime   interface{} // 检测完成时间戳（毫秒）
	RiskType         interface{} // 标识风险类型。 0：正常 100：涉政 200：色情 210：性感 300：广告 310：二维码 320：水印 400：暴恐 500：违规 510：不良场景 530：行业违规
	RiskSource       interface{} // 风险来源。 1000：无风险。 1001：文字有风险（OCR 图片识别文字）。 1002：截帧有风险。
	MatchedDetail    interface{} // 命中所有敏感词名单的详情。 调用 开始视频流审核 时，如果 ImageType[] 传值包括 5（广告识别）、10（OCR），会触发敏感词识别。
	MatchedItem      interface{} // 命中的具体敏感词
	MatchedList      interface{} // 命中敏感词所在的名单名称
	Description      interface{} // 策略规则风险原因描述。
	DetectType       interface{} // 用来区分截帧图片是否过了检测。 1：截帧图片过了检测 2：截帧图片没过检测
	Similarity       interface{} // 与上一张截帧图片的相似概率值 取值范围 [0-1]，数值越接近 1 越相似。
	StillTime        interface{} // 展示静止画面时间（秒）
	ImageUrl         interface{} // 截帧图片 url 地址。
	ImageTime        interface{} // 视频流截帧图片违规发生的时间戳（毫秒）。
	ImageText        interface{} // 视频中画面识别出的文字内容
	CreateTime       interface{} // 创建时间
	UpdateTime       interface{} // 更新时间
}

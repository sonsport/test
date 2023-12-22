// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomReviewRecord is the golang structure for table room_review_record.
type RoomReviewRecord struct {
	Id               int64   `json:"id"               description:""`
	RoomId           string  `json:"roomId"           description:"房间id"`
	UserId           int64   `json:"userId"           description:"用户id"`
	StatCode         int     `json:"statCode"         description:"审核状态。0：审核中1：审核结束"`
	RiskLevel        int     `json:"riskLevel"        description:"风险级别  0：pass，正常内容， 1：review，可疑内容 2：reject，违规内容"`
	ContentType      int     `json:"contentType"      description:"用来区分音频和图片回调。 1：该回调为图片回调 2：该回调为音频回调"`
	StreamId         string  `json:"streamId"         description:"开始视频流审核时输入的流 ID。"`
	ProcessStartTime int64   `json:"processStartTime" description:"开始处理时间戳（毫秒）。"`
	ProcessEndTime   int64   `json:"processEndTime"   description:"检测完成时间戳（毫秒）"`
	RiskType         int     `json:"riskType"         description:"标识风险类型。 0：正常 100：涉政 200：色情 210：性感 300：广告 310：二维码 320：水印 400：暴恐 500：违规 510：不良场景 530：行业违规"`
	RiskSource       int     `json:"riskSource"       description:"风险来源。 1000：无风险。 1001：文字有风险（OCR 图片识别文字）。 1002：截帧有风险。"`
	MatchedDetail    string  `json:"matchedDetail"    description:"命中所有敏感词名单的详情。 调用 开始视频流审核 时，如果 ImageType[] 传值包括 5（广告识别）、10（OCR），会触发敏感词识别。"`
	MatchedItem      string  `json:"matchedItem"      description:"命中的具体敏感词"`
	MatchedList      string  `json:"matchedList"      description:"命中敏感词所在的名单名称"`
	Description      string  `json:"description"      description:"策略规则风险原因描述。"`
	DetectType       int     `json:"detectType"       description:"用来区分截帧图片是否过了检测。 1：截帧图片过了检测 2：截帧图片没过检测"`
	Similarity       float64 `json:"similarity"       description:"与上一张截帧图片的相似概率值 取值范围 [0-1]，数值越接近 1 越相似。"`
	StillTime        int     `json:"stillTime"        description:"展示静止画面时间（秒）"`
	ImageUrl         string  `json:"imageUrl"         description:"截帧图片 url 地址。"`
	ImageTime        int64   `json:"imageTime"        description:"视频流截帧图片违规发生的时间戳（毫秒）。"`
	ImageText        string  `json:"imageText"        description:"视频中画面识别出的文字内容"`
	CreateTime       int64   `json:"createTime"       description:"创建时间"`
	UpdateTime       int64   `json:"updateTime"       description:"更新时间"`
}

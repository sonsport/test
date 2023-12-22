// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RedEnvelopeInfo is the golang structure for table red_envelope_info.
type RedEnvelopeInfo struct {
	Id                 int64  `json:"id"                 description:""`
	UserId             int64  `json:"userId"             description:"用户id"`
	AnchorUserId       int64  `json:"anchorUserId"       description:"主播id"`
	RoomId             string `json:"roomId"             description:"房间id"`
	RedEnvelopeType    int    `json:"redEnvelopeType"    description:"红包类型 1 房间红包 2 全站红包"`
	RedEnvelopeCount   int    `json:"redEnvelopeCount"   description:"红包个数"`
	RedEnvelopeDiamond int    `json:"redEnvelopeDiamond" description:"红包钻石"`
	SpentDiamond       int    `json:"spentDiamond"       description:"已抢钻石"`
	BeginTime          int64  `json:"beginTime"          description:"开始时间"`
	EndTime            int64  `json:"endTime"            description:"结束时间"`
	CreateTime         int64  `json:"createTime"         description:""`
	UpdateTime         int64  `json:"updateTime"         description:""`
}

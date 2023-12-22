// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RedEnvelopeRecord is the golang structure for table red_envelope_record.
type RedEnvelopeRecord struct {
	Id            int64  `json:"id"            description:""`
	RId           int64  `json:"rId"           description:"红包id"`
	UserId        int64  `json:"userId"        description:"用户id"`
	RoomId        string `json:"roomId"        description:"房间id"`
	Diamond       int    `json:"diamond"       description:"红包钻石"`
	BeforeDiamond int    `json:"beforeDiamond" description:"红包之前余额"`
	AfterDiamond  int    `json:"afterDiamond"  description:"红包之后余额"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

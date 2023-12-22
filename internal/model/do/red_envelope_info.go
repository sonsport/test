// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RedEnvelopeInfo is the golang structure of table red_envelope_info for DAO operations like Where/Data.
type RedEnvelopeInfo struct {
	g.Meta             `orm:"table:red_envelope_info, do:true"`
	Id                 interface{} //
	UserId             interface{} // 用户id
	AnchorUserId       interface{} // 主播id
	RoomId             interface{} // 房间id
	RedEnvelopeType    interface{} // 红包类型 1 房间红包 2 全站红包
	RedEnvelopeCount   interface{} // 红包个数
	RedEnvelopeDiamond interface{} // 红包钻石
	SpentDiamond       interface{} // 已抢钻石
	BeginTime          interface{} // 开始时间
	EndTime            interface{} // 结束时间
	CreateTime         interface{} //
	UpdateTime         interface{} //
}

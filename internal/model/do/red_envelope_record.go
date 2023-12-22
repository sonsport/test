// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RedEnvelopeRecord is the golang structure of table red_envelope_record for DAO operations like Where/Data.
type RedEnvelopeRecord struct {
	g.Meta        `orm:"table:red_envelope_record, do:true"`
	Id            interface{} //
	RId           interface{} // 红包id
	UserId        interface{} // 用户id
	RoomId        interface{} // 房间id
	Diamond       interface{} // 红包钻石
	BeforeDiamond interface{} // 红包之前余额
	AfterDiamond  interface{} // 红包之后余额
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

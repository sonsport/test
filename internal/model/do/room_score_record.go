// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomScoreRecord is the golang structure of table room_score_record for DAO operations like Where/Data.
type RoomScoreRecord struct {
	g.Meta        `orm:"table:room_score_record, do:true"`
	Id            interface{} //
	UserId        interface{} //
	RoomId        interface{} //
	OperationType interface{} // 操作类型
	OperationId   interface{} // 操作id
	Score         interface{} // 排序分值
	BeforeScore   interface{} // before排序分值
	AfterScore    interface{} // after排序分值
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

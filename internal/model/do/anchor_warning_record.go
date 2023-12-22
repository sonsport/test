// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AnchorWarningRecord is the golang structure of table anchor_warning_record for DAO operations like Where/Data.
type AnchorWarningRecord struct {
	g.Meta      `orm:"table:anchor_warning_record, do:true"`
	Id          interface{} // 自增id
	Icon        interface{} // 警告图片证据
	Appeal      interface{} // 申诉
	State       interface{} // 0待申诉，1申诉中，2成功，3失败
	UserId      interface{} // 主播id
	RoomId      interface{} // 直播间id
	WarnType    interface{} // 警告类型
	Reason      interface{} // 警告原因
	WarnBigType interface{} //
	AdminId     interface{} // 提醒大类
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

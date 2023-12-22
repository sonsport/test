// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RoomActivityBanner is the golang structure of table room_activity_banner for DAO operations like Where/Data.
type RoomActivityBanner struct {
	g.Meta       `orm:"table:room_activity_banner, do:true"`
	Id           interface{} //
	Icon         interface{} // 图片
	ActivityType interface{} // 活动类型 1 首冲 2 h5
	Link         interface{} // 链接地址
	State        interface{} // 状态1启用
	ShowType     interface{} // 展示类型，0 全部，1 未充值用户 2 充值用户 3 主播
}

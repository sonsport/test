// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityPageConfig is the golang structure of table activity_page_config for DAO operations like Where/Data.
type ActivityPageConfig struct {
	g.Meta       `orm:"table:activity_page_config, do:true"`
	Id           interface{} //
	AcId         interface{} // 活动Id activity_config 自增id
	ResUrl       interface{} // 资源地址
	Duration     interface{} // 持续时间 单位秒 >0才倒计时
	LinkAddress  interface{} // 链接地址
	ShowStyle    interface{} // 0无配置 1全屏 2半屏 后续其他扩展
	PageType     interface{} // 页面类别：启动页(start_up)、首页(home_page)、直播间(room_live)、后续扩展
	PageLocation interface{} // 页面位置：启动页：只有一个位置不用配置，首页：弹窗（pop_up）、浮窗（float）；直播间：弹窗（pop_up）、浮窗（float）、底部图标（bottom_icon）、直播间-礼物（gift_top）；
	StartTime    interface{} // 活动开始时间
	EndTime      interface{} // 活动结束时间
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

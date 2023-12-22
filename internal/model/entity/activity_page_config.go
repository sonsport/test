// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityPageConfig is the golang structure for table activity_page_config.
type ActivityPageConfig struct {
	Id           int    `json:"id"           description:""`
	AcId         int    `json:"acId"         description:"活动Id activity_config 自增id"`
	ResUrl       string `json:"resUrl"       description:"资源地址"`
	Duration     int    `json:"duration"     description:"持续时间 单位秒 >0才倒计时"`
	LinkAddress  string `json:"linkAddress"  description:"链接地址"`
	ShowStyle    int    `json:"showStyle"    description:"0无配置 1全屏 2半屏 后续其他扩展"`
	PageType     string `json:"pageType"     description:"页面类别：启动页(start_up)、首页(home_page)、直播间(room_live)、后续扩展"`
	PageLocation string `json:"pageLocation" description:"页面位置：启动页：只有一个位置不用配置，首页：弹窗（pop_up）、浮窗（float）；直播间：弹窗（pop_up）、浮窗（float）、底部图标（bottom_icon）、直播间-礼物（gift_top）；"`
	StartTime    int64  `json:"startTime"    description:"活动开始时间"`
	EndTime      int64  `json:"endTime"      description:"活动结束时间"`
	CreateTime   int64  `json:"createTime"   description:""`
	UpdateTime   int64  `json:"updateTime"   description:""`
}

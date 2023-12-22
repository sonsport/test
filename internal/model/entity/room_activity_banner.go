// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomActivityBanner is the golang structure for table room_activity_banner.
type RoomActivityBanner struct {
	Id           int    `json:"id"           description:""`
	Icon         string `json:"icon"         description:"图片"`
	ActivityType int    `json:"activityType" description:"活动类型 1 首冲 2 h5"`
	Link         string `json:"link"         description:"链接地址"`
	State        int    `json:"state"        description:"状态1启用"`
	ShowType     int    `json:"showType"     description:"展示类型，0 全部，1 未充值用户 2 充值用户 3 主播"`
}

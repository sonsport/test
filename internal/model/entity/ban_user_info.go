// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BanUserInfo is the golang structure for table ban_user_info.
type BanUserInfo struct {
	Id         int64  `json:"id"         description:"自增id"`
	UserId     uint64 `json:"userId"     description:"封禁用户id"`
	AdminId    uint64 `json:"adminId"    description:"管理员id"`
	LinkId     int64  `json:"linkId"     description:"关联警告表主键id"`
	Type       int    `json:"type"       description:"封禁类型，1禁发私信，2禁发公屏聊天，3禁开播，4禁止登录，5封禁设备"`
	BanTime    int64  `json:"banTime"    description:"封禁时间"`
	DeviceId   string `json:"deviceId"   description:"封禁设备"`
	BanReason  string `json:"banReason"  description:"封禁原因"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

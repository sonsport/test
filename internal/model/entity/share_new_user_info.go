// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShareNewUserInfo is the golang structure for table share_new_user_info.
type ShareNewUserInfo struct {
	Id               uint64 `json:"id"               description:""`
	UserId           int64  `json:"userId"           description:""`
	ShareId          int    `json:"shareId"          description:"分享ID"`
	OwnerUserId      int64  `json:"ownerUserId"      description:"分享者uid"`
	WatchLiveTime    int    `json:"watchLiveTime"    description:"观看直播时长"`
	RegisterTime     int64  `json:"registerTime"     description:"注册时间"`
	Ip               string `json:"ip"               description:"登陆地IP"`
	Smid             string `json:"smid"             description:"设备ID"`
	DeviceName       string `json:"deviceName"       description:"设备型号"`
	Status           int    `json:"status"           description:"0:初始化；1:待加币；2:加币成功"`
	WatchLiveRewards int    `json:"watchLiveRewards" description:"1:看播奖励有效；2:看播奖励无效"`
	IsOldUser        int    `json:"isOldUser"        description:"1:老用户；0:新用户"`
	FullReasonCode   int    `json:"fullReasonCode"   description:"失败原因 status -1时有效"`
	CreateTime       int64  `json:"createTime"       description:""`
	UpdateTime       int64  `json:"updateTime"       description:""`
}

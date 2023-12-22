// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserBaseInfo is the golang structure for table user_base_info.
type UserBaseInfo struct {
	Id               int64  `json:"id"               description:""`
	UserId           uint64 `json:"userId"           description:"用户Id"`
	FirstLiveTime    int64  `json:"firstLiveTime"    description:"首播日期"`
	FirstAuthTime    int64  `json:"firstAuthTime"    description:"首次认证日期"`
	LastAuthTime     int64  `json:"lastAuthTime"     description:"最后认证日期"`
	LiveCallSwitch   int    `json:"liveCallSwitch"   description:"连麦开关 1打开"`
	LiveCallType     int    `json:"liveCallType"     description:"连麦类型 1 所有人 2 关注我的人 3我的好友"`
	ReviewUpUser     int    `json:"reviewUpUser"     description:"多人房上麦审核"`
	MultipleUserType int    `json:"multipleUserType" description:"多人房上麦权限  1 所有人 2 关注我的人 3我的好友"`
	CreatedTime      int64  `json:"createdTime"      description:""`
	UpdatedTime      int64  `json:"updatedTime"      description:""`
	ServerId         int64  `json:"serverId"         description:"运营号id"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivePersonStatistics is the golang structure for table active_person_statistics.
type ActivePersonStatistics struct {
	Id               uint   `json:"id"               description:"主键ID"`
	BeginTime        int64  `json:"beginTime"        description:"开始时间"`
	EndTime          int64  `json:"endTime"          description:"结束时间"`
	AppChannel       string `json:"appChannel"       description:"渠道"`
	AppVersion       string `json:"appVersion"       description:"app版本"`
	ActiveType       int    `json:"activeType"       description:"类型1日活，2周活，3月活"`
	ActiveUsersCount int64  `json:"activeUsersCount" description:"活跃用户"`
}

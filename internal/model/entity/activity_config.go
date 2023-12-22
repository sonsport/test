// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ActivityConfig is the golang structure for table activity_config.
type ActivityConfig struct {
	Id          int    `json:"id"          description:""`
	Title       string `json:"title"       description:"活动标题"`
	ResUrl      string `json:"resUrl"      description:"资源图 活动横图"`
	LinkAddress string `json:"linkAddress" description:"活动地址"`
	StartTime   int64  `json:"startTime"   description:"活动开始时间"`
	EndTime     int64  `json:"endTime"     description:"活动结束时间"`
	State       int    `json:"state"       description:"状态 0下架 1上架"`
	Sort        int    `json:"sort"        description:"排序，顺序排列"`
	Remark      string `json:"remark"      description:"备注"`
	CreateTime  int64  `json:"createTime"  description:""`
	UpdateTime  int64  `json:"updateTime"  description:""`
}

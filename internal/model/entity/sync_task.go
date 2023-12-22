// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SyncTask is the golang structure for table sync_task.
type SyncTask struct {
	Id        uint64      `json:"id"        description:"任务id"`
	SrcTable  string      `json:"srcTable"  description:"数据来源表"`
	DstTable  string      `json:"dstTable"  description:"数据去向表"`
	Remark    string      `json:"remark"    description:"备注"`
	SyncId    uint64      `json:"syncId"    description:"已同步/统计的数据来源表的最大记录id"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}

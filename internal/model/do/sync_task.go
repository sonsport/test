// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SyncTask is the golang structure of table sync_task for DAO operations like Where/Data.
type SyncTask struct {
	g.Meta    `orm:"table:sync_task, do:true"`
	Id        interface{} // 任务id
	SrcTable  interface{} // 数据来源表
	DstTable  interface{} // 数据去向表
	Remark    interface{} // 备注
	SyncId    interface{} // 已同步/统计的数据来源表的最大记录id
	UpdatedAt *gtime.Time //
}

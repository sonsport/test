// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityConfig is the golang structure of table activity_config for DAO operations like Where/Data.
type ActivityConfig struct {
	g.Meta      `orm:"table:activity_config, do:true"`
	Id          interface{} //
	Title       interface{} // 活动标题
	ResUrl      interface{} // 资源图 活动横图
	LinkAddress interface{} // 活动地址
	StartTime   interface{} // 活动开始时间
	EndTime     interface{} // 活动结束时间
	State       interface{} // 状态 0下架 1上架
	Sort        interface{} // 排序，顺序排列
	Remark      interface{} // 备注
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

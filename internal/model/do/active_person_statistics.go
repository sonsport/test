// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ActivePersonStatistics is the golang structure of table active_person_statistics for DAO operations like Where/Data.
type ActivePersonStatistics struct {
	g.Meta           `orm:"table:active_person_statistics, do:true"`
	Id               interface{} // 主键ID
	BeginTime        interface{} // 开始时间
	EndTime          interface{} // 结束时间
	AppChannel       interface{} // 渠道
	AppVersion       interface{} // app版本
	ActiveType       interface{} // 类型1日活，2周活，3月活
	ActiveUsersCount interface{} // 活跃用户
}

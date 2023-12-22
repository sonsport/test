// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MedalInfo is the golang structure of table medal_info for DAO operations like Where/Data.
type MedalInfo struct {
	g.Meta          `orm:"table:medal_info, do:true"`
	Id              interface{} //
	MedalIcon       interface{} // 勋章图片
	MedalGreyIcon   interface{} // 勋章灰色图片
	MedalNameConfig interface{} // 勋章名称配置json
	MedalDescConfig interface{} // 勋章描述配置
	ConditionsType  interface{} // 获取类型
	BeginTime       interface{} // 获取勋章开始时间
	EndTime         interface{} // 获取勋章结束时间
	MedalStates     interface{} // 勋章状态 1
	MedalSort       interface{} // 排序
	Remark          interface{} // 备注
	CreateTime      interface{} //
	UpdateTime      interface{} //
}

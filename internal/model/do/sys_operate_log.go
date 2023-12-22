// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysOperateLog is the golang structure of table sys_operate_log for DAO operations like Where/Data.
type SysOperateLog struct {
	g.Meta    `orm:"table:sys_operate_log, do:true"`
	Id        interface{} // ID
	AdminId   interface{} // 用户ID
	BeginTime interface{} // 开始时间
	EndTime   interface{} // 结束时间
	Ip        interface{} // 请求地址
	Uri       interface{} // ip
	Params    interface{} // 请求参数
	Method    interface{} // HTTP方法
	Response  interface{} // 返回信息
	Host      interface{} // 请求主机
	Brower    interface{} // 浏览器
}

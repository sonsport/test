// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysOperateLog is the golang structure for table sys_operate_log.
type SysOperateLog struct {
	Id        int64  `json:"id"        description:"ID"`
	AdminId   int64  `json:"adminId"   description:"用户ID"`
	BeginTime int64  `json:"beginTime" description:"开始时间"`
	EndTime   int64  `json:"endTime"   description:"结束时间"`
	Ip        string `json:"ip"        description:"请求地址"`
	Uri       string `json:"uri"       description:"ip"`
	Params    string `json:"params"    description:"请求参数"`
	Method    string `json:"method"    description:"HTTP方法"`
	Response  string `json:"response"  description:"返回信息"`
	Host      string `json:"host"      description:"请求主机"`
	Brower    string `json:"brower"    description:"浏览器"`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdjustMessage is the golang structure for table adjust_message.
type AdjustMessage struct {
	Id          uint64 `json:"id"          description:""`
	InstallTime int64  `json:"installTime" description:"install_time"`
	UserId      int64  `json:"userId"      description:"用户Id"`
	UserName    string `json:"userName"    description:""`
	OrderId     string `json:"orderId"     description:"用户Id"`
	EventToken  string `json:"eventToken"  description:""`
	GaId        string `json:"gaId"        description:"订单id"`
	DeviceId    string `json:"deviceId"    description:"设备id"`
	AppName     string `json:"appName"     description:""`
	AppId       string `json:"appId"       description:""`
	Tracker     string `json:"tracker"     description:""`
	Network     string `json:"network"     description:""`
	AdjustId    string `json:"adjustId"    description:""`
	Currency    string `json:"currency"    description:""`
	Revenue     string `json:"revenue"     description:""`
	CreateTime  int64  `json:"createTime"  description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LoginRecord is the golang structure for table login_record.
type LoginRecord struct {
	Id         uint64 `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:"用户id"`
	AppName    string `json:"appName"    description:"app名称"`
	AppChannel string `json:"appChannel" description:"app渠道"`
	AppVersion string `json:"appVersion" description:"app版本"`
	DeviceId   string `json:"deviceId"   description:"设备号"`
	Ip         string `json:"ip"         description:"登陆ip"`
	Language   string `json:"language"   description:"语言"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
	AppToken   string `json:"appToken"   description:"app登录token"`
	LoginType  int    `json:"loginType"  description:"1 gg 2 fb 3 apple"`
	AppPhone   string `json:"appPhone"   description:"app手机型号"`
}

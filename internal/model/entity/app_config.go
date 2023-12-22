// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AppConfig is the golang structure for table app_config.
type AppConfig struct {
	Id         int    `json:"id"         description:"自增Id"`
	System     int    `json:"system"     description:"所属系统；0所有,1安卓 2ios"`
	Version    string `json:"version"    description:"版本"`
	Channel    string `json:"channel"    description:"渠道"`
	Key        string `json:"key"        description:"key"`
	Value      string `json:"value"      description:"值"`
	Remark     string `json:"remark"     description:"备注"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

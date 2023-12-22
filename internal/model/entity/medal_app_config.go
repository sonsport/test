// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MedalAppConfig is the golang structure for table medal_app_config.
type MedalAppConfig struct {
	Id              int64  `json:"id"              description:"主键id"`
	MedalId         int64  `json:"medalId"         description:"勋章id"`
	AppName         string `json:"appName"         description:"app名称"`
	MedalNameConfig string `json:"medalNameConfig" description:"勋章名称配置json"`
	MedalDescConfig string `json:"medalDescConfig" description:"勋章描述配置json"`
}

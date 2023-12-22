// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MedalAppConfig is the golang structure of table medal_app_config for DAO operations like Where/Data.
type MedalAppConfig struct {
	g.Meta          `orm:"table:medal_app_config, do:true"`
	Id              interface{} // 主键id
	MedalId         interface{} // 勋章id
	AppName         interface{} // app名称
	MedalNameConfig interface{} // 勋章名称配置json
	MedalDescConfig interface{} // 勋章描述配置json
}

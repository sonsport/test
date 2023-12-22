// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MedalInfo is the golang structure for table medal_info.
type MedalInfo struct {
	Id              int64  `json:"id"              description:""`
	MedalIcon       string `json:"medalIcon"       description:"勋章图片"`
	MedalGreyIcon   string `json:"medalGreyIcon"   description:"勋章灰色图片"`
	MedalNameConfig string `json:"medalNameConfig" description:"勋章名称配置json"`
	MedalDescConfig string `json:"medalDescConfig" description:"勋章描述配置"`
	ConditionsType  int    `json:"conditionsType"  description:"获取类型"`
	BeginTime       int64  `json:"beginTime"       description:"获取勋章开始时间"`
	EndTime         int64  `json:"endTime"         description:"获取勋章结束时间"`
	MedalStates     int    `json:"medalStates"     description:"勋章状态 1"`
	MedalSort       int    `json:"medalSort"       description:"排序"`
	Remark          string `json:"remark"          description:"备注"`
	CreateTime      int64  `json:"createTime"      description:""`
	UpdateTime      int64  `json:"updateTime"      description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NobleInfo is the golang structure for table noble_info.
type NobleInfo struct {
	Id            int    `json:"id"            description:""`
	NobleName     string `json:"nobleName"     description:"贵族名称"`
	NobleStates   int    `json:"nobleStates"   description:"贵族状态 1 上架"`
	NobleIcon     string `json:"nobleIcon"     description:"贵族图标"`
	CurrentPrice  int    `json:"currentPrice"  description:"现价"`
	OriginalPrice int    `json:"originalPrice" description:"原价"`
	NobleType     int    `json:"nobleType"     description:"贵族类型 1 vip"`
	NobleLevel    int    `json:"nobleLevel"    description:"贵族等级"`
	NobleDay      int    `json:"nobleDay"      description:"贵族天数"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

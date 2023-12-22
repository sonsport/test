// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NobleConfig is the golang structure for table noble_config.
type NobleConfig struct {
	Id            int64 `json:"id"            description:""`
	NobleId       int64 `json:"nobleId"       description:"贵族id"`
	PrivilegeId   int64 `json:"privilegeId"   description:"特权id"`
	PrivilegeLink int64 `json:"privilegeLink" description:"特权值 类型为加成类型时val为倍数、商品时为商品id"`
	CreateTime    int64 `json:"createTime"    description:""`
	UpdateTime    int64 `json:"updateTime"    description:""`
}

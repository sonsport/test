// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserDressUp is the golang structure for table user_dress_up.
type UserDressUp struct {
	Id         uint64 `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:"用户Id"`
	ShopId     uint   `json:"shopId"     description:"商品ID"`
	CategoryId uint   `json:"categoryId" description:"分类ID"`
	StartTime  int64  `json:"startTime"  description:"开始时间"`
	EndTime    int64  `json:"endTime"    description:"结束时间 -1标识永久配置"`
	Status     int    `json:"status"     description:"状态 0-未佩戴 1-佩戴"`
	EffectDays int    `json:"effectDays" description:"有效期 （单位：天）-1标识永久"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

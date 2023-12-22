// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserDressUpLog is the golang structure for table user_dress_up_log.
type UserDressUpLog struct {
	Id         uint64 `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:"用户Id"`
	ShopId     uint   `json:"shopId"     description:"商品ID"`
	CategoryId uint   `json:"categoryId" description:"分类ID"`
	DressType  int    `json:"dressType"  description:"0：卸载装扮 1：穿戴装扮"`
	CreateTime int64  `json:"createTime" description:"创建时间"`
	UpdateTime int64  `json:"updateTime" description:"更新时间"`
}

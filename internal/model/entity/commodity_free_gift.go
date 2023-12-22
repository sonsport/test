// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommodityFreeGift is the golang structure for table commodity_free_gift.
type CommodityFreeGift struct {
	Id       int64 `json:"id"       description:""`
	Cid      int64 `json:"cid"      description:""`
	State    int   `json:"state"    description:"状态 1 启用"`
	GiftType int   `json:"giftType" description:"赠送类型	1 装扮  2 经验"`
	GiftLink int64 `json:"giftLink" description:"赠送关联	商品id"`
	GiftVal  int   `json:"giftVal"  description:"赠送数	装扮为天 经验为经验值"`
}

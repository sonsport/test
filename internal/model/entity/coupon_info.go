// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CouponInfo is the golang structure for table coupon_info.
type CouponInfo struct {
	Id          int64 `json:"id"          description:""`
	UserId      int64 `json:"userId"      description:"用户Id"`
	BeginTime   int64 `json:"beginTime"   description:"开始时间"`
	EndTime     int64 `json:"endTime"     description:"结束时间"`
	CouponState int   `json:"couponState" description:"状态 1未使用 2已使用"`
	UseTime     int64 `json:"useTime"     description:"使用时间"`
	CouponType  int   `json:"couponType"  description:"优惠券类型 1 钻石赠送券"`
	CouponRatio int   `json:"couponRatio" description:"赠送比例 30 则赠送订单总额的30%"`
	MinShift    int   `json:"minShift"    description:"最小支持充值档位"`
	MaxShift    int   `json:"maxShift"    description:"最大支持充值档位"`
	LinkId      int64 `json:"linkId"      description:"优惠券来源"`
	CreateTime  int64 `json:"createTime"  description:""`
	UpdateTime  int64 `json:"updateTime"  description:""`
}

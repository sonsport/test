// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NobleUser is the golang structure for table noble_user.
type NobleUser struct {
	Id                  int64 `json:"id"                  description:""`
	UserId              int64 `json:"userId"              description:""`
	NobleId             int64 `json:"nobleId"             description:"贵族id"`
	ExpirationTime      int64 `json:"expirationTime"      description:"过期时间"`
	NobleStates         int   `json:"nobleStates"         description:"贵族状态 1 有效"`
	PaidThenDiamonds    int64 `json:"paidThenDiamonds"    description:"购买vip时的付费钻石数"`
	PaidProductDay      int64 `json:"paidProductDay"      description:"购买的规格天"`
	PaidProductDiamond  int64 `json:"paidProductDiamond"  description:"购买的钻石数"`
	RenewInspectionTime int64 `json:"renewInspectionTime" description:"自动续费检查时间"`
	CreateTime          int64 `json:"createTime"          description:""`
	UpdateTime          int64 `json:"updateTime"          description:""`
}

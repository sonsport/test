// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CouponInfo is the golang structure of table coupon_info for DAO operations like Where/Data.
type CouponInfo struct {
	g.Meta      `orm:"table:coupon_info, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户Id
	BeginTime   interface{} // 开始时间
	EndTime     interface{} // 结束时间
	CouponState interface{} // 状态 1未使用 2已使用
	UseTime     interface{} // 使用时间
	CouponType  interface{} // 优惠券类型 1 钻石赠送券
	CouponRatio interface{} // 赠送比例 30 则赠送订单总额的30%
	MinShift    interface{} // 最小支持充值档位
	MaxShift    interface{} // 最大支持充值档位
	LinkId      interface{} // 优惠券来源
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

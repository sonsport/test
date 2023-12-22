// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RechargePointsRecord is the golang structure of table recharge_points_record for DAO operations like Where/Data.
type RechargePointsRecord struct {
	g.Meta       `orm:"table:recharge_points_record, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户id
	Points       interface{} // 积分
	PointsType   interface{} // 积分类型
	BeforePoints interface{} // 之前积分
	AfterPoints  interface{} // 之后积分
	BeforeRank   interface{} // 之前排名
	AfterRank    interface{} // 之后排名
	OrderId      interface{} // 订单id
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

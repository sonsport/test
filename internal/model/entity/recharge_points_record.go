// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RechargePointsRecord is the golang structure for table recharge_points_record.
type RechargePointsRecord struct {
	Id           int   `json:"id"           description:""`
	UserId       int64 `json:"userId"       description:"用户id"`
	Points       int   `json:"points"       description:"积分"`
	PointsType   int   `json:"pointsType"   description:"积分类型"`
	BeforePoints int   `json:"beforePoints" description:"之前积分"`
	AfterPoints  int   `json:"afterPoints"  description:"之后积分"`
	BeforeRank   int   `json:"beforeRank"   description:"之前排名"`
	AfterRank    int   `json:"afterRank"    description:"之后排名"`
	OrderId      int64 `json:"orderId"      description:"订单id"`
	CreateTime   int64 `json:"createTime"   description:""`
	UpdateTime   int64 `json:"updateTime"   description:""`
}

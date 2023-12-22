// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ThirdPartyBusiGameOrder is the golang structure for table third_party_busi_game_order.
type ThirdPartyBusiGameOrder struct {
	Id          uint   `json:"id"          description:""`
	OrderType   int    `json:"orderType"   description:"1   : 增加钻石, 2:扣减钻石"`
	BusiId      string `json:"busiId"      description:"第三方业务唯一标识"`
	HimiUid     int    `json:"himiUid"     description:"uid"`
	CoinNum     int    `json:"coinNum"     description:"订单虚拟币额度"`
	BusiUid     string `json:"busiUid"     description:"第三方业务用户唯一标识"`
	OrderStatus int    `json:"orderStatus" description:"订单状态 0-下单,1-执行成功"`
	RoomId      string `json:"roomId"      description:"房间id"`
	CreateTime  int64  `json:"createTime"  description:""`
	UpdateTime  int64  `json:"updateTime"  description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ThirdPartyBusiGameOrder is the golang structure of table third_party_busi_game_order for DAO operations like Where/Data.
type ThirdPartyBusiGameOrder struct {
	g.Meta      `orm:"table:third_party_busi_game_order, do:true"`
	Id          interface{} //
	OrderType   interface{} // 1   : 增加钻石, 2:扣减钻石
	BusiId      interface{} // 第三方业务唯一标识
	HimiUid     interface{} // uid
	CoinNum     interface{} // 订单虚拟币额度
	BusiUid     interface{} // 第三方业务用户唯一标识
	OrderStatus interface{} // 订单状态 0-下单,1-执行成功
	RoomId      interface{} // 房间id
	CreateTime  interface{} //
	UpdateTime  interface{} //
}

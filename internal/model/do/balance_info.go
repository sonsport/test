// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceInfo is the golang structure of table balance_info for DAO operations like Where/Data.
type BalanceInfo struct {
	g.Meta              `orm:"table:balance_info, do:true"`
	UserId              interface{} // 用户Id
	TotalFee            interface{} // 总花费、只有付费购买才累计到这里 印尼盾 单位：分
	TotalPayoutDiamonds interface{} // 总流水，凡是用户花费的钻石都记录到这里 累计
	TotalPayDiamonds    interface{} // 总购买钻石 只累计购买的
	RemainDiamonds      interface{} // 剩余钻石数
	CreateTime          interface{} //
	UpdateTime          interface{} //
}

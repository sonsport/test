// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PayoutRecord is the golang structure of table payout_record for DAO operations like Where/Data.
type PayoutRecord struct {
	g.Meta      `orm:"table:payout_record, do:true"`
	Id          interface{} //
	PayoutId    interface{} // 代付id
	OrderNo     interface{} // 系统订单号
	TradeNo     interface{} // 交易订单号
	Status      interface{} // 状态
	FailureCode interface{} // 失败状态
	CreateTime  interface{} // 创建时间
	UpdateTime  interface{} // 修改时间
}

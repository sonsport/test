// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GoogleOrder is the golang structure of table google_order for DAO operations like Where/Data.
type GoogleOrder struct {
	g.Meta        `orm:"table:google_order, do:true"`
	Id            interface{} //
	UserId        interface{} // 用户Id
	Receipt       interface{} // inAppPurchaseData 回调
	Signature     interface{} // google签名
	InternalIp    interface{} // 总购买钻石 只累计购买的
	ProductId     interface{} // 商品Id
	GpOrderId     interface{} // google订单idGPA.0000
	PurchaseToken interface{} //
	AppName       interface{} // app名称
	Package       interface{} // 包名 com.aaa.bb
	OrderId       interface{} // 系统订单ID
	State         interface{} // 状态0待处理 1处理完成
	CreateTime    interface{} //
	UpdateTime    interface{} //
}

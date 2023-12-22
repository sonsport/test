// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GoogleOrder is the golang structure for table google_order.
type GoogleOrder struct {
	Id            int64  `json:"id"            description:""`
	UserId        int64  `json:"userId"        description:"用户Id"`
	Receipt       string `json:"receipt"       description:"inAppPurchaseData 回调"`
	Signature     string `json:"signature"     description:"google签名"`
	InternalIp    uint64 `json:"internalIp"    description:"总购买钻石 只累计购买的"`
	ProductId     string `json:"productId"     description:"商品Id"`
	GpOrderId     string `json:"gpOrderId"     description:"google订单idGPA.0000"`
	PurchaseToken string `json:"purchaseToken" description:""`
	AppName       string `json:"appName"       description:"app名称"`
	Package       string `json:"package"       description:"包名 com.aaa.bb"`
	OrderId       string `json:"orderId"       description:"系统订单ID"`
	State         int    `json:"state"         description:"状态0待处理 1处理完成"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShareNewUserRecord is the golang structure of table share_new_user_record for DAO operations like Where/Data.
type ShareNewUserRecord struct {
	g.Meta            `orm:"table:share_new_user_record, do:true"`
	Id                interface{} //
	UserId            interface{} //
	OwnerUid          interface{} // 分享者uid
	Type              interface{} // 1:新用户注册；2:新用户充值
	Status            interface{} // 0:初始化；1:待加币；2:加币成功
	OwnerOrderId      interface{} // 分享者加币订单号
	OwnerDiamondNum   interface{} // 分享者获得的钻石数
	NewUserOrderId    interface{} // 新用户加币订单
	NewUserDiamondNum interface{} // 新用户获得的钻石数
	CodeId            interface{} // 分享ID
	RegisterTime      interface{} // 注册时间
	RechargeNum       interface{} // 充值的虚拟币数量
	WatchLiveTime     interface{} // 观看直播时长
	Ip                interface{} // 登陆地IP
	Smid              interface{} // 设备ID
	DeviceName        interface{} // 设备型号
	UpdateTime        interface{} //
	CreateTime        interface{} //
}

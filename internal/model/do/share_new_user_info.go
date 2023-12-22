// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShareNewUserInfo is the golang structure of table share_new_user_info for DAO operations like Where/Data.
type ShareNewUserInfo struct {
	g.Meta           `orm:"table:share_new_user_info, do:true"`
	Id               interface{} //
	UserId           interface{} //
	ShareId          interface{} // 分享ID
	OwnerUserId      interface{} // 分享者uid
	WatchLiveTime    interface{} // 观看直播时长
	RegisterTime     interface{} // 注册时间
	Ip               interface{} // 登陆地IP
	Smid             interface{} // 设备ID
	DeviceName       interface{} // 设备型号
	Status           interface{} // 0:初始化；1:待加币；2:加币成功
	WatchLiveRewards interface{} // 1:看播奖励有效；2:看播奖励无效
	IsOldUser        interface{} // 1:老用户；0:新用户
	FullReasonCode   interface{} // 失败原因 status -1时有效
	CreateTime       interface{} //
	UpdateTime       interface{} //
}

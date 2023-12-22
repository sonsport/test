// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserBaseInfo is the golang structure of table user_base_info for DAO operations like Where/Data.
type UserBaseInfo struct {
	g.Meta           `orm:"table:user_base_info, do:true"`
	Id               interface{} //
	UserId           interface{} // 用户Id
	FirstLiveTime    interface{} // 首播日期
	FirstAuthTime    interface{} // 首次认证日期
	LastAuthTime     interface{} // 最后认证日期
	LiveCallSwitch   interface{} // 连麦开关 1打开
	LiveCallType     interface{} // 连麦类型 1 所有人 2 关注我的人 3我的好友
	ReviewUpUser     interface{} // 多人房上麦审核
	MultipleUserType interface{} // 多人房上麦权限  1 所有人 2 关注我的人 3我的好友
	CreatedTime      interface{} //
	UpdatedTime      interface{} //
	ServerId         interface{} // 运营号id
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NobleUser is the golang structure of table noble_user for DAO operations like Where/Data.
type NobleUser struct {
	g.Meta              `orm:"table:noble_user, do:true"`
	Id                  interface{} //
	UserId              interface{} //
	NobleId             interface{} // 贵族id
	ExpirationTime      interface{} // 过期时间
	NobleStates         interface{} // 贵族状态 1 有效
	PaidThenDiamonds    interface{} // 购买vip时的付费钻石数
	PaidProductDay      interface{} // 购买的规格天
	PaidProductDiamond  interface{} // 购买的钻石数
	RenewInspectionTime interface{} // 自动续费检查时间
	CreateTime          interface{} //
	UpdateTime          interface{} //
}

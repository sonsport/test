// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserDressUp is the golang structure of table user_dress_up for DAO operations like Where/Data.
type UserDressUp struct {
	g.Meta     `orm:"table:user_dress_up, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户Id
	ShopId     interface{} // 商品ID
	CategoryId interface{} // 分类ID
	StartTime  interface{} // 开始时间
	EndTime    interface{} // 结束时间 -1标识永久配置
	Status     interface{} // 状态 0-未佩戴 1-佩戴
	EffectDays interface{} // 有效期 （单位：天）-1标识永久
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserDressUpLog is the golang structure of table user_dress_up_log for DAO operations like Where/Data.
type UserDressUpLog struct {
	g.Meta     `orm:"table:user_dress_up_log, do:true"`
	Id         interface{} //
	UserId     interface{} // 用户Id
	ShopId     interface{} // 商品ID
	CategoryId interface{} // 分类ID
	DressType  interface{} // 0：卸载装扮 1：穿戴装扮
	CreateTime interface{} // 创建时间
	UpdateTime interface{} // 更新时间
}

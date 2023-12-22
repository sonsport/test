// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BannerInfo is the golang structure of table banner_info for DAO operations like Where/Data.
type BannerInfo struct {
	g.Meta     `orm:"table:banner_info, do:true"`
	Bid        interface{} // bid
	Ranking    interface{} // 排序
	Title      interface{} // 标题
	ClientOs   interface{} // 0 全部；1 android；2 iOS
	AppName    interface{} // app名称
	Cover      interface{} // 封面URL
	Target     interface{} // 跳转类型，0 网页，1 app内页
	Link       interface{} // 跳转URL
	ShowType   interface{} // 展示类型，0 全部，1 未充值用户 2 充值用户 3 主播
	CreateTime interface{} //
	UpdateTime interface{} //
	State      interface{} // 状态1有效、2失效
}

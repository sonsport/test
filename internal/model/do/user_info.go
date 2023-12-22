// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta             `orm:"table:user_info, do:true"`
	UserId             interface{} // 用户Id 自增
	Username           interface{} // 用户名
	AreaCode           interface{} // 地区编码
	Nickname           interface{} // 昵称
	Password           interface{} // 密码  空字符串说明没有密码
	Gender             interface{} // 性别 1男 2女 0未知
	Birthday           interface{} // 生日 时间戳
	Height             interface{} // 身高 cm
	Education          interface{} // 学历，枚举 小学 高中 大学 本科 扩展
	Industry           interface{} // 行业，枚举 学生、文化/艺术 扩展
	Intro              interface{} // 自我介绍 签名
	Portrait           interface{} // 头像
	LiveState          interface{} // 0未开播 1开播
	WhatsApp           interface{} // whatsapp号码
	Role               interface{} // 0普通用户 1主播 2管理员 3公会长 4运营号 后续扩展
	BanTime            interface{} // 封禁时间，0代表未封禁
	LastHeartbeatAt    interface{} // 用户心跳
	DeviceId           interface{} // 设备号
	AppName            interface{} // app名称
	AppChannel         interface{} // 渠道
	AppVersion         interface{} // app版本
	PhoneMode          interface{} // 手机型号
	PhoneSystemVersion interface{} // 手机系统版本
	LastIp             interface{} // 最后一次登录ip
	System             interface{} // 1为安卓 2为ios
	UserLanguage       interface{} // 用户语言
	CurrencyCode       interface{} // 国家数字码，比如印尼：360；马来：458  0为默认
	Remarks            interface{} // 备注
	CreateTime         interface{} // 创建时间
	UpdateTime         interface{} // 最后更新时间
}

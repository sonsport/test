// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	UserId             uint64 `json:"userId"             description:"用户Id 自增"`
	Username           string `json:"username"           description:"用户名"`
	AreaCode           int    `json:"areaCode"           description:"地区编码"`
	Nickname           string `json:"nickname"           description:"昵称"`
	Password           string `json:"password"           description:"密码  空字符串说明没有密码"`
	Gender             int    `json:"gender"             description:"性别 1男 2女 0未知"`
	Birthday           uint64 `json:"birthday"           description:"生日 时间戳"`
	Height             uint   `json:"height"             description:"身高 cm"`
	Education          uint   `json:"education"          description:"学历，枚举 小学 高中 大学 本科 扩展"`
	Industry           uint   `json:"industry"           description:"行业，枚举 学生、文化/艺术 扩展"`
	Intro              string `json:"intro"              description:"自我介绍 签名"`
	Portrait           string `json:"portrait"           description:"头像"`
	LiveState          uint   `json:"liveState"          description:"0未开播 1开播"`
	WhatsApp           string `json:"whatsApp"           description:"whatsapp号码"`
	Role               int    `json:"role"               description:"0普通用户 1主播 2管理员 3公会长 4运营号 后续扩展"`
	BanTime            uint64 `json:"banTime"            description:"封禁时间，0代表未封禁"`
	LastHeartbeatAt    int64  `json:"lastHeartbeatAt"    description:"用户心跳"`
	DeviceId           string `json:"deviceId"           description:"设备号"`
	AppName            string `json:"appName"            description:"app名称"`
	AppChannel         string `json:"appChannel"         description:"渠道"`
	AppVersion         string `json:"appVersion"         description:"app版本"`
	PhoneMode          string `json:"phoneMode"          description:"手机型号"`
	PhoneSystemVersion string `json:"phoneSystemVersion" description:"手机系统版本"`
	LastIp             string `json:"lastIp"             description:"最后一次登录ip"`
	System             int    `json:"system"             description:"1为安卓 2为ios"`
	UserLanguage       string `json:"userLanguage"       description:"用户语言"`
	CurrencyCode       int    `json:"currencyCode"       description:"国家数字码，比如印尼：360；马来：458  0为默认"`
	Remarks            string `json:"remarks"            description:"备注"`
	CreateTime         int64  `json:"createTime"         description:"创建时间"`
	UpdateTime         int64  `json:"updateTime"         description:"最后更新时间"`
}

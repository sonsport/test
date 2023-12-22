package model

import "github.com/gogf/gf/v2/frame/g"

type RegisterInput struct {
	Name         string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Avatar       string `json:"avatar"       description:"头像"`
	Password     string `json:"password"     description:"密码" v:"password"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterOutput struct {
	Id uint
}

type LoginInput struct {
	Name     string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password string `json:"password"     description:"密码" v:"password"`
}

type UpdatePasswordInput struct {
	Password     string `json:"password"     description:""`
	UserSalt     string `json:"userSalt,omitempty"     description:"加密盐 生成密码用"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type UpdatePasswordOutput struct {
	Id uint
}

type UserDetailInput struct {
	UserId uint64 `json:"userId"             description:"用户Id 自增"`
}

type UserDetailOutput struct {
	UserId          uint64 `json:"userId"             description:"用户Id 自增"`
	Username        string `json:"username"           description:"用户名"`
	AreaCode        int    `json:"areaCode"           description:"地区编码"`
	Nickname        string `json:"nickname"           description:"昵称"`
	Password        string `json:"password"           description:"密码  空字符串说明没有密码"`
	Gender          int    `json:"gender"             description:"性别 1男 2女 0未知"`
	Birthday        uint64 `json:"birthday"           description:"生日 时间戳"`
	Height          uint   `json:"height"             description:"身高 cm"`
	Education       uint   `json:"education"          description:"学历，枚举 小学 高中 大学 本科 扩展"`
	Industry        uint   `json:"industry"           description:"行业，枚举 学生、文化/艺术 扩展"`
	Intro           string `json:"intro"              description:"自我介绍 签名"`
	Portrait        string `json:"portrait"           description:"头像"`
	LiveState       uint   `json:"liveState"          description:"0未开播 1开播"`
	WhatsApp        string `json:"whatsApp"           description:"whatsapp号码"`
	LastHeartbeatAt int64  `json:"lastHeartbeatAt"    description:"用户心跳"`
	LastIp          string `json:"lastIp"             description:"最后一次登录ip"`
}

type UserInfoBase struct {
	g.Meta `orm:"table:user_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

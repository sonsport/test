package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta       `path:"/register" method:"post" tags:"前台用户" summary:"用户注册"`
	Name         string `json:"name"         description:"用户名" v:"required#用户名必填"`
	Password     string `json:"password"     description:"密码" v:"password"`
	Avatar       string `json:"avatar"       description:"头像"`
	UserSalt     string `json:"userSalt"     description:"加密盐 生成密码用"`
	Sex          int    `json:"sex"          description:"1男 2女"`
	Status       int    `json:"status"       description:"1正常 2拉黑冻结"`
	Sign         string `json:"sign"         description:"个性签名"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type RegisterRes struct {
	Id uint `json:"id"`
}

// LoginRes for gtoken
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserInfoBase
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"前台用户" summary:"当前登录用户信息"`
}

type UserInfoRes struct {
	UserInfoBase
}

// UserInfoBase 可以复用的，一定要抽取出来
type UserInfoBase struct {
	UserId    uint64 `json:"userId"             description:"用户Id 自增"`
	Username  string `json:"username"           description:"用户名"`
	AreaCode  int    `json:"areaCode"           description:"地区编码"`
	Nickname  string `json:"nickname"           description:"昵称"`
	Password  string `json:"password"           description:"密码  空字符串说明没有密码"`
	Gender    int    `json:"gender"             description:"性别 1男 2女 0未知"`
	Birthday  uint64 `json:"birthday"           description:"生日 时间戳"`
	Height    uint   `json:"height"             description:"身高 cm"`
	Education uint   `json:"education"          description:"学历，枚举 小学 高中 大学 本科 扩展"`
	Industry  uint   `json:"industry"           description:"行业，枚举 学生、文化/艺术 扩展"`
	Intro     string `json:"intro"              description:"自我介绍 签名"`
	Portrait  string `json:"portrait"           description:"头像"`
	LiveState uint   `json:"liveState"          description:"0未开播 1开播"`
	WhatsApp  string `json:"whatsApp"           description:"whatsapp号码"`
}

// UpdatePasswordReq 修改密码
type UpdatePasswordReq struct {
	g.Meta       `path:"/update/password" method:"post" tag:"前台用户" summary:"修改密码"`
	Password     string `json:"password"  v:"password"   description:""`
	UserSalt     string `json:"userSalt,omitempty"     description:"加密盐 生成密码用"`
	SecretAnswer string `json:"secretAnswer" description:"密保问题的答案"`
}

type UpdatePasswordRes struct {
	Id uint `json:"id"`
}

// UserDetailReq 获取用户详情
type UserDetailReq struct {
	g.Meta `path:"/user/detail" method:"get" tag:"api" summary:"获取用户详情"`
	UserId int64 `json:"userId"  v:"userId"   description:"userId"`
}

type UserDetailRes struct {
	UserInfoBase
}

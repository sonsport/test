package cmd

import (
	"context"
	"strconv"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/dao"
	"fuya-ark/internal/model"
	"fuya-ark/internal/model/entity"
	"fuya-ark/utility"
	"fuya-ark/utility/response"
)

// StartFrontendGToken 前台登录gtoken相关
func StartFrontendGToken() (gfFrontendToken *gtoken.GfToken, err error) {
	gfFrontendToken = &gtoken.GfToken{
		CacheMode:        consts.CacheModeRedis,
		ServerName:       consts.BackendServerName,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFuncFrontend,
		LoginAfterFunc:   loginAfterFuncFrontend,
		LogoutPath:       "/logout",
		AuthPaths:        g.SliceStr{"/*"},
		AuthExcludePaths: g.SliceStr{"/app/*", "/notify/*"}, // 不拦截路径
		AuthAfterFunc:    authAfterFuncFrontend,
		MultiLogin:       consts.FrontendMultiLogin,
	}
	//去掉全局校验，只用cmd中的路由组校验
	err = gfFrontendToken.Start()
	return
}

// for 前台项目
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	//验证账号密码是否正确
	userInfo := &entity.UserInfo{}
	err := dao.UserInfo.Ctx(r.GetCtx()).Where(dao.UserInfo.Columns().Username, name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	if userInfo == nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, userInfo.Username) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	// 唯一标识，扩展参数user data
	return consts.GTokenFrontendPrefix + gconv.String(userInfo.UserId), userInfo
}

// 自定义的登录之后的函数 for前台项目
func loginAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenFrontendPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := &model.LoginOutput{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
		}
		data.Username = userInfo.Username
		data.Portrait = userInfo.Portrait
		data.LiveState = userInfo.LiveState
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 登录鉴权中间件for前台
func authAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	//这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	client := 1
	clientOs := r.Header.Get("client-os")
	if len(clientOs) > 0 {
		client, _ = strconv.Atoi(clientOs)
	}

	r.SetCtxVar(consts.CtxUserId, userInfo.UserId)
	r.SetCtxVar(consts.CtxUserName, userInfo.Username)
	r.SetCtxVar(consts.CtxAppName, r.Header.Get("app-name"))                 //app名称
	r.SetCtxVar(consts.CtxAppChannel, r.Header.Get("app-channel"))           //app渠道
	r.SetCtxVar(consts.CtxClientLanguage, r.Header.Get("client-language"))   //用户语言
	r.SetCtxVar(consts.CtxDeviceId, r.Header.Get("device-id"))               //设备号
	r.SetCtxVar(consts.CtxAppVersion, r.Header.Get("app-v"))                 //app版本
	r.SetCtxVar(consts.CtxAppVersionCode, gconv.Int(r.Header.Get("app-vc"))) //手机型号
	r.SetCtxVar(consts.CtxClientMode, r.Header.Get("client-m"))              //手机型号
	r.SetCtxVar(consts.CtxClientSystemVersion, r.Header.Get("client-os-v"))  //手机系统版本
	r.SetCtxVar(consts.CtxClientIP, r.GetClientIp())                         //请求Ip
	r.SetCtxVar(consts.CtxClientOs, client)                                  //1为安卓 2为ios
	r.Middleware.Next()
}

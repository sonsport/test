package cmd

import (
	"context"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/controller"
	"fuya-ark/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			//订单超时未评价默认好评
			err = UserOrderDefaultComments(ctx)
			//if err != nil {
			//	panic(err)
			//}

			// 启动前台项目gtoken
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			//前台项目路由组
			s.Group("/ark-api", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.User.Register, //用户注册
				)
				//需要登录鉴权的路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					//需要登录鉴权的接口放到这里
					group.Bind(
						//controller.User.Info,           //当前登录用户的信息
						//controller.User.UpdatePassword, //当前用户修改密码
						controller.User.UserDetail, //当前登录用户的详细信息
					)
				})
			})
			s.SetPort(8000) //设置端口
			s.Run()
			return nil
		},
	}
)

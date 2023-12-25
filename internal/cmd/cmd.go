package cmd

import (
	"context"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

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
			s := g.Server(consts.ProjectName)
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
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				err := frontendToken.Middleware(ctx, group)
				if err != nil {
					return
				}
				//需要登录鉴权的路由组
				group.Group("/activity", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Activity,
					)
				})
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(
						//controller.User.UpdatePassword, //当前用户修改密码
						//controller.User.Info,           //当前登录用户的信息
						controller.User.UserDetail, //当前登录用户的详细信息
						controller.User.Register,   //用户注册
					)
				})
			})

			s.SetPort(8000) //设置端口
			s.Run()
			return nil
		},
	}
)

package cmd

import (
	"context"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	goredislib "github.com/redis/go-redis/v9"

	"fuya-ark/internal/consts"
	"fuya-ark/internal/router"
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
			initRedisLock() //初始化redis锁
			//前台项目路由组
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS, service.Middleware().Ctx, service.Middleware().ResponseHandler,
				)
				//s.Group("/", func(group *ghttp.RouterGroup) {
				//	//分享短链跳转
				//	group.GET("/shr/{code}", srv.Shr)
				//	//分享短链跳转
				//	group.GET("/apk/{code}", srv.Shr)
				//})
				err := frontendToken.Middleware(ctx, group)
				if err != nil {
					return
				}
				//需要登录鉴权的路由组
				router.InitActivityRouter(group)

			})

			s.SetPort(8000) //设置端口
			s.Run()
			return nil
		},
	}
)

func initRedisLock() {
	gRedisConfig, ok := gredis.GetConfig()
	if ok {
		return
	}
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     gRedisConfig.Address,
		Password: gRedisConfig.Pass,
		DB:       gRedisConfig.Db,
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	consts.RedisLock = redsync.New(pool)
}

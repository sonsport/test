package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"fuya-ark/internal/controller"
	"fuya-ark/internal/model"
)

func InitActivityRouter(s *ghttp.RouterGroup) {
	s.Group("/inviter", func(group *ghttp.RouterGroup) {
		//获取分享渠道及分享链接
		group.GET("/share/channels", BindRequest(model.ShareChannelsInput{}).Do(controller.Activity.ShareChannels))
		//分享登录绑定
		group.POST("/user/login/reward", BindRequest(model.NewUserLoginRewardInput{}).Do(controller.Activity.NewUserLoginReward))
		////获取拉新总览
		//group.GET("/get-inviter-info", BindRequest(srv.PullUserListReq{}).Do(srv.GetInviterInfo))
		////拉新详情
		//group.GET("/pull-user/records", BindRequest(base.PageSizeReq{}).Do(srv.PullUserRecords))
		////获取最新邀请用户奖励
		//group.GET("/get-invited-users", BindRequest(base.PageSizeReq{}).Do(srv.GetInvitedUsers))
		////分享回调
		//group.POST("/share-success", BindRequest(srv.ShareSuccessReq{}).Do(srv.ShareSuccess))
		////获取分享链接 获取分享
		//group.GET("/get-share-url", BindRequest(srv.GetShareUrlReq{}).Do(srv.GetShareUrl))
		////分享邀请成功  被邀请人打开调用
		//group.POST("/share-invitation-success", BindRequest(srv.ShareInvitationSuccessReq{}).Do(srv.ShareInvitationSuccess))
		////邀新成功排行榜
		//group.POST("/share-success-rank", BindRequest(srv.ShareSuccessRankReq{}).Do(srv.ShareSuccessRank))
		//
		////获取我的周统计直播数据
		//group.GET("/board/week", BindRequest(srv.AnchorBoardRequest{}).Do(srv.GetAnchorBoardWeekData))
		//group.GET("/room/list", BindRequest(base.EmptyReq{}).Do(srv.ListRoomActivity))
		//
		////公会直播活动奖励首播、四周
		//s.Group("/guild/reward/live", func(group *ghttp.RouterGroup) {
		//	//总计
		//	group.GET("/total", BindRequest(base.EmptyReq{}).Do(srv.GetGuildLiveRewardTotal))
		//	//结果
		//	group.GET("/list", BindRequest(srv.ActivityRewardForGuildRequest{}).Do(srv.GetGuildLiveRewardList))
		//	//公会推荐公会活动
		//	group.GET("/pull-guild/list", BindRequest(srv.ActivityRewardForPullGuildRequest{}).Do(srv.GetPullGuildLiveRewardList))
		//
		//})
		//// 钻石转盘抽奖活动
		//s.Group("/diamond", func(group *ghttp.RouterGroup) {
		//	// 主页列表
		//	group.GET("/turntable", BindRequest(base.EmptyReq{}).Do(srv.GetHomepageInfo))
		//	// 抽奖结果
		//	group.POST("/draw", BindRequest(base.EmptyReq{}).Do(srv.PrizeDraw))
		//	// 获取抽奖记录
		//	group.GET("/get-turntable-record", BindRequest(srv.GetTurntableRecordListReq{}).Do(srv.GetRecordList))
		//})
		//
		//// 道具抽奖活动
		//s.Group("/props", func(group *ghttp.RouterGroup) {
		//	// 获取抽奖配置
		//	group.GET("/get-prize-config", BindRequest(srv.GetConfigReq{}).Do(complexPrizeService.GetComplexPrizeConfig))
		//	// 抽奖
		//	group.POST("/draw", BindRequest(srv.GetConfigReq{}).Do(complexPrizeService.DoComplexPrize))
		//	// 获取抽奖记录
		//	group.GET("/get-prize-record", BindRequest(srv.GetConfigReq{}).Do(complexPrizeService.GetComplexPrizeRecord))
		//	// 获取大奖记录
		//	group.POST("/get-grand-prize-record", BindRequest(base.PageSizeReq{}).Do(complexPrizeService.GetGrandPrizeRecord))
		//})
		//// 主播开播排行活动
		//s.Group("/live", func(group *ghttp.RouterGroup) {
		//	// 直播时长排行
		//	group.GET("/live-time-rank", BindRequest(statistic.AnchorLiveTimeRanReq{}).Do(service.GetAnchorLiveTimeRank))
		//	group.GET("/insert-apply", BindRequest(srv.SaveAnchorLabel{}).Do(anchorLabelService.InsertApply))
		//	// 获取前一日排行
		//	group.GET("/get-rank-by-day", BindRequest(statistic.AnchorLiveTimeRanReq{}).Do(service.GetRankByDay))
		//})
		//
		//// 主播开播排行活动
		//s.Group("/recharge", func(group *ghttp.RouterGroup) {
		//	service := srv.NewRechargePointsRecordService()
		//	// 获取充值排行榜
		//	group.GET("/get-recharge-points-rank", BindRequest(v1.EmptyReq{}).Do(service.GetRechargePointsRank))
		//	// 获取排行榜记录
		//	group.GET("/get-recharge-points-record", BindRequest(v1.EmptyReq{}).Do(service.GetRechargePointsRecord))
		//	// 获取自己排名
		//	group.GET("/get-my-recharge-points", BindRequest(v1.EmptyReq{}).Do(service.GetMyRechargePoints))
		//})
		//
		//// 主播开播排行活动
		//s.Group("/order-referrer", func(group *ghttp.RouterGroup) {
		//	// 获取充值房间主播排行榜
		//	group.GET("/get-order-referrer-rank", BindRequest(statistic.OrderReferrerStatisticsReq{}).Do(service.GetAnchorOrderReferrerRank))
		//	// 获取自己的充值排行榜
		//	group.GET("/get-my-order-referrer-rank", BindRequest(base.EmptyReq{}).Do(service.GetMyAnchorOrderReferrerRank))
		//})
	})

	// 活动中心配置
	//s.Group("/inviter-center", func(group *ghttp.RouterGroup) {
	//	//活动配置
	//	group.GET("/config", BindRequest(base.EmptyReq{}).Do(srv.GetActivityCenterConfig))
	//	//活动配置列表
	//	group.GET("/config/list", BindRequest(base.PageSizeReq{}).Do(srv.GetActivityCenterConfigList))
	//	//更具类型获取活动配置
	//	group.GET("/config-by-type", BindRequest(srv.ConfigReq{}).Do(srv.GetActivityCenterByTypeConfig))
	//})

}

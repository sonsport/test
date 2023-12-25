package consts

const (
	OneMinute    = 60
	FiveMinute   = OneMinute * 5
	OneDay       = 60 * 60 * 24
	Board3Day    = OneDay * 3
	Board33Day   = OneDay * 33
	Board66Day   = OneDay * 66
	BoardForever = 0

	BoardUserSingleByDay   = "Board:User:Single:List:Day:"   // 当前主播-用户日贡献榜
	BoardUserSingleByWeek  = "Board:User:Single:List:Week:"  // 当前主播-用户周贡献榜
	BoardUserSingleByMonth = "Board:User:Single:List:Month:" // 当前主播-用户月贡献榜
	BoardUserSingleByAll   = "Board:User:Single:List:All:"   // 当前主播-用户总贡献榜

	BoardUserSingleTotalByDay   = "Board:User:Single:Total:Day:"   //当前主播当天累计收益
	BoardUserSingleTotalByWeek  = "Board:User:Single:Total:Week:"  //当前主播当周累计收益
	BoardUserSingleTotalByMonth = "Board:User:Single:Total:Month:" //当前主播当月累计收益
	BoardUserSingleTotalByAll   = "Board:User:Single:Total:All:"   //当前主播总共累计收益

	BoardUserPlatformByDay   = "Board:User:Platform:List:Day:"    // 平台-日 用户贡献榜
	BoardUserPlatformByWeek  = "Board:User:Platform:List:Week:"   // 平台-周 用户贡献榜
	BoardUserPlatformByMonth = "Board:User:Platform:List:Month:"  // 平台-月 用户贡献榜
	BoardUserPlatformByAll   = "Board:User:Platform:List:All:All" // 平台-总 用户贡献榜

	BoardUserPlatformTotalByDay   = "Board:User:Platform:Total:Day:"    //平台当日累计收益
	BoardUserPlatformTotalByWeek  = "Board:User:Platform:Total:Week:"   //平台当周累计收益
	BoardUserPlatformTotalByMonth = "Board:User:Platform:Total:Month:"  //平台当月累计收益
	BoardUserPlatformTotalByAll   = "Board:User:Platform:Total:All:All" //平台总共累计收益

	BoardAuthorPlatformByDay   = "Board:Author:Platform:List:Day:"    // 平台-日 主播收益榜
	BoardAuthorPlatformByWeek  = "Board:Author:Platform:List:Week:"   // 平台-周 主播收益榜
	BoardAuthorPlatformByMonth = "Board:Author:Platform:List:Month:"  // 平台-月 主播收益榜
	BoardAuthorPlatformByAll   = "Board:Author:Platform:List:All:All" // 平台-总 主播收益榜

	BoardAuthorPlatformTotalByDay   = "Board:Author:Platform:Total:Day:"    //当天累计消耗
	BoardAuthorPlatformTotalByWeek  = "Board:Author:Platform:Total:Week:"   //当周累计消耗
	BoardAuthorPlatformTotalByMonth = "Board:Author:Platform:Total:Month:"  //当月累计消耗
	BoardAuthorPlatformTotalByAll   = "Board:Author:Platform:Total:All:All" //总共累计消耗

	BoardAuthorDiamondPlatformByDay   = "Board:Author:Diamond:Platform:List:Day:"    // 平台-日 主播钻石流水榜
	BoardAuthorDiamondPlatformByWeek  = "Board:Author:Diamond:Platform:List:Week:"   // 平台-周 主播钻石流水榜
	BoardAuthorDiamondPlatformByMonth = "Board:Author:Diamond:Platform:List:Month:"  // 平台-月 主播钻石流水榜
	BoardAuthorDiamondPlatformByAll   = "Board:Author:Diamond:Platform:List:All:All" // 平台-总 主播钻石流水榜

	BoardAuthorDiamondPlatformTotalByDay   = "Board:Author:Diamond:Platform:Total:Day:"    //当天主播累计流水
	BoardAuthorDiamondPlatformTotalByWeek  = "Board:Author:Diamond:Platform:Total:Week:"   //当周主播累计流水
	BoardAuthorDiamondPlatformTotalByMonth = "Board:Author:Diamond:Platform:Total:Month:"  //当月主播累计流水
	BoardAuthorDiamondPlatformTotalByAll   = "Board:Author:Diamond:Platform:Total:All:All" //总共主播累计流水

	RankUserByRoom   = "Board:User:Room:All:"   // 单个房间- 用户贡献榜
	RankAuthorByRoom = "Board:Author:Room:All:" // 单个房间主播累计收益

	//LastBillBoardCacheKey 上一个榜单key用来判断发放勋章 :%v boardType
	LastBillBoardCacheKey = "Board:Last:CacheKey:"

	//AnchorSettlementIncomeRankZSetKey 主播结算收益排行
	AnchorSettlementIncomeRankZSetKey = "Anchor:Settlement:IncomeRank:ZSetKey"
	//AnchorSettlementIncomeUserRankIncomeKey 主播结算收益用户排名收益
	AnchorSettlementIncomeUserRankIncomeKey = "Anchor:Settlement:IncomeRank:User:Income"
	//BoardGameUserDiamondPlatform 游戏钻石流水排行
	BoardGameUserDiamondPlatform = "Board:Game:User:Diamond:Platform" // 平台 游戏钻石流水榜

)

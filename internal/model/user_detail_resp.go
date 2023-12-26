package model

// DetailResp 用户详情返回包装对象
type DetailResp struct {
	User           *InfoVo        `json:"user"`
	Balance        *BalanceVo     `json:"balances"`
	UserDressUps   *DressUpVo     `json:"userDressUps"`
	GuildAnchor    *GuildAnchorVo `json:"guildAnchor"`
	ServerWhatsApp string         `json:"serverWhatsApp"`
}

type InfoVo struct {
	UserId               int64                       `json:"userId"`               //用户Id 自增
	SuperId              int64                       `json:"superId"`              //靓号 默认为0
	SuperTime            int64                       `json:"superTime"`            //靓号过期时间
	Nickname             string                      `json:"nickname"`             //昵称
	Gender               int64                       `json:"gender"`               //性别 1男 2女 0未知
	Birthday             int64                       `json:"birthday"`             //生日 时间戳
	Height               int64                       `json:"height"`               //身高 cm
	Education            int64                       `json:"education"`            //学历，枚举 小学 高中 大学 本科 扩展
	Industry             int64                       `json:"industry"`             //行业，枚举 学生、文化/艺术 扩展
	Intro                string                      `json:"intro"`                //自我介绍 签名
	Portrait             string                      `json:"portrait"`             //头像
	SmallImage           string                      `json:"smallImage"`           //缩略图
	Role                 int                         `json:"role"`                 //0普通用户,1认证主播,2取消认证主播,3房间管理员,4超级管理员
	LiveState            int64                       `json:"liveState"`            //0未开播 1开播
	WhatsApp             string                      `json:"whatsApp"`             //whatsapp号码
	UserLevel            int                         `json:"userLevel"`            //用户等级，根据用户打赏给主播的钻石来定，达到一定数额就升级
	AnchorLevel          int                         `json:"anchorLevel"`          //主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	LastHeartbeatAt      int64                       `json:"lastHeartbeatAt"`      //用户心跳
	FollowCount          int                         `json:"followCount"`          //我关注别人的数量
	FollowedCount        int                         `json:"followedCount"`        //我的粉丝数量（别人关注我）
	IsNewUser            uint8                       `json:"isNewUser"`            //0为老用户 1为新用户
	CurrencyCode         int64                       `json:"currencyCode"`         //国家数字码，比如印尼：360；马来：458  0为默认
	Remarks              string                      `json:"remarks"`              //备注
	AnchorHighQuality    int64                       `json:"anchorHighQuality"`    //是否高质量主播 1高质量
	FollowStatus         int                         `json:"followStatus"`         //关注状态
	TotalDiamonds        int64                       `json:"totalDiamonds"`        //送出钻石
	TotalIncome          int64                       `json:"totalIncome"`          //收益
	FeedCount            int                         `json:"feedCount"`            //动态条数
	RoomId               string                      `json:"roomId"`               //直播间ID全局唯一
	StreamId             string                      `json:"streamId"`             //直播流id
	PrivateStatus        int                         `json:"private_status"`       //0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
	CreateTime           int64                       `json:"createTime"`           //创建时间
	UpdateTime           int64                       `json:"updateTime"`           //创建时间
	StarLevel            int                         `json:"starLevel"`            //主播星级
	BirthdayStart        string                      `json:"birthdayStart"`        //非审核 fail 审核success
	BirthdayTime         bool                        `json:"birthdayTime"`         //非审核 false true
	GenderState          int                         `json:"genderState"`          //非审核 0 100
	IntroDetails         string                      `json:"introDetails"`         //非审核  nothing     details
	AgeTime              int                         `json:"ageTime"`              //非审核 0 100
	UserDressUps         *DressUpVo                  `json:"userDressUps"`         //10
	BeautifulNumDressUps *ShopGoodsBubbleDressUpItem `json:"beautifulNumDressUps"` //靓号框
	LoginType            int64                       `json:"loginType"`            //1 gg 2 fb 3 apple
	ShowWhatsApp         bool                        `json:"showWhatsApp"`         //显示whatsapp入口
}

type BalanceVo struct {
	UserId              int64 `json:"userId"`              // 用户Id
	TotalFee            int64 `json:"totalFee"`            // 总花费、只有付费购买才累计到这里 印尼盾 单位：分
	TotalPayoutDiamonds int64 `json:"totalPayoutDiamonds"` // 总流水，凡是用户花费的钻石都记录到这里 累计
	TotalPayDiamonds    int64 `json:"totalPayDiamonds"`    // 总购买钻石 只累计购买的
	RemainDiamonds      int64 `json:"remainDiamonds"`      // 剩余钻石数
}

type GuildAnchorVo struct {
	ID        int   `json:"id"`         // 自增ID
	Uid       int64 `json:"uid"`        // 用户uid
	GuildID   int64 `json:"guild_id"`   // 公会ID
	GuildRole int8  `json:"guild_role"` // 公会角色 1 工会长；2 成员
	Status    int8  `json:"status"`     // 状态 1 正常；2 失效
	StarLevel int   `json:"star_level"` // 主播星级
	JoinTime  int64 `json:"join_time"`  // 加入时间
}

// DressUpVo 装扮信息
type DressUpVo struct {
	BubbleBox      *ShopGoodsBubbleDressUpItem `json:"bubble_box"`      // 气泡框
	GoldenLight    *DressUpItem                `json:"golden_light"`    // 金光
	AvatarFrame    *DressUpItem                `json:"avatar_frame"`    // 头像框
	Ride           *DressUpItem                `json:"ride"`            // 座驾
	Badge          *DressUpItem                `json:"badge"`           // 徽章
	ActivityBadges []*DressUpItem              `json:"activity_badges"` // 活动徽章
	HonoraryTitles []*HonoraryTitles           `json:"honorary_titles"` // 荣誉称号
}

// BeautifulNumVo 靓号信息
type BeautifulNumVo struct {
	Id             int64  `json:"id"`              //靓号id
	SuperId        int64  `json:"super_id"`        //靓号
	IconUrl        string `json:"icon_url"`        //icon地址
	Experience     int64  `json:"experience"`      //经验值
	ExpirationTime int64  `json:"expiration_time"` //过期时间
	ReclaimTime    int64  `json:"reclaim_time"`    //回收时间
}

// RoomViewerResp 直播间观众列表
type RoomViewerResp struct {
	RoomViewers []*RoomViewer `json:"room_viewer_list"`
	Total       int           `json:"total"` //观众数
}

// RoomViewer 直播间观众列表
type RoomViewer struct {
	UserId               int64                       `json:"userId"`               //用户Id
	SuperId              int64                       `json:"superId"`              //靓号
	Nickname             string                      `json:"nickname"`             //昵称
	PrivilegeType        int                         `json:"privilege_type"`       //特权类型
	Gender               int                         `json:"gender"`               //性别 1男 2女 0未知
	Portrait             string                      `json:"portrait"`             //头像
	UserLevel            int                         `json:"userLevel"`            //用户等级，根据用户打赏给主播的钻石来定，达到一定数额就升级
	AnchorLevel          int                         `json:"anchorLevel"`          //主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	UserDressUps         *DressUpVo                  `json:"userDressUps"`         //用户装扮
	BeautifulNumDressUps *ShopGoodsBubbleDressUpItem `json:"beautifulNumDressUps"` // 靓号框
}

// LiveAwardRecord 直播奖励记录
type LiveAwardRecord struct {
	Diamonds int64 `json:"diamonds"`  // 钻石
	UserId   int64 `json:"user_id"`   // 用户Id
	LinkId   int64 `json:"link_id"`   // 关联id
	LiveTime int64 `json:"live_time"` // 直播时长
	EndTime  int64 `json:"end_time"`  //直播开始时间

}

type ServerInfoResp struct {
	WhatsApp string `json:"whatsApp"`
}

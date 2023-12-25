package model

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

// LoginOutput for gtoken
type LoginOutput struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserDetailOutput
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

type UserInfo struct {
	UserId         int64                       `json:"userId"`         // 用户Id 自增
	Nickname       string                      `json:"nickname"`       //
	Portrait       string                      `json:"portrait"`       //
	PrivilegeType  int                         `json:"privilege_type"` //特权类型
	UserLevel      int                         `json:"userLevel"`
	AuthorLevel    int                         `json:"authorLevel"`
	Gender         int                         `json:"gender"`
	LiveState      int64                       `json:"liveState"`       //0未开播、1开播
	StarLevel      int                         `json:"starLevel"`       //主播星级
	RoomId         string                      `json:"roomId"`          //开播房间id
	StreamId       string                      `json:"streamId"`        //直播流id
	BubbleBox      *ShopGoodsBubbleDressUpItem `json:"bubble_box"`      //装扮信息 // 气泡框
	GoldenLight    *DressUpItem                `json:"golden_light"`    //金光
	AvatarFrame    *DressUpItem                `json:"avatar_frame"`    //头像框
	Ride           *DressUpItem                `json:"ride"`            //座驾
	Badge          *DressUpItem                `json:"badge"`           //徽章
	ActivityBadges []*DressUpItem              `json:"activity_badges"` //活动徽章
	HomeCover      *DressUpItem                `json:"home_cover"`      //首页封面
}

type AgentRechargeRes struct {
	Id               int64    `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId           int64    `gorm:"column:user_id" db:"source_user_id" json:"user_id" form:"user_id"`                          //用户Id
	TargetUserId     int64    `gorm:"column:target_user_id" db:"target_user_id" json:"target_user_id" form:"target_user_id"`     //对方用户id
	RecordType       int8     `gorm:"column:record_type" db:"record_type" json:"record_type" form:"record_type"`                 //消费类型 1支出 2收入
	OnePrice         float64  `gorm:"column:one_price" db:"one_price" json:"one_price" form:"one_price"`                         //售出单价
	CostPrice        float64  `gorm:"column:cost_price" db:"cost_price" json:"cost_price" form:"cost_price"`                     //成本单价
	Diamonds         int64    `gorm:"column:diamonds" db:"diamonds" json:"diamonds" form:"diamonds"`                             //钻石数额
	BeforeChange     int64    `gorm:"column:before_change" db:"before_change" json:"before_change" form:"before_change"`         //变更之前余额
	AfterChange      int64    `gorm:"column:after_change" db:"after_change" json:"after_change" form:"after_change"`             //变更之后余额
	LinkId           int64    `gorm:"column:link_id" db:"link_id" json:"link_id" form:"link_id"`                                 //引起余额变更的id
	CertificateUrl   string   `gorm:"column:certificate_url" db:"certificate_url" json:"certificate_url" form:"certificate_url"` //凭证url
	PaymentId        int64    `gorm:"column:payment_id" db:"payment_id" json:"payment_id" form:"payment_id"`                     //支付id
	PaymentName      string   `gorm:"column:payment_name" db:"payment_name" json:"payment_name" form:"payment_name"`             //支付名称
	CreateTime       int64    `db:"create_time" gorm:"column:create_time" json:"create_time"`                                    // 创建时间
	UpdateTime       int64    `db:"update_time" gorm:"column:update_time" json:"update_time"`                                    // 最后更新时间
	TotalFee         int64    `json:"total_fee" gorm:"column:total_fee"`                                                         //总充值 印尼盾 单位：分
	AgentsPropsNames []string `json:"agents_props_names"`                                                                        //代理道具名称列表
}

type AnchorWeekRankRes struct {
	UserId           int64          `json:"userId"`           //用户Id
	Nickname         string         `json:"nickname"`         //昵称
	Portrait         string         `json:"portrait"`         //头像
	LiveState        int64          `json:"liveState"`        //0未开播 1开播
	UserLevel        int            `json:"userLevel"`        //用户等级，根据用户打赏给主播的钻石来定，达到一定数额就升级
	AnchorLevel      int            `json:"anchorLevel"`      //主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	BalancesIncome   int64          `json:"balancesIncome"`   //收益
	WeekWaterDiamond int64          `json:"weekWaterDiamond"` //流水
	RoomId           string         `json:"roomId"`           //直播间ID全局唯一
	StreamId         string         `json:"streamId"`         //直播流id
	StarLevel        int            `json:"starLevel"`        //主播星级
	IsStandard       int            `json:"isStandard"`       //是否达标 达标
	AvatarFrame      *DressUpItem   `json:"avatarFrame"`      //头像框
	Badge            *DressUpItem   `json:"badge"`            //徽章
	ActivityBadges   []*DressUpItem `json:"activityBadges"`   //活动徽章
}

type OrderRankRes struct {
	UserId   int64 `json:"user_id"`  //用户Id
	Diamonds int64 `json:"diamonds"` //充值钻石
}

type OtherAchievement struct {
	// 勋章图标url
	Medals []string `json:"medals"`
	// 贡献榜头像url
	BillboardPortraits []string `json:"billboard_portraits"`
	// 公会成员头像url
	MemberPortraits []string `json:"member_portraits"`
	// pk段位
	// 天梯赛是否开启
	LadderTournamentStates bool `json:"ladder_tournament_states"`
	BigRank                int  `json:"big_rank"`
	DetailRank             int  `json:"detail_rank"`
	// 装扮
	DressUp []*DressUpItem `json:"dress_up,omitempty"`
}

// UpdateUserReq 跟新用户请求
type UpdateUserReq struct {
	// 昵称
	Nickname string `json:"nickname"`
	// 密码  空字符串说明没有密码
	Password string `json:"password"`
	// 性别 1男 2女 0未知
	Gender int `json:"gender"`
	// 生日 时间戳
	Birthday int64 `json:"birthday"`
	// 身高 cm
	Height int `json:"height"`
	// 学历，枚举 0小学 1高中 2大学 3本科 4扩展
	Education int `json:"education"`
	// 行业，枚举 1学生、2文化/艺术 3扩展
	Industry int `json:"industry"`
	// 自我介绍 签名
	Intro string `json:"intro"`
	// 头像
	Portrait string `json:"portrait"`
	// whatsapp号码
	WhatsApp string `json:"whatsApp"`
}

type HeartbeatReq struct {
}

type GetUserReq struct {
	UserId int64 `json:"userId"`
}

type SearchUserReq struct {
	SearchKey string `json:"searchKey" v:"required"`
}

type CheckUserType struct {
	CheckType int `json:"check_type"` // 用户Id

}

type UpdateMessageStateReq struct {
	Id    int64 `json:"id"`
	State int   `json:"state"`
}

type UploadAppMaiDianReq struct {
	Id   string `json:"log_type"`
	Data string `json:"data"`
}

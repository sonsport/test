package model

import (
	"time"
)

// 分享来源方式
var PullNewUserSource = "pull_new_user" //分享拉新用户

var ShareSourceMap = map[string]int{
	PullNewUserSource: 1,
}

const (
	//奖励方式
	//AwardInviteUser   = 1 //新用户注册
	//AwardUserRecharge = 2 //新用户充值
	//RoomShare         = 3 // 房间内分享的

	InviteStatPass    = 1  // 邀请状态正常
	InviteStatInvalid = -1 // 邀请状态失效
)

type GetPullNewUserShareUrlParams struct {
	UserId  int64                  `json:"userId"`       //分享人uid
	ShareId int64                  `json:"share_id"`     //分享ID
	Source  string                 `json:"source"`       //分享来源 获取短链必须参数
	Custom  map[string]interface{} `json:"custom_query"` //自定义拼接参数
}

type ShareParam struct {
	ShareId string `json:"share_id"`
}

type ShareOwnerInfoVo struct {
	Id              int64 `json:"id" gorm:"column:id"`
	UserId          int64 `json:"user_id" gorm:"column:user_id"`
	Stat            int   `json:"stat" gorm:"column:stat"`
	InviteUserCount int   `json:"invite_user_count" gorm:"column:invite_user_count"`
	OwnerReward     int   `json:"owner_reward" gorm:"column:owner_reward"`       // 邀请人奖励
	RechargeReward  int   `json:"recharge_reward" gorm:"column:recharge_reward"` // 邀请人充值奖励
	NewUserReward   int   `json:"new_user_reward" gorm:"column:new_user_reward"` // 被邀请人奖励
	//Operator        string `json:"operator" gorm:"column:operator"`
	PrepaidCount int `json:"prepaid_count"`
}

type ShareNewUserInfoVo struct {
	Id                  int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId              int64  `gorm:"column:user_id;NOT NULL" json:"user_id"`
	ShareId             int64  `gorm:"column:share_id;NOT NULL" json:"share_id"`                  // 分享ID
	OwnerUserId         int64  `gorm:"column:owner_user_id;NOT NULL" json:"owner_user_id"`        // 分享者用户id
	OwnerUserReward     int64  `gorm:"column:owner_user_reward" json:"owner_user_reward"`         // 分享者获得的钻石数
	NewUserReward       int64  `gorm:"column:new_user_reward" json:"new_user_reward"`             // 新用户获得的钻石数
	OwnerRechargeReward int64  `gorm:"column:owner_recharge_reward" json:"owner_recharge_reward"` // 邀请者获得新用户的充值奖励
	RechargeNum         int64  `gorm:"column:recharge_num;NOT NULL" json:"recharge_num"`          // 充值的虚拟币数量
	WatchLiveTime       int64  `gorm:"column:watch_live_time;NOT NULL" json:"watch_live_time"`    // 观看直播时长
	RegisterTime        int64  `gorm:"column:register_time" json:"register_time"`                 // 注册时间
	IP                  string `gorm:"column:ip" json:"ip"`                                       // 登陆地IP
	Smid                string `gorm:"column:smid" json:"smid"`                                   // 设备ID
	DeviceName          string `gorm:"column:device_name" json:"device_name"`                     // 设备型号
	Status              int64  `gorm:"column:status" json:"status"`
	Nickname            string `db:"nickname" form:"nickname" json:"nickname"`
}

type ShareNewUserInfoV1 struct {
	UserId      int64  `json:"user_id"`
	Status      int    `json:"status"`
	Nickname    string `json:"nickname"`
	RechargeNum int64  `json:"recharge_num"`
	IsOldUser   int    `json:"is_old_user"` //1:老用户；0:新用户
}

// GuildSettlementTotalResp 结算总计
type GuildSettlementTotalResp struct {
	//PayCount       int   `json:"pay_count"`        //已经结算的人数
	GuildHostNum      int   `json:"guild_host_num"`       //公会下主播个数（标星）
	GuildPullGuildNum int   `json:"guild_pull_guild_num"` //公会推荐公会个数
	UnPayTotal        int64 `json:"unpay_total"`          //未支付总和
	PayTotal          int64 `json:"pay_total"`            //已支付总和  pay_count
	UnPayHostCount    int   `json:"unpay_host_count"`     //未结算的人数(GuildHostNum-PayCount)
	MaxTotal          int64 `json:"max_total"`            //最多可得到的奖励(GuildHostNum*50000)
	TargetPay         int64 `json:"target_pay"`           //未支付总和 等同于UnPayTotal
}

type GuildSettlementItemResp struct {
	UserId      int64  `json:"user_id"`      // 用户uid
	Nickname    string `json:"nickname"`     // 昵称
	LiveTime    int64  `json:"live_time"`    //有效直播时长
	RewardMoney int64  `json:"reward_money"` //奖励
	RewardState int8   `json:"reward_state"` //奖励状态 0未发放 1已发放
	EffectDay   int    `json:"effect_day"`   //有效直播天
	RewardDate  int64  `json:"reward_date"`  //发放日期
}

type PullGuildSettlementItemResp struct {
	UserId      int64  `json:"user_id"`      // 用户uid
	Nickname    string `json:"nickname"`     // 昵称
	LiveTime    int64  `json:"live_time"`    //有效直播时长 单位：秒
	RewardMoney int64  `json:"reward_money"` //奖励
	RewardState int8   `json:"reward_state"` //奖励状态 0未发放 1已发放
	RewardDate  int64  `json:"reward_date"`  //发放日期
}

type DiamondTurntableRecordListResp struct {
	UserId     int64  `json:"user_id"`     // 用户uid
	Nickname   string `json:"nickname"`    // 昵称
	Name       string `json:"name"`        // 礼物名称
	Num        int64  `json:"num"`         // 数量
	CreateTime int64  `json:"create_time"` // 中奖时间
}

type DiamondTurntableHomepageInfo struct {
	Id   int64  `json:"id"`   // 礼物id
	Name string `json:"name"` // 礼物名称
	Num  int64  `json:"num"`  // 数量
	Icon string `json:"icon"` // 礼物图标
}

type DiamondTurntableHomepageResp struct {
	Count     int                            `json:"count"`      // 抽奖次数
	WatchTime int                            `json:"watch_time"` // 观看时间 单位分钟
	GiftList  []DiamondTurntableHomepageInfo `json:"gift_list"`  // 配置礼物列表
}

type DiamondTurntableDrawResp struct {
	Id int64 `json:"id"` // 抽中奖品在配置表的id
}

type DiamondWatchInfo struct {
	RoomId         string `json:"room_id"`          //房间id
	HeartBeatCount int64  `json:"heart_beat_count"` //当前房间心跳次数
	IsReward       bool   `json:"isreward"`         //是否发放奖励 false没发过 ture发过
}

// ConfigResp  活动中心配置
type ConfigResp struct {
	StartUpPage  *ConfigBase   `json:"startup_page"`   //启动页
	HomePage     *HomePage     `json:"home_page"`      //首页
	RoomLivePage *RoomLivePage `json:"room_live_page"` //直播间
}

// HomePage 首页
type HomePage struct {
	PopUpList  []*ConfigBase `json:"pop_up_list"` //弹窗
	FloatList  []*ConfigBase `json:"float_list"`  //浮窗
	BannerList []*ConfigBase `json:"banner_list"` //banner
}

// ConfigBase 活动基础配置
type ConfigBase struct {
	AcId        int    `json:"ac_id"`        //活动Id activity_config 自增id
	Title       string `json:"title"`        //活动标题
	ResUrl      string `json:"res_url"`      //资源地址
	Duration    int    `json:"duration"`     //持续时间 单位秒 >0才倒计时
	LinkAddress string `json:"link_address"` //链接地址
	ShowStyle   int8   `json:"show_style"`   //0无配置 1全屏 2半屏
}

// RoomLivePage 直播间活动
type RoomLivePage struct {
	PopUpList   []*ConfigBase `json:"pop_up_list"`   //弹窗
	FloatList   []*ConfigBase `json:"float_list"`    //浮窗
	Bottom      *ConfigBase   `json:"bottom"`        //底部
	GiftTopList []*ConfigBase `json:"gift_top_list"` //直播间礼物顶部
}

// ConfigPlusResp 活动中心列表
type ConfigPlusResp struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`        //活动标题
	ResUrl      string `json:"res_url"`      //资源图 活动横图
	LinkAddress string `json:"link_address"` //活动地址
	StartTime   int64  `json:"start_time"`   //活动开始时间
	EndTime     int64  `json:"end_time"`     //活动结束时间
	State       int    `json:"state"`        //状态 0下架 1上架
}

type ComplexPrizeConfigRes struct {
	Id           int64  `json:"id"`
	ActivityType int    `json:"activity_type"` //抽奖活动类型 1 圣诞 2 元旦
	Icon         string `json:"icon"`          //图标
	DrawType     int    `json:"draw_type"`     //抽奖类型 1装扮 2 钻石 3 优惠券 4 实物 5 礼物
	PrizeName    string `json:"prize_name"`    //奖品名称
	Num          int    `json:"num"`           //数量
	PrizeCount   int    `json:"prize_count"`   //奖池期数
}

type ComplexPrizeRecordRes struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`       //用户id
	NickName     string `json:"nickname"`      //用户昵称
	CId          int64  `json:"c_id"`          //抽奖配置id
	ActivityType int    `json:"activity_type"` //抽奖活动类型 1 圣诞 2 元旦
	Icon         string `json:"icon"`          //图标
	DrawType     int    `json:"draw_type"`     //抽奖类型 1装扮 2 钻石 3 优惠券 4 实物 5 礼物
	LinkId       string `json:"link_id"`       //关联商品id  1,2  逗号分割
	PrizeName    string `json:"prize_name"`    //奖品名称
	PrizeCount   int    `json:"prize_count"`   //奖池期数
	CreateTime   int64  `json:"create_time"`   //创建时间
}

// AnchorLiveTimeRankResp 每日主播直播时间排行
type AnchorLiveTimeRankResp struct {
	Id                int64     `json:"id"`                  //ID
	UserId            int64     `json:"user_id"`             //用户ID
	Nickname          string    `json:"nickname"`            //昵称
	RankDateTime      time.Time `json:"rank_date_time"`      //排行日期
	EffectLiveSeconds int64     `json:"effect_live_seconds"` //有效直播时长
}

type YesterAnchorLiveTimeRankResp struct {
	UserId            int64 `json:"userId"`       //用户ID
	EffectLiveSeconds int64 `json:"live_seconds"` //有效直播时长
}

type RechargeUserRes struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Portrait string `json:"portrait"`
	Points   int    `json:"points"`
	Level    int    `json:"level"`
	Rank     int    `json:"rank"`
}

// LiveTimeRankResp 每日主播直播时间排行
type LiveTimeRankResp struct {
	Sort              int       `json:"sort"`                //排名
	GuildUserId       int64     `json:"guild_user_id"`       //工会章id
	UserId            int64     `json:"user_id"`             //用户ID
	Nickname          string    `json:"nickname"`            //昵称
	RankDateTime      time.Time `json:"rank_date_time"`      //排行日期
	EffectLiveSeconds int64     `json:"effect_live_seconds"` //有效直播时长
	Label             string    `json:"label"`               //标签
	StarLevel         int       `json:"star_level"`          //标签
	Reward            int       `json:"reward"`              //奖励
}

type OrderDiamondRankRes struct {
	UserId      int64    `json:"user_id"`      //用户id
	NickName    string   `json:"nickname"`     //昵称
	Portrait    string   `json:"portrait"`     //头像
	Diamonds    int      `json:"diamonds"`     //钻石
	UserLevel   int      `json:"user_level"`   //用户等级
	AnchorLevel int      `json:"anchor_level"` //主播等级
	Portraits   []string `json:"portraits"`
}

type ShareSuccessRankRes struct {
	UserId   int64  `json:"user_id"`  //用户id
	NickName string `json:"nickname"` //昵称
	Portrait string `json:"portrait"` //头像
	Diamonds int    `json:"diamonds"` //钻石
}

type ExportAnchorLiveRankRsp struct {
	Rank         int     `json:"rank"`           //rank
	UserId       int64   `json:"userId"`         //用户ID
	RankDateTime string  `json:"rank_date_time"` //时间
	Label        string  `json:"label"`          //标签
	StarLevel    int     `json:"star_level"`     //星级
	LiveTime     float64 `json:"live_time"`      //直播时长小时
	LiveSeconds  float64 `json:"live_seconds"`   //直播时长小时
	GuildId      int64   `json:"guild_id"`       //公会id
	BankName     string  `json:"bank_name"`      //银行名
	BankAccounts string  `json:"bank_accounts"`  //银行账号
	BankPerson   string  `json:"bank_person"`    //持卡人姓名
}

type OrderDiamondRankResArr []*OrderDiamondRankRes

func (a OrderDiamondRankResArr) Len() int {
	return len(a)
}

func (a OrderDiamondRankResArr) Less(i, j int) bool {
	return a[i].Diamonds > a[j].Diamonds
}

func (a OrderDiamondRankResArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

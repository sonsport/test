package model

import "time"

type UserWeekDataResp struct {
	//UserLevel    int            `json:"userLevel"`    // 用户等级，根据用户打赏给主播的钻石来定，达到一定数额就升级
	UserId               int64   `json:"user_id" gorm:"column:user_id"`                   // 用户uid
	Nickname             string  `json:"nickname"`                                        // 昵称
	Portrait             string  `json:"portrait"`                                        // 头像
	WeekWaterDiamond     int64   `json:"week_water_diamond"`                              // 本周流水
	LastWeekWaterDiamond int64   `json:"last_week_water_diamond"`                         // 上周流水
	AnchorLevel          int     `json:"anchorLevel"`                                     // 主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	GuildId              int64   `json:"guild_id" gorm:"column:guild_id"`                 // 公会ID
	GuildName            string  `json:"guild_name" gorm:"column:guild_name"`             // 公会名
	GuildCutRate         string  `json:"guild_cut_rate" gorm:"column:guild_cut_rate"`     // 公会抽成
	LastWeekIncome       float64 `json:"last_week_income" gorm:"column:last_week_income"` // 上周收益
	ThisWeekIncome       float64 `json:"this_week_income" gorm:"column:this_week_income"` //本周钻收益
}

type UserBoardWeekDataResp struct {
	BoardList []UserBoard `json:"board_list" gorm:"column:board_list"` // 排行榜
	MyBoard   *MyBoard    `json:"my_board" gorm:"column:my_board"`     // 我的排行榜
}
type UserBoard struct {
	UserId      int64   `json:"user_id" gorm:"column:user_id"`           // 用户uid
	Nickname    string  `json:"nickname"`                                // 昵称
	Portrait    string  `json:"portrait"`                                // 头像
	AnchorLevel int     `json:"anchorLevel"`                             // 主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	TotalIncome float64 `json:"total_income" gorm:"column:total_income"` // 收益
	Rank        int     `json:"rank" gorm:"column:rank"`                 //排行
}

type MyBoard struct {
	UserId      int64   `json:"user_id" gorm:"column:user_id"`             // 用户uid
	Nickname    string  `json:"nickname"`                                  // 昵称
	Portrait    string  `json:"portrait"`                                  // 头像
	AnchorLevel int     `json:"anchorLevel"`                               // 主播等级，用户打赏给主播的钻石来定，达到一定数额就升级
	TotalIncome float64 `json:"total_income" gorm:"column:total_income"`   // 收益
	Rank        int     `json:"rank" gorm:"column:rank"`                   //排行 》0已上榜 0未上榜
	IsUp        int     `json:"is_up" gorm:"column:is_up"`                 //是否达标 1达标 0未达标
	LiveTime    int64   `json:"live_time" gorm:"column:live_time"`         //有效直播时长
	ToLiveTime  int64   `json:"to_live_time" gorm:"column:to_live_time"`   //需要达到的有效直播时长
	EffectDay   int     `json:"effect_day" gorm:"column:effect_day"`       //有效天
	ToEffectDay int     `json:"to_effect_day" gorm:"column:to_effect_day"` //需要达到的有效天
}

type GuildTotalDataResp struct {
	TotalIncome      float64 `json:"total_income" gorm:"column:total_income"`
	HostTotalIncome  float64 `json:"host_total_income" gorm:"column:host_total_income"`   //主播总收入 印尼盾 分
	GuildTotalIncome float64 `json:"guild_total_income" gorm:"column:guild_total_income"` //公会总收入 印尼盾 分
	WaterDiamond     int64   `json:"water_diamond"`
	Total            float64 `json:"total" gorm:"column:total"` //总计
}

type AnchorLiveResp struct {
	StatisticsTime  time.Time `json:"statistics_time" gorm:"column:statistics_time"`
	LiveDate        string    `json:"live_date" `
	DayWaterDiamond int64     `json:"day_water_diamond"`
	LiveTime        int64     `json:"live_time" gorm:"column:live_time"`
	AllIncome       float64   `json:"all_income" gorm:"column:all_income"`
}

type GuildAnchorLiveResp struct {
	UserId          int64   `json:"user_id" gorm:"column:user_id"`   // 用户uid
	Nickname        string  `json:"nickname" gorm:"column:nickname"` // 用户uid
	EffectLiveDay   int64   `json:"effect_live_day" gorm:"column:effect_live_day"`
	LiveTime        int64   `json:"live_time" gorm:"column:live_time"`
	AllIncome       float64 `json:"all_income" gorm:"column:all_income"`
	AllWaterDiamond int64   `json:"all_water_diamond"`
}

type GuildTotalDataDetailResp struct {
	UserId                 int64   `json:"user_id" gorm:"column:user_id"`                                       // 用户uid
	Nickname               string  `json:"nickname" gorm:"column:nickname"`                                     // 用户昵称
	RewardMoneyForHost     float64 `json:"reward_money_for_host" gorm:"column:reward_money_for_host"`           //主播金币档位奖励给主播 印尼盾 分
	TotalIncome            float64 `json:"total_income" gorm:"column:total_income"`                             //
	HostRate               int     `json:"host_rate" gorm:"column:host_rate"`                                   //对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数
	HostRateValue          string  `json:"host_rate_value" gorm:"column:host_rate_value"`                       //对应中reward_formula_for_host的最后奖励比例，比如：50% *100的整数
	GiftIncome             float64 `json:"gift_income" gorm:"column:gift_income"`                               //礼物收入=reward_formula_for_host 印尼盾 分
	GuildRate              int     `json:"guild_rate" gorm:"column:guild_rate"`                                 //公会礼物提成比列 *100的整数
	GuildRateValue         string  `json:"guild_rate_value" gorm:"column:guild_rate_value"`                     //公会礼物提成比列 *100的整数
	GuildGiftIncome        float64 `json:"guild_gift_income" gorm:"column:guild_gift_income"`                   //公会礼物提成收入 印尼盾 分
	WeekRankRewardForHost  float64 `json:"week_rank_reward_for_host" gorm:"column:week_rank_reward_for_host"`   //周榜主播奖励（印尼盾）同时需要按照例如：周1至周3工作几天时长多久来发放，如果排行榜靠前，时长或出勤天未达到同样不发放奖励 印尼盾 分
	WeekRankRewardForGuild float64 `json:"week_rank_reward_for_guild" gorm:"column:week_rank_reward_for_guild"` //周榜公会获得提成收入 印尼盾 分
	HostTotalIncome        float64 `json:"host_total_income" gorm:"column:host_total_income"`                   //主播总收入 印尼盾 分
	GuildCutRate           string  `gorm:"column:guild_cut_rate" json:"guild_cut_rate"`                         //公会抽成 设置范围0-100
	GuildReward            float64 `json:"guild_reward" gorm:"column:guild_reward"`                             //入驻奖励=live_reward_for_guild+activity__reward_for_guild（分4周）印尼盾 分
	GuildTotalIncome       float64 `json:"guild_total_income" gorm:"column:guild_total_income"`                 //公会总收入 印尼盾 分
	ActivityRewardForGuild float64 `json:"activity_reward_for_guild"`                                           //活动奖励给公会 印尼盾 分 比如分四周奖励 活动奖励给公会 印尼盾 分 比如分四周奖励 该字段不受是否结算字段影响（settle_state）：即便主播未达到结算标准，该字段一样结算给公会
	WeekWaterDiamond       int64   `json:"week_water_diamond"`                                                  //本周流水
}

type AnchorLastLiveResp struct {
	MasterId     int64  `json:"master_id"`   //公会长id
	MasterName   string `json:"master_name"` //公会长id
	UserId       int64  `json:"user_id"`     // 用户uid
	ServerId     int64  `json:"server_id"`   // 运营号id
	LastLiveTime int64  `json:"last_live_time" `
}

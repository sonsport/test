package model

type ALevelAnchorInfo struct {
	UserId         int64  `json:"user_id"` //用户ID
	Level          string `json:"level"`   //标签
	CreateTime     int64  `json:"create_time"`
	EffectLiveTime int64  `json:"effect_live_time"`
	EffectDay      int    `json:"effect_day"`
	FeedNum        int    `json:"feed_num"`
	RewardDayNum   int64  `json:"reward_day_num"`
}

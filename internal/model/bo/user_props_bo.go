package bo

type GiftPropsBO struct {
	UserId    int64 `json:"user_id"`
	FromType  int   `json:"from_type"`  //来源类型 4 使用himilive 5 点赞 6关注 7排行榜 8 看播1分钟 9 看播5分钟 10 看播30分钟 11 公屏聊天 12 充值 13 送礼物 14 购买装扮
	EventTime int64 `json:"event_time"` //事件时间
	TotalFee  int64 `json:"total_fee"`  //充值数额
}

type QuestBO struct {
	UserId        int64 `json:"user_id"`
	TargetUserId  int64 `json:"target_user_id"` //对方id 用于关注被关注  分享被分享
	QuestType     int   `json:"quest_type"`     //任务类型 common.QuestTypeSignIn
	QuestAchieved int64 `json:"quest_achieved"` //增加任务达成值 1 次数任务代表1次 计时为1分钟
	EventTime     int64 `json:"event_time"`     //事件时间
}

type ActivitObject struct {
	UserId       int64  `json:"user_id"`
	ActivityType int    `json:"activity_type"` //活动类型
	EventTime    int64  `json:"event_time"`    //事件时间
	ObjectExtra  string `json:"object_extra"`  //扩展对象
}

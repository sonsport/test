package model

type BeginLiveReq struct {
	RoomId        string `json:"room_id" v:"required"` //直播间名称
	RoomName      string `json:"room_name"`            //房间名
	ImagePath     string `json:"image_path"`           //图片路径
	GameId        int64  `json:"game_id"`              //游戏在系统内的id
	PrivateStatus int    `json:"private_status"`       //0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
	UnlockPrice   int64  `json:"unlock_price"`         //钻石
	Password      string `json:"password"`             //密码
	WishGiftId    int64  `json:"wish_gift_id"`         //心愿礼物id
	WishGiftCount int64  `json:"wish_gift_count"`      //心愿礼物个数
	RoomType      int    `json:"room_type"`            //房间类型 0 普通直播间  2 多人直播间
	ReviewUpUser  int64  `json:"review_up_user"`       //上麦审核
	MaxMultiple   int    `json:"max_multiple"`         //最大值窗口数
	LayoutSet     int    `json:"layout_set"`           //布局设置  1 突出 2 平铺
}

type ApplyLiveMultipleReq struct {
	RoomId       string `json:"room_id" v:"required"`        //直播间id
	AnchorUserId int64  `json:"anchor_user_id" v:"required"` //直播间id
}

type CancelLiveMultipleReq struct {
	RoomId       string `json:"room_id" v:"required"`        //直播间id
	AnchorUserId int64  `json:"anchor_user_id" v:"required"` //直播间id
}

type UpdateMultipleUserSetReq struct {
	RoomId           string `json:"room_id" v:"required"` //直播间id
	ReviewUpUser     int64  `json:"review_up_user"`       //多人房上麦审核
	MultipleUserType int64  `json:"multiple_user_type"`   //多人房上麦权限  1 所有人 2 关注我的人 3我的好友
	RoomType         int64  `json:"room_type"`            //房间类型 0 普通直播间  2 多人直播间
	LayoutSet        int64  `json:"layout_set"`           //布局设置  1 突出 2 平铺
	VideoControl     int64  `json:"video_control"`        //全员视频控制 1 禁止
	SoundControl     int64  `json:"sound_control"`        //全员音频控制 1 禁止
}

type UpdateMultipleReq struct {
	RoomId         string `json:"room_id" v:"required"`          //直播间id
	RoomMultipleId int64  `json:"room_multiple_id" v:"required"` //多人房队列id
	LiveStates     int64  `json:"live_states"`                   //多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略
	LiveControl    int64  `json:"live_control"`                  //直播控制 1 正常 2 禁视频
	SoundControl   int64  `json:"sound_control"`                 //直播控制 1 正常 2 禁音频
}

type GetMultipleReq struct {
	RoomId string `json:"room_id" v:"required"` //直播间id
}

type InviteLiveMultipleReq struct {
	RoomId string `json:"room_id" v:"required"` //直播间id
	UserId int64  `json:"user_id" v:"required"` //直播间id
}

type AgreeInviteLiveMultipleReq struct {
	RoomId string `json:"room_id" v:"required"` //直播间id
}

type GetStatisticsByRoomReq struct {
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

type MultipleListRes struct {
	Id               int64      `json:"id"`
	UserId           int64      `json:"user_id"`            //用户id
	Nickname         string     `json:"nickname"`           //昵称
	Portrait         string     `json:"portrait"`           //头像
	RoomId           string     `json:"room_id"`            //房间id
	AnchorUserId     int64      `json:"anchor_user_id"`     //主播用户id
	LiveControl      int64      `json:"live_control"`       //直播控制 1 正常 2 禁视频
	SoundControl     int64      `json:"sound_control"`      //直播控制 1 正常 2 静音
	LiveStates       int64      `json:"live_states"`        //多人房状态 0 申请中 1 上麦中 2 已拒绝 3 已结束 4 忽略
	StreamId         string     `json:"stream_id"`          //流id
	MultipleUserRole int64      `json:"multiple_user_role"` //多人房身份 1 多人房主播 2 上麦者
	MainStatus       int64      `json:"main_status"`        //主麦状态 1主麦
	Top              int32      `json:"top"`                //为该输入流设置图片，用于覆盖视频内容；使用图片时，不显示视频内容。
	Left             int32      `json:"left"`
	Bottom           int32      `json:"bottom"`
	Right            int32      `json:"right"`
	UserDressUps     *DressUpVo `json:"user_dress_ups"` //用户装扮
}

type MultipleUserRes struct {
	ReviewUpUser     int64 `json:"review_up_user"`     //多人房上麦审核
	MultipleUserType int64 `json:"multiple_user_type"` //多人房上麦权限  1 所有人 2 关注我的人 3我的好友
	VideoControl     int64 `json:"video_control"`      //全员视频控制 1 禁止
	SoundControl     int64 `json:"sound_control"`      //全员音频控制 1 禁止
}

type RoomDetailRes struct {
	UserId            int64  `json:"user_id" form:"user_id"`                 //主播ID
	RoomName          string `json:"room_name" form:"room_name"`             //直播间名称
	RoomId            string `json:"room_id" form:"room_id"`                 //第三方直播间ID全局唯一
	StreamId          string `json:"stream_id" form:"stream_id"`             //推拉流ID全局唯一
	ImagePath         string `json:"image_path" form:"image_path"`           //图片路径
	Status            int    `json:"status" form:"status"`                   //直播间状态 0 预开播 1 直播中 2 直播结束 3 直播异常
	BeginTime         int64  `json:"begin_time" form:"begin_time"`           //直播开始时间
	Role              int8   `json:"role" form:"role"`                       //主播类型 0 用户 1 主播
	FocusCount        int64  `json:"focus_count" form:"focus_count"`         //本次直播涨粉数量
	BattleStatus      int    `json:"battle_status" form:"battle_status"`     //pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK
	PrivateStatus     int    `json:"private_status" form:"private_status"`   //0 正常直播间，1 小黑屋，2 钻石房间 3 密码房间
	UnlockPrice       int64  `json:"unlock_price" form:"unlock_price"`       //解锁价格
	Password          string `json:"password" form:"password"`               //解锁价格
	GameId            int64  `json:"game_id" form:"game_id"`                 //游戏id
	WishGiftId        int64  `json:"wish_gift_id" form:"wish_gift_id"`       //心愿礼物id
	WishGiftCount     int64  `json:"wish_gift_count" form:"wish_gift_count"` //心愿礼物个数
	RoomType          int64  `json:"room_type"  form:"room_type"`            //房间类型 0 普通直播间  2 多人直播间
	MaxMultiple       int64  `json:"max_multiple"  form:"max_multiple"`      //最大值窗口数
	LayoutSet         int64  `json:"layout_set"  form:"layout_set"`          //布局设置  1 突出 2 平铺
	MixedStreamId     string `json:"mixed_stream_id"`                        //混流id
	WideMixedStreamId string `json:"wide_mixed_stream_id"`                   //宽屏混流id
	GuildId           int64  `json:"guild_id"`                               //工会id
	GuildName         string `json:"guild_name"`                             //工会name
	ZegoToken         string `json:"zego_token"`                             //zego token
	GuestId           string `json:"guest_id"`                               //
	NickName          string `json:"nick_name"`                              //
	Portrait          string `json:"portrait"`                               //
}

type UpdateMainMultipleReq struct {
	MaxMultiple int64  `json:"max_multiple"`         //最大值窗口数
	RoomId      string `json:"room_id" v:"required"` //直播间id
	Id          int64  `json:"id"`
	MainStatus  int64  `json:"main_status"`
}

type UpdateMaxMultipleReq struct {
	MaxMultiple int64  `json:"max_multiple"`         //最大值窗口数
	RoomId      string `json:"room_id" v:"required"` //直播间id
}

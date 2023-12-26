package bo

type RoomLiveBusinessObject struct {
	RoomId    string `json:"room_id"`
	UserId    int64  `json:"user_id"`
	BeginTime string `json:"begin_time"`
}

type RoomHeartBO struct {
	Score     int   `json:"score"`
	CountTime int   `json:"count"`
	BeginTime int64 `json:"beginTime"`
	//PreLastHeartbeatTime  int64 `json:"pre_last_heartbeat_time"`
	LastHeartbeatTime int64 `json:"last_heartbeat_time"`
}

// ShortScoreSct 房间临时积分
type ShortScoreSct struct {
	UserId        int64 `json:"user_id"`
	BeginTime     int64 `json:"begin_time"`
	Score         int64 `json:"score"`
	TimeScore     int   `json:"time_score"`
	FixedPosition int   `json:"fixedPosition"` //固定第几位，配合系统临时积分使用，0 没有固定，>0 有固定位置，由管理后台配置
}

type RoomAuthorAwardBO struct {
	Diamonds  int   `json:"diamonds"`
	Count     int   `json:"count"`
	BeginTime int64 `json:"beginTime"`
	LastTime  int64 `json:"last_heartbeat_time"`
}

// LiveRoomBO 直播间记录表
type LiveRoomBO struct {
	Id                int64  `db:"id" json:"id" form:"id"`                                                    //直播间Id
	UserId            int64  `db:"user_id" json:"user_id" form:"user_id"`                                     //主播ID
	RoomName          string `db:"room_name" json:"room_name" form:"room_name"`                               //直播间名称
	RoomId            string `db:"room_id" json:"room_id" form:"room_id"`                                     //第三方直播间ID全局唯一
	StreamId          string `db:"stream_id" json:"stream_id" form:"stream_id"`                               //推拉流ID全局唯一
	Title             string `db:"title" json:"title" form:"title"`                                           //直播间标题
	ImagePath         string `db:"image_path" json:"image_path" form:"image_path"`                            //图片路径
	Score             int    `db:"score" json:"score" form:"score"`                                           //排序分值
	MemberCount       int64  `db:"member_count" json:"member_count" form:"member_count"`                      //累计观看人数
	Status            int    `db:"status" json:"status" form:"status"`                                        //直播间状态 0 预开播 1 直播中 2 直播结束 3 直播异常
	CoinAmount        int64  `db:"coin_amount" json:"coin_amount" form:"coin_amount"`                         //打赏金币数量 放大100倍
	BeginTime         int64  `db:"begin_time" json:"begin_time" form:"begin_time"`                            //直播开始时间
	EndTime           int64  `db:"end_time" json:"end_time" form:"end_time"`                                  //直播结束时间
	LastHeartbeatTime int64  `db:"last_heartbeat_time" json:"last_heartbeat_time" form:"last_heartbeat_time"` //最后一次心跳时间
	Role              int8   `db:"role" json:"role" form:"role"`                                              //主播类型 0 用户 1 主播
	FocusCount        int64  `db:"focus_count" json:"focus_count" form:"focus_count"`                         //本次直播涨粉数量
	BattleStatus      int    `db:"battle_status" json:"battle_status" form:"battle_status"`                   //pk状态 0 未pk 1 随机PK匹配中 2 邀请PK中 3 正在PK
	PrivateStatus     int    `db:"private_status" json:"private_status" form:"private_status"`                //0 正常直播间，1 小黑屋，2私密直播间
	CreateTime        int64  `db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime        int64  `db:"update_time" json:"update_time" form:"update_time"`
}

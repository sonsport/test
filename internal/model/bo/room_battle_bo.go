package bo

type RoomBattleBusinessObject struct {
	BattleId     int64  `json:"battle_id"`
	MsgType      int64  `json:"msg_type"`
	RoomId       string `json:"room_id"`
	AnchorUserId int64  `json:"anchor_user_id"`
	BeginTime    string `json:"begin_time"`
}

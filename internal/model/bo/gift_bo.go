package bo

type SendGiftBusinessObject struct {
	Gid             int64  `json:"gid"` //礼物id
	UserId          int64  `json:"userId"`
	ReceiverId      int64  `json:"receiverId"`      // 接收人
	Number          int64  `json:"number"`          // 个数
	RoomId          string `json:"roomId"`          // 房间id
	WinningMultiple int64  `json:"winningMultiple"` // 倍数
	BalanceLogId    int64  `json:"balanceLogId"`    // 钻石数
	Diamonds        int64  `json:"diamonds"`        // 钻石数
	GiftLogId       int64  `json:"giftLogId"`       // 钻石数
	CashbackRatio   int64  `json:"cashback_ratio"`  // 返现比例
}

type ActivityGiftMQObject struct {
	UserId int64  `json:"userId"`
	Gid    int64  `json:"gid"`  // 礼物Id 自增
	Name   string `json:"name"` // 礼物名称

}

type AnchorIncomePushBo struct {
	UserId     int64  `json:"userId"`
	ReceiverId int64  `json:"receiverId"` // 接收人
	RoomId     string `json:"roomId"`     // 房间id
	IsQuest    int64  `json:"is_quest"`
}

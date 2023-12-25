package game

type UserRes struct {
	UID      int    `json:"uid"`
	Nick     string `json:"nick"`
	Gender   int    `json:"gender"`
	Portrait string `json:"portrait"`
}

type CocosUserRes struct {
	UID      int    `json:"uid"`
	Nick     string `json:"nick"`
	Gender   int    `json:"gender"`
	Portrait string `json:"portrait"`
	Birthday string `json:"birthday"`
}

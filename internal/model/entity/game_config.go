// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GameConfig is the golang structure for table game_config.
type GameConfig struct {
	Id                 uint    `json:"id"                 description:""`
	GameId             string  `json:"gameId"             description:"游戏ID"`
	GameKey            string  `json:"gameKey"            description:"第三方业务唯一标识"`
	GameName           string  `json:"gameName"           description:"游戏名称"`
	IconUrl            string  `json:"iconUrl"            description:"图片地址"`
	MiniIconUrl        string  `json:"miniIconUrl"        description:"小图标"`
	GameUrl            string  `json:"gameUrl"            description:"游戏地址"`
	GameType           int     `json:"gameType"           description:"游戏类型 0自研 1joyplay 2cocos"`
	IsMini             int     `json:"isMini"             description:"是否半屏 0全屏 1半屏"`
	IsHot              int     `json:"isHot"              description:"是否热门 0不是 1是热门"`
	SystemRate         int     `json:"systemRate"         description:"抽水比例  *10000的整数 最小为万分之一"`
	AnchorRate         int     `json:"anchorRate"         description:"主播分成  *10000的整数 最小为万分之一"`
	Desc               string  `json:"desc"               description:"备注"`
	GameVersion        string  `json:"gameVersion"        description:"游戏版本"`
	Md5Version         string  `json:"md5Version"         description:"MD5版本"`
	CocosEngineVersion string  `json:"cocosEngineVersion" description:"cocos引擎版本"`
	ShowSystem         int     `json:"showSystem"         description:"显示系统 0 全部 1 android 2 ios"`
	WinRatio           float64 `json:"winRatio"           description:"游戏宽高比"`
	Entrypoint         string  `json:"entrypoint"         description:"Ws接入点"`
	Status             int     `json:"status"             description:"状态 1 正常 2 下架"`
	Sort               int     `json:"sort"               description:"排序"`
	Orientation        string  `json:"orientation"        description:"游戏方向 横竖屏"`
	IsFullWin          int     `json:"isFullWin"          description:"是否全屏 1全屏"`
	CreateTime         int64   `json:"createTime"         description:""`
	UpdateTime         int64   `json:"updateTime"         description:""`
}

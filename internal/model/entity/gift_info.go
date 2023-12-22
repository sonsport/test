// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GiftInfo is the golang structure for table gift_info.
type GiftInfo struct {
	Gid              uint64 `json:"gid"              description:"礼物Id 自增"`
	Name             string `json:"name"             description:"礼物名称"`
	Diamonds         uint   `json:"diamonds"         description:"价格-钻石"`
	OriginalDiamonds int    `json:"originalDiamonds" description:"原价"`
	Type             uint   `json:"type"             description:"礼物大类 0普通礼物  1幸运礼物"`
	GiftType         uint   `json:"giftType"         description:"0普通小礼物（无特效）1特效礼物  2 房间红包雨 3 全站红包雨 4 全站通知礼物"`
	Icon             string `json:"icon"             description:"礼物图标"`
	AnimEffectUrl    string `json:"animEffectUrl"    description:"特效礼物url"`
	FileMd5          string `json:"fileMd5"          description:"礼物md5值"`
	State            int    `json:"state"            description:"0下架  1上架"`
	Sort             uint   `json:"sort"             description:"排序 顺序"`
	ResVoiceUrl      string `json:"resVoiceUrl"      description:"声音"`
	ResType          int    `json:"resType"          description:"1 icon/2 gif/3 文件资源；4为3d；5为面部识别；6为svga；"`
	IsBatter         int    `json:"isBatter"         description:"是否连送"`
	MinLevel         int    `json:"minLevel"         description:"最小等级"`
	LuckPoolType     int    `json:"luckPoolType"     description:"奖池类型 0 90% 1 95%"`
	Badge            string `json:"badge"            description:"礼物徽章"`
	StartTime        int64  `json:"startTime"        description:""`
	EndTime          int64  `json:"endTime"          description:""`
	CreatedTime      int64  `json:"createdTime"      description:""`
	UpdatedTime      int64  `json:"updatedTime"      description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GiftInfo is the golang structure of table gift_info for DAO operations like Where/Data.
type GiftInfo struct {
	g.Meta           `orm:"table:gift_info, do:true"`
	Gid              interface{} // 礼物Id 自增
	Name             interface{} // 礼物名称
	Diamonds         interface{} // 价格-钻石
	OriginalDiamonds interface{} // 原价
	Type             interface{} // 礼物大类 0普通礼物  1幸运礼物
	GiftType         interface{} // 0普通小礼物（无特效）1特效礼物  2 房间红包雨 3 全站红包雨 4 全站通知礼物
	Icon             interface{} // 礼物图标
	AnimEffectUrl    interface{} // 特效礼物url
	FileMd5          interface{} // 礼物md5值
	State            interface{} // 0下架  1上架
	Sort             interface{} // 排序 顺序
	ResVoiceUrl      interface{} // 声音
	ResType          interface{} // 1 icon/2 gif/3 文件资源；4为3d；5为面部识别；6为svga；
	IsBatter         interface{} // 是否连送
	MinLevel         interface{} // 最小等级
	LuckPoolType     interface{} // 奖池类型 0 90% 1 95%
	Badge            interface{} // 礼物徽章
	StartTime        interface{} //
	EndTime          interface{} //
	CreatedTime      interface{} //
	UpdatedTime      interface{} //
}

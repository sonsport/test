// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GameConfig is the golang structure of table game_config for DAO operations like Where/Data.
type GameConfig struct {
	g.Meta             `orm:"table:game_config, do:true"`
	Id                 interface{} //
	GameId             interface{} // 游戏ID
	GameKey            interface{} // 第三方业务唯一标识
	GameName           interface{} // 游戏名称
	IconUrl            interface{} // 图片地址
	MiniIconUrl        interface{} // 小图标
	GameUrl            interface{} // 游戏地址
	GameType           interface{} // 游戏类型 0自研 1joyplay 2cocos
	IsMini             interface{} // 是否半屏 0全屏 1半屏
	IsHot              interface{} // 是否热门 0不是 1是热门
	SystemRate         interface{} // 抽水比例  *10000的整数 最小为万分之一
	AnchorRate         interface{} // 主播分成  *10000的整数 最小为万分之一
	Desc               interface{} // 备注
	GameVersion        interface{} // 游戏版本
	Md5Version         interface{} // MD5版本
	CocosEngineVersion interface{} // cocos引擎版本
	ShowSystem         interface{} // 显示系统 0 全部 1 android 2 ios
	WinRatio           interface{} // 游戏宽高比
	Entrypoint         interface{} // Ws接入点
	Status             interface{} // 状态 1 正常 2 下架
	Sort               interface{} // 排序
	Orientation        interface{} // 游戏方向 横竖屏
	IsFullWin          interface{} // 是否全屏 1全屏
	CreateTime         interface{} //
	UpdateTime         interface{} //
}

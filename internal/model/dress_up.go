package model

type UserDressUp struct {
	Id         int64 `json:"id" gorm:"column:id"`
	UserId     int64 `json:"user_id" gorm:"column:user_id"`         // uid
	ShopId     int64 `json:"shop_id" gorm:"column:shop_id"`         // 商品ID
	CategoryId int   `json:"category_id" gorm:"column:category_id"` // 分类ID
	StartTime  int64 `json:"start_time" gorm:"column:start_time"`   // 开始时间
	EndTime    int64 `json:"end_time" gorm:"column:end_time"`       // 结束时间 -1标识永久配置
	EffectDays int   `json:"effect_days" db:"column:effect_days"`   // 未佩戴过的装扮的有效期，如果佩戴过设置为0
	Status     int8  `json:"status" gorm:"column:status"`           // 状态  0-未佩戴 1-佩戴；2取消佩戴 只有新购买为0
	CreateTime int64 `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime int64 `json:"update_time" gorm:"column:update_time"` // 更新时间
}

// UserDressUpListItem 装扮列表项
type UserDressUpListItem struct {
	//ID          int64  `json:"id" gorm:"column:id"`
	ShopID      int64  `json:"shop_id" gorm:"column:shop_id"`         // 商品ID
	ShopName    string `json:"shop_name" gorm:"column:shop_name"`     // 商品名称
	ResUrl      string `json:"res_url" gorm:"column:res_url"`         // 资源地址
	IconUrl     string `json:"icon_url" gorm:"column:icon_url"`       // 图片地址
	CoinKind    string `json:"coin_kind" gorm:"column:coin_kind"`     // 币种
	Price       int64  `json:"price" gorm:"column:price"`             // 单价
	Desc        string `json:"desc" gorm:"column:desc"`               // 商品描述
	CategoryID  int    `json:"category_id" gorm:"column:category_id"` // 分类ID
	StartTime   int64  `json:"start_time" gorm:"column:start_time"`   // 开始时间
	EndTime     int64  `json:"end_time" gorm:"column:end_time"`       // 结束时间
	IsCountDown bool   `json:"is_count_down"`                         // 是否进行倒计时
	IsRenew     bool   `json:"is_renew"`                              // 是否续费
	IsForever   bool   `json:"is_forever"`                            // 是否为永久礼物
	Status      int8   `json:"status" gorm:"column:status"`           // 状态 0-未佩戴 1-佩戴 2-过期
}

var (
	//MonthDays = 30 // 单月天数

	// 装扮状态

	UserDressUpStatusOfUnWear int8 = 0 // 未佩戴
	UserDressUpStatusOfWear   int8 = 1 // 佩戴

	// 个性装扮H5显示状态

	UserDressUpStatusOfExpire = 2 // 过期

	UserDressUpEffectDaysOfForever = -1 // 永久日期
)

// UserDressUpListCategory 装扮列表
type UserDressUpListCategory struct {
	ID    int                   `json:"id"`
	Name  string                `json:"name"`
	Desc  string                `json:"desc"`
	Items []UserDressUpListItem `json:"items"`
}

type DressUpItem struct {
	ShopId  int64  `json:"shop_id"`
	ResUrl  string `json:"res_url"`
	Name    string `json:"name"`
	IconUrl string `json:"icon_url"`
}

type HonoraryTitles struct {
	Title      string `json:"title"`
	StyleColor string `json:"style_color"`
}

type HonoraryTitlesExtra struct {
	StyleColor   string                             `json:"style_color"`
	LanguageList []*HonoraryTitlesExtraLanguageList `json:"language_list"`
}

type HonoraryTitlesExtraLanguageList struct {
	Lang  string `json:"lang"`
	Title string `json:"title"`
}

type ShopGoodsBubbleFrameExtra struct {
	BorderColor     string                                 `json:"border_color"`     // 边框颜色
	BackgroundColor string                                 `json:"background_color"` // 背景颜色
	BorderImages    []ShopGoodsBubbleFrameExtraBorderImage `json:"border_images"`    // 边框图片
}

type ShopGoodsBubbleFrameExtraBorderImage struct {
	ResUrl string `json:"res_url"` // 资源地址
	Width  int    `json:"width"`   // 宽度
	Height int    `json:"height"`  // 高度
}

type ShopGoodsBubbleDressUpItem struct {
	SuperId int64 `json:"super_id"` //靓号
	DressUpItem
	ShopGoodsBubbleFrameExtra
}

type UserHadWearDressUp struct {
	BubbleBox      *ShopGoodsBubbleDressUpItem `json:"bubble_box"`      // 气泡框
	GoldenLight    *DressUpItem                `json:"golden_light"`    // 金光
	HomeCover      *DressUpItem                `json:"home_cover"`      // 首页封面
	AvatarFrame    *DressUpItem                `json:"avatar_frame"`    // 头像框
	Ride           *DressUpItem                `json:"ride"`            // 座驾
	Badge          *DressUpItem                `json:"badge"`           // 徽章
	ActivityBadges []*DressUpItem              `json:"activity_badges"` // 活动徽章
	HonoraryTitles []*HonoraryTitles           `json:"honorary_titles"` // 荣誉称号
}

// DressUpEndTime 装扮倒计时
type DressUpEndTime struct {
	EndTime     int64 `json:"end_time"`
	IsCountDown bool  `json:"is_count_down"` // 是否进行倒计时
}

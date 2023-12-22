// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AreaInfo is the golang structure for table area_info.
type AreaInfo struct {
	Id           uint64 `json:"id"           description:""`
	AreaCode     int    `json:"areaCode"     description:"地区"`
	Path         string `json:"path"         description:""`
	Title        string `json:"title"        description:""`
	TitleCn      string `json:"titleCn"      description:""`
	Currency     string `json:"currency"     description:"货币编码"`
	CurrencyCode int    `json:"currencyCode" description:"国家编码"`
	Sort         int    `json:"sort"         description:""`
	IsRich       int    `json:"isRich"       description:"是否富有，0普通，1富有"`
	IsShow       int    `json:"isShow"       description:""`
	CreatedTime  int64  `json:"createdTime"  description:""`
	UpdatedTime  int64  `json:"updatedTime"  description:""`
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BeautifulNum is the golang structure for table beautiful_num.
type BeautifulNum struct {
	Id             int64  `json:"id"             description:""`
	SuperId        int64  `json:"superId"        description:"靓号"`
	ClassifyId     int64  `json:"classifyId"     description:"分类id"`
	IconUrl        string `json:"iconUrl"        description:"icon地址"`
	ShopStatus     int    `json:"shopStatus"     description:"状态 1可用 2不可用"`
	SortWeight     int    `json:"sortWeight"     description:"排序权重"`
	SkuListPrice   string `json:"skuListPrice"   description:"规格列表价格"`
	UserLevelLimit int    `json:"userLevelLimit" description:"用户等级限制"`
	Price          int    `json:"price"          description:"单价"`
	UserId         int64  `json:"userId"         description:"用户id"`
	IsWear         int    `json:"isWear"         description:"是否佩戴"`
	Experience     int64  `json:"experience"     description:"经验值"`
	ExpirationTime int64  `json:"expirationTime" description:"过期时间"`
	ReclaimTime    int64  `json:"reclaimTime"    description:"回收时间"`
	CreateTime     int64  `json:"createTime"     description:""`
	UpdateTime     int64  `json:"updateTime"     description:""`
}

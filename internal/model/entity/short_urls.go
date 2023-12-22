// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShortUrls is the golang structure for table short_urls.
type ShortUrls struct {
	Id         uint64 `json:"id"         description:""`
	UserId     int64  `json:"userId"     description:""`
	RawUrl     string `json:"rawUrl"     description:""`
	Source     string `json:"source"     description:""`
	ShortCode  string `json:"shortCode"  description:""`
	CreateTime int64  `json:"createTime" description:""`
}

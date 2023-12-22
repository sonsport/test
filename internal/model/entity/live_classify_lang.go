// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LiveClassifyLang is the golang structure for table live_classify_lang.
type LiveClassifyLang struct {
	Id               int    `json:"id"               description:""`
	ClassifyId       int    `json:"classifyId"       description:"分类id"`
	ClassifyLang     string `json:"classifyLang"     description:"分类语言"`
	ClassifyName     string `json:"classifyName"     description:"分类名称"`
	ClassifyIcon     string `json:"classifyIcon"     description:"分类图标"`
	ClassifyPickIcon string `json:"classifyPickIcon" description:""`
	CreateTime       int64  `json:"createTime"       description:""`
	UpdateTime       int64  `json:"updateTime"       description:""`
}

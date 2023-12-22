// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShortUrls is the golang structure of table short_urls for DAO operations like Where/Data.
type ShortUrls struct {
	g.Meta     `orm:"table:short_urls, do:true"`
	Id         interface{} //
	UserId     interface{} //
	RawUrl     interface{} //
	Source     interface{} //
	ShortCode  interface{} //
	CreateTime interface{} //
}

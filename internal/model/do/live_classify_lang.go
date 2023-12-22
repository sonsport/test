// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LiveClassifyLang is the golang structure of table live_classify_lang for DAO operations like Where/Data.
type LiveClassifyLang struct {
	g.Meta           `orm:"table:live_classify_lang, do:true"`
	Id               interface{} //
	ClassifyId       interface{} // 分类id
	ClassifyLang     interface{} // 分类语言
	ClassifyName     interface{} // 分类名称
	ClassifyIcon     interface{} // 分类图标
	ClassifyPickIcon interface{} //
	CreateTime       interface{} //
	UpdateTime       interface{} //
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// BeautifulNumClassify is the golang structure of table beautiful_num_classify for DAO operations like Where/Data.
type BeautifulNumClassify struct {
	g.Meta       `orm:"table:beautiful_num_classify, do:true"`
	ClassifyId   interface{} // 分类id
	ClassifyName interface{} //
	TranslateKey interface{} // 翻译key
	IsShow       interface{} // 是否显示
	CreateTime   interface{} //
	UpdateTime   interface{} //
}

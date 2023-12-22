// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LiveClassify is the golang structure of table live_classify for DAO operations like Where/Data.
type LiveClassify struct {
	g.Meta              `orm:"table:live_classify, do:true"`
	Id                  interface{} //
	ParentId            interface{} // 上级id 0 为顶级
	ClassifyType        interface{} // 1 推荐 2 最新 3 热门 4 pk 5 密码房 6 钻石房 7 标签主播 8 无标签主播 9 新秀 10 夜播
	ClassifyDefaultName interface{} // 默认分类名
	ClassifyRemark      interface{} // 分类备注
	ClassifyStatus      interface{} // 1 正常
	ClassifyRank        interface{} // 排序
	CreateTime          interface{} //
	UpdateTime          interface{} //
}

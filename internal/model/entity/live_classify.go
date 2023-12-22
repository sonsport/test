// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LiveClassify is the golang structure for table live_classify.
type LiveClassify struct {
	Id                  int    `json:"id"                  description:""`
	ParentId            int    `json:"parentId"            description:"上级id 0 为顶级"`
	ClassifyType        int    `json:"classifyType"        description:"1 推荐 2 最新 3 热门 4 pk 5 密码房 6 钻石房 7 标签主播 8 无标签主播 9 新秀 10 夜播"`
	ClassifyDefaultName string `json:"classifyDefaultName" description:"默认分类名"`
	ClassifyRemark      string `json:"classifyRemark"      description:"分类备注"`
	ClassifyStatus      int    `json:"classifyStatus"      description:"1 正常"`
	ClassifyRank        int    `json:"classifyRank"        description:"排序"`
	CreateTime          int64  `json:"createTime"          description:""`
	UpdateTime          int64  `json:"updateTime"          description:""`
}

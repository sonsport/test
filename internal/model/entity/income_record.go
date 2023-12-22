// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// IncomeRecord is the golang structure for table income_record.
type IncomeRecord struct {
	Id           uint64 `json:"id"           description:"id"`
	UserId       int64  `json:"userId"       description:"用户id"`
	SourceId     int64  `json:"sourceId"     description:"收益来源用户的id"`
	UserRecordId int64  `json:"userRecordId" description:"对应用户扣费记录id"`
	IncomeType   int    `json:"incomeType"   description:"收益类型"`
	LinkId       int64  `json:"linkId"       description:"引起变更的id"`
	Income       int64  `json:"income"       description:"收益变动数额"`
	BeforeChange int64  `json:"beforeChange" description:"变更之前余额"`
	AfterChange  int64  `json:"afterChange"  description:"变更之后余额"`
	CreateTime   int64  `json:"createTime"   description:"createdAt"`
	UpdateTime   int64  `json:"updateTime"   description:"updatedAt"`
}

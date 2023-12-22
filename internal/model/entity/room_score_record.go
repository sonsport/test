// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RoomScoreRecord is the golang structure for table room_score_record.
type RoomScoreRecord struct {
	Id            int64  `json:"id"            description:""`
	UserId        int64  `json:"userId"        description:""`
	RoomId        string `json:"roomId"        description:""`
	OperationType int    `json:"operationType" description:"操作类型"`
	OperationId   int64  `json:"operationId"   description:"操作id"`
	Score         int    `json:"score"         description:"排序分值"`
	BeforeScore   int    `json:"beforeScore"   description:"before排序分值"`
	AfterScore    int    `json:"afterScore"    description:"after排序分值"`
	CreateTime    int64  `json:"createTime"    description:""`
	UpdateTime    int64  `json:"updateTime"    description:""`
}

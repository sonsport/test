// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShareNewUserRecord is the golang structure for table share_new_user_record.
type ShareNewUserRecord struct {
	Id                uint64 `json:"id"                description:""`
	UserId            int64  `json:"userId"            description:""`
	OwnerUid          int64  `json:"ownerUid"          description:"分享者uid"`
	Type              int    `json:"type"              description:"1:新用户注册；2:新用户充值"`
	Status            int    `json:"status"            description:"0:初始化；1:待加币；2:加币成功"`
	OwnerOrderId      string `json:"ownerOrderId"      description:"分享者加币订单号"`
	OwnerDiamondNum   int    `json:"ownerDiamondNum"   description:"分享者获得的钻石数"`
	NewUserOrderId    string `json:"newUserOrderId"    description:"新用户加币订单"`
	NewUserDiamondNum int    `json:"newUserDiamondNum" description:"新用户获得的钻石数"`
	CodeId            int    `json:"codeId"            description:"分享ID"`
	RegisterTime      int64  `json:"registerTime"      description:"注册时间"`
	RechargeNum       int    `json:"rechargeNum"       description:"充值的虚拟币数量"`
	WatchLiveTime     int    `json:"watchLiveTime"     description:"观看直播时长"`
	Ip                string `json:"ip"                description:"登陆地IP"`
	Smid              string `json:"smid"              description:"设备ID"`
	DeviceName        string `json:"deviceName"        description:"设备型号"`
	UpdateTime        int64  `json:"updateTime"        description:""`
	CreateTime        int64  `json:"createTime"        description:""`
}

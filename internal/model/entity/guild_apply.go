// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GuildApply is the golang structure for table guild_apply.
type GuildApply struct {
	Id         uint   `json:"id"         description:"自增id"`
	UserId     int64  `json:"userId"     description:"申请人用户Id"`
	GuildName  string `json:"guildName"  description:"公会名称"`
	AnchorNum  int    `json:"anchorNum"  description:"主播数量"`
	Whatsapp   string `json:"whatsapp"   description:"联系WhatsApp"`
	State      int    `json:"state"      description:"状态：0待审核 1通过 2不通过"`
	Operator   string `json:"operator"   description:"审核人"`
	Remark     string `json:"remark"     description:"备注"`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

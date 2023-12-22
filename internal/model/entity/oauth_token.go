// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OauthToken is the golang structure for table oauth_token.
type OauthToken struct {
	UserId     uint64 `json:"userId"     description:""`
	OpenId     string `json:"openId"     description:""`
	LoginType  int    `json:"loginType"  description:""`
	CreateTime int64  `json:"createTime" description:""`
	UpdateTime int64  `json:"updateTime" description:""`
}

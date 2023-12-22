// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OauthToken is the golang structure of table oauth_token for DAO operations like Where/Data.
type OauthToken struct {
	g.Meta     `orm:"table:oauth_token, do:true"`
	UserId     interface{} //
	OpenId     interface{} //
	LoginType  interface{} //
	CreateTime interface{} //
	UpdateTime interface{} //
}

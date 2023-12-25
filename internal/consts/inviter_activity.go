package consts

// PullNewUserSource 分享来源方式
var PullNewUserSource = "pull_new_user" //分享拉新用户

var ShareSourceMap = map[string]int{
	PullNewUserSource: 1,
}

const (
	//奖励方式
	//AwardInviteUser   = 1 //新用户注册
	//AwardUserRecharge = 2 //新用户充值
	//RoomShare         = 3 //房间内分享的

	InviteStatPass    = 1  // 邀请状态正常
	InviteStatInvalid = -1 // 邀请状态失效

	//AdjustShareUrl ad分享链接
	AdjustShareUrl = "https://2h8v.adj.st/?adj_t=15vpaqob_12tfj5er&adj_deeplink=openappfuyalive%3A%2F%2Ffuya%3A80%2Fproduct&adj_label=#adj_label"
)

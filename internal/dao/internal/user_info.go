// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfoDao is the data access object for table user_info.
type UserInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserInfoColumns // columns contains all the column names of Table for convenient usage.
}

// UserInfoColumns defines and stores column names for table user_info.
type UserInfoColumns struct {
	UserId             string // 用户Id 自增
	Username           string // 用户名
	AreaCode           string // 地区编码
	Nickname           string // 昵称
	Password           string // 密码  空字符串说明没有密码
	Gender             string // 性别 1男 2女 0未知
	Birthday           string // 生日 时间戳
	Height             string // 身高 cm
	Education          string // 学历，枚举 小学 高中 大学 本科 扩展
	Industry           string // 行业，枚举 学生、文化/艺术 扩展
	Intro              string // 自我介绍 签名
	Portrait           string // 头像
	LiveState          string // 0未开播 1开播
	WhatsApp           string // whatsapp号码
	Role               string // 0普通用户 1主播 2管理员 3公会长 4运营号 后续扩展
	BanTime            string // 封禁时间，0代表未封禁
	LastHeartbeatAt    string // 用户心跳
	DeviceId           string // 设备号
	AppName            string // app名称
	AppChannel         string // 渠道
	AppVersion         string // app版本
	PhoneMode          string // 手机型号
	PhoneSystemVersion string // 手机系统版本
	LastIp             string // 最后一次登录ip
	System             string // 1为安卓 2为ios
	UserLanguage       string // 用户语言
	CurrencyCode       string // 国家数字码，比如印尼：360；马来：458  0为默认
	Remarks            string // 备注
	CreateTime         string // 创建时间
	UpdateTime         string // 最后更新时间
}

// userInfoColumns holds the columns for table user_info.
var userInfoColumns = UserInfoColumns{
	UserId:             "user_id",
	Username:           "username",
	AreaCode:           "area_code",
	Nickname:           "nickname",
	Password:           "password",
	Gender:             "gender",
	Birthday:           "birthday",
	Height:             "height",
	Education:          "education",
	Industry:           "industry",
	Intro:              "intro",
	Portrait:           "portrait",
	LiveState:          "live_state",
	WhatsApp:           "whats_app",
	Role:               "role",
	BanTime:            "ban_time",
	LastHeartbeatAt:    "last_heartbeat_at",
	DeviceId:           "device_id",
	AppName:            "app_name",
	AppChannel:         "app_channel",
	AppVersion:         "app_version",
	PhoneMode:          "phone_mode",
	PhoneSystemVersion: "phone_system_version",
	LastIp:             "last_ip",
	System:             "system",
	UserLanguage:       "user_language",
	CurrencyCode:       "currency_code",
	Remarks:            "remarks",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
}

// NewUserInfoDao creates and returns a new DAO object for table data access.
func NewUserInfoDao() *UserInfoDao {
	return &UserInfoDao{
		group:   "default",
		table:   "user_info",
		columns: userInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserInfoDao) Columns() UserInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

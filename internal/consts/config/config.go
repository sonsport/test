package config

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct {
	UrlConfig      UrlConfig       `json:"urlConfig"`
	AppChannels    []*AppChannel   `json:"appChannel"`
	ServerConfig   *ServerConfig   `json:"serverConfig"`
	XxlJob         *XxlJob         `json:"xxlJob"`
	BattleConfig   *BattleConfig   `json:"battleConfig"`
	RabbitMqConfig *RabbitMqConfig `json:"rabbitMqConfig"`
	IosAppStore    *IosAppStore    `json:"iosAppStore"`
	PayerMaxConfig *PayerMaxConfig `json:"payerMaxConfig"`
	HaiPayConfig   *HaiPayConfig   `json:"haiPayConfig"`
}

type IosAppStore struct {
	AppIdMap  map[string]string `json:"appIdMap"`
	BundleMap map[string]string `json:"bundleMap"`
}

type PayerMaxConfig struct {
	AppId            string `json:"appId"`
	MerchantNo       string `json:"merchantNo"`
	PayUrl           string `json:"payUrl"`
	FrontCallBackUrl string `json:"frontCallBackUrl"`
	CallbackUrl      string `json:"callbackUrl"`
}

type HaiPayConfig struct {
	MerchantNo  string `json:"merchantNo"`
	MerApiKey   string `json:"merApiKey"`
	PayUrl      string `json:"payUrl"`
	VaUrl       string `json:"vaUrl"`
	QrisUrl     string `json:"qrisUrl"`
	RedirectUrl string `json:"redirectUrl"`
	NotifyUrl   string `json:"notifyUrl"`
}

type RabbitMqConfig struct {
	Hosts       string `json:"hosts"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        int64  `json:"port"`
	VirtualHost string `json:"virtualHost"`
}

type BattleConfig struct {
	BattleTime  int `json:"battleTime"`
	PenaltyTime int `json:"penaltyTime"`
}

type XxlJob struct {
	Url         string `json:"url"`
	Port        string `json:"port"`
	RegistryKey string `json:"registryKey"`
}

type UrlConfig struct {
	H5Domain              string `json:"h5Domain"`
	PullNewUserUrl        string `json:"pullNewUserUrl"`
	ShareDomain           string `json:"shareDomain"`
	WatchLiveTime         int    `json:"watchLiveTime"`
	OwnerRewardDiamonds   int64  `json:"ownerRewardDiamonds"`
	NewUserRewardDiamonds int64  `json:"newUserRewardDiamonds"`
}

type AppChannel struct {
	AppName    string `json:"appName"`    //appName
	AppChannel string `json:"appChannel"` //appChannel
	AppSystem  string `json:"appSystem"`  //android or ios
	AppId      string `json:"appId"`      //fb appid
	AppSecret  string `json:"appSecret"`  //fb appSecret
}

type ServerConfig struct {
	DiamondsToIncome       int64                    `json:"diamondsToIncome"`       //钻石转换收益倍数 1:100
	LuckyGiftIncome        float64                  `json:"luckyGiftIncome"`        //幸运礼物收益率 钻石*0.15
	LuckyGiftAllRoomConfig []LuckyGiftAllRoomConfig `json:"luckyGiftAllRoomConfig"` //幸运礼物全站飘窗倍数
	LuckyDiamondsToExp     int64                    `json:"luckyDiamondsToExp"`     //幸运礼物 5 钻石=1经验值
	NormalDiamondsToExp    int64                    `json:"normalDiamondsToExp"`    //普通礼物 1 钻石=1经验值
	IncomeToExp            int64                    `json:"incomeToExp"`            //收益礼物经验 1 金币=1经验值
	IsOpenGreenLive        bool                     `json:"isOpenGreenLive"`        //是否开启直播绿网
	Ants                   Ants                     `json:"ants"`                   // 协程池配置
	ApplePopup             bool                     `json:"applePopup"`             // 苹果评分引导
}

type Ants struct {
	AnchorStatisticV2Size int `json:"anchorStatisticV2Size"` // 主播每日统计协程池大小
}

type LuckyGiftAllRoomConfig struct {
	MinMultiple int64  `json:"minMultiple"` // 最小倍数
	Colors      string `json:"colors"`      // 颜色
	ShowType    int    `json:"showType"`    // 类型
}

// GetAppChannelByNameAndChannel 根据appName和appChannel获取appChannel对象
func GetAppChannelByNameAndChannel(appName string, appChannel string) (*AppChannel, error) {
	for _, channel := range ConfigS.AppChannels {
		if channel.AppName == appName && channel.AppChannel == appChannel {
			return channel, nil
		}
	}
	return nil, fmt.Errorf("未找到该app渠道配置")
}

func GetServerConfig() *ServerConfig {
	return ConfigS.ServerConfig
}

func GetBattleConfig() *BattleConfig {
	return ConfigS.BattleConfig
}

// GeAppleAppIdByName 根据appName获取AppleAppId对象
func GeAppleAppIdByName(appName string) string {
	for key, val := range ConfigS.IosAppStore.AppIdMap {
		if appName == key {
			return val
		}
	}
	return ""
}

// GeAppleBundleByName 根据appName获取AppleBundle
func GeAppleBundleByName(appName string) string {
	for key, val := range ConfigS.IosAppStore.BundleMap {
		if appName == key {
			return val
		}
	}
	return ""
}

var ConfigS *Config

func InitConfig() error {
	ConfigS = new(Config)
	data, err := gcfg.Instance().Data(context.Background())
	if err != nil {
		return err
	}
	gconv.Scan(data, ConfigS)
	return nil
}

package utility

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/wxnacy/wgo/arrays"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"fuya-ark/internal/consts"
)

var i18n *gi18n.Manager

func init() {
	if i18n == nil {
		i18n = gi18n.New()
	}
}

func Str2Uids(s string) []int64 {
	idStrs := strings.Split(s, ",")
	var ids []int64
	for _, v := range idStrs {
		uid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		if uid <= 0 {
			continue
		}
		ids = append(ids, uid)
	}
	return ids
}

// SubStr 截取字符串，支持多字节字符
// start：起始下标，负数从从尾部开始，最后一个为-1
// length：截取长度，负数表示截取到末尾
func SubStr(str string, start int, length int) (result string) {
	s := []rune(str)
	total := len(s)
	if total == 0 {
		return
	}
	// 允许从尾部开始计算
	if start < 0 {
		start = total + start
		if start < 0 {
			return
		}
	}
	if start > total {
		return
	}
	// 到末尾
	if length < 0 {
		length = total
	}

	end := start + length
	if end > total {
		result = string(s[start:])
	} else {
		result = string(s[start:end])
	}

	return
}

// GetRemoteClientIp 获取远程客户端IP
func GetRemoteClientIp(r *ghttp.Request) string {
	var remoteIp string // r.RemoteAddr
	remoteIp = r.Header.Get("X-Forwarded-For")
	if len(remoteIp) == 0 || (len(remoteIp) > 0 && "unknown" == strings.ToLower(remoteIp)) {
		remoteIp = r.Header.Get("Proxy-Client-IP")
	}
	if len(remoteIp) == 0 || (len(remoteIp) > 0 && "unknown" == strings.ToLower(remoteIp)) {
		remoteIp = r.Header.Get("WL-Proxy-Client-IP")
	}
	if len(remoteIp) == 0 || (len(remoteIp) > 0 && "unknown" == strings.ToLower(remoteIp)) {
		remoteIp = r.Header.Get("x-real-ip")
	}
	if len(remoteIp) == 0 || (len(remoteIp) > 0 && "unknown" == strings.ToLower(remoteIp)) {
		remoteIp = r.RemoteAddr
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}
	//本地ip
	if remoteIp == "::1" {
		remoteIp = consts.LocalHost
	}
	//后续这里加个","切割判断 取第一个为真是ip
	ipArray := strings.Split(remoteIp, ",")
	if len(ipArray) > 1 {
		remoteIp = ipArray[0]
	}
	return remoteIp
}

// GetIps  获取本机ip地址
func GetIps() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isVailIpNet := address.(*net.IPNet)
		if isVailIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	fmt.Println("ips = ", ips)
	return ips
}

// StringIpToInt ip转int
func StringIpToInt(ipString string) (ipInt int) {
	ipSegs := strings.Split(ipString, ".")
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// IpIntToString int转ip串
func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var ipLen = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < ipLen; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[ipLen-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < ipLen; i++ {
		buffer.WriteString(ipSegs[i])
		if i < ipLen-1 {
			buffer.WriteString(".")
		}
	}

	return buffer.String()
}

func GoSafe(ctx context.Context, fn func()) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(ctx, "goroutine error:%v", err)
		}
	}()

	fn()
}

func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

//
//func TranslateByKey(ctx context.Context, key string) string {
//	reqData := common.Context.GetRequestData(ctx)
//	var header *model.ReqHeader
//	if reqData != nil {
//		header = reqData.Header
//	}
//	language := common.LanguageEN //默认为英文
//	if header != nil {
//		clientLanguage := header.ClientLanguage
//		if clientLanguage == common.LanguageEN || clientLanguage == common.LanguageID || clientLanguage == common.LanguageMS {
//			language = header.ClientLanguage
//		}
//	}
//	return i18n.GetContent(gi18n.WithLanguage(ctx, language), key)
//}
//
//func GetCurrentLanguage(ctx context.Context) string {
//	reqData := common.Context.GetRequestData(ctx)
//	var header *model.ReqHeader
//	if reqData != nil {
//		header = reqData.Header
//	}
//	language := common.LanguageEN //默认为英文
//	if header != nil {
//		language = header.ClientLanguage
//	}
//	return language
//}
//
//func TranslateByKeyAndLanguage(ctx context.Context, key, language string) string {
//	if len(language) != 2 {
//		language = common.LanguageEN //默认为英文
//	}
//	switch language {
//	case common.LanguageEN, common.LanguageID, common.LanguageMS:
//		break
//	default:
//		language = common.LanguageEN
//	}
//	return i18n.GetContent(gi18n.WithLanguage(ctx, language), key)
//}
//
//func TranslateFormat(ctx context.Context, content string, values ...interface{}) string {
//	reqData := common.Context.GetRequestData(ctx)
//	var header *model.ReqHeader
//	if reqData != nil {
//		header = reqData.Header
//	}
//	language := common.LanguageEN //默认为英文
//	if header != nil {
//		language = header.ClientLanguage
//	}
//	switch language {
//	case common.LanguageEN, common.LanguageID, common.LanguageMS:
//		break
//	default:
//		language = common.LanguageEN
//	}
//	return i18n.Tf(gi18n.WithLanguage(ctx, language), content, values)
//}

// GetRandomWithAll 包含上下限 [min, max]
func GetRandomWithAll(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// BuildVisitorNickname 匿名昵称
func BuildVisitorNickname(nickname string) string {
	return "Himi-" + MaskString(nickname, 1)
}

// BuildVisitorPortrait 匿名头像
func BuildVisitorPortrait() string {
	return consts.RoomAnonymousAvatar
}

// BuildMysteryManAvatarPortrait 匿名头像
func BuildMysteryManAvatarPortrait() string {
	return consts.MysteryManAvatar
}

func MaskString(input string, start int) string {
	if len(input) <= start {
		return input
	}
	inputRune := []rune(input)
	lenStart := len([]rune(input))
	switch {
	case lenStart <= 3:
		return string(inputRune[:start]) + strings.Repeat("*", lenStart-1)
	case 3 < lenStart && lenStart <= 5:
		return string(inputRune[:start+1]) + strings.Repeat("*", lenStart-2) + string(inputRune[lenStart+start-1:])
	case 5 < lenStart && lenStart <= 10:
		return string(inputRune[:start+2]) + strings.Repeat("*", lenStart-4) + string(inputRune[lenStart+start-2:])
	case lenStart > 10:
		return string(inputRune[:start+4]) + strings.Repeat("*", lenStart-8) + string(inputRune[lenStart+start-4:])
	default:
		return ""
	}
}

func InArray(array []int64, element int64) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}

func IsShowNightRoom(appChannel string, appVersionCode, privateStatus int) bool {
	if privateStatus != consts.PrivateStatusDiamond && privateStatus != consts.PrivateStatusPassword {
		return true
	}
	if strings.Contains(appChannel, "apple") {
		return false
	}
	if (appChannel == "himi_offline_android" && appVersionCode < 35) ||
		(appChannel == "himi_google" && appVersionCode < 20) {
		return false
	}

	if appChannel == "fuya_google" && appVersionCode < 5 {
		return false
	}
	return true
}

func MaxInt64(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//
//func GetWarnTypeName(ctx context.Context, warnType int) string {
//	switch warnType {
//	case 1:
//		return TranslateByKey(ctx, "liveWhileDriving")
//	case 2:
//		return TranslateByKey(ctx, "underageLive")
//	case 3:
//		return TranslateByKey(ctx, "play2AppsAtTheSameTime")
//	case 4:
//		return TranslateByKey(ctx, "notTheSameStreamer")
//	case 5:
//		return TranslateByKey(ctx, "sleepLive")
//	case 6:
//		return TranslateByKey(ctx, "blackSow")
//	case 7:
//		return TranslateByKey(ctx, "noisyOrUnnclear")
//	case 8:
//		return TranslateByKey(ctx, "idleForTooLong")
//	}
//	return ""
//}

func Sha1Hash(text string) string {
	// 创建SHA-1哈希对象
	sha1Hash := sha1.New()

	// 将文本转换为字节流并更新哈希对象
	sha1Hash.Write([]byte(text))

	// 获取SHA-1哈希值的十六进制表示
	encryptedText := fmt.Sprintf("%x", sha1Hash.Sum(nil))

	return encryptedText
}

// FormatFloat32 分子, denominator 分母
func FormatFloat32(ctx context.Context, numerator, denominator float32) float32 {

	if math.Abs(float64(denominator)) < 0.00001 || math.Abs(float64(numerator)) < 0.00001 {
		g.Log().Infof(ctx, "chyInfo denominator = %v , numerator = %v", denominator, numerator)
		return 0.0
	}
	result, err := strconv.ParseFloat(fmt.Sprintf("%.2f", numerator/denominator), 32)
	if err != nil {
		g.Log().Infof(ctx, "chyInfo FormatFloat32 err = %v", err)
		return 0.0
	}
	return float32(result)
}

//720 * 1280

// Top Left Bottom Right
// protrudeLayout 突出样式
var protrudeLayout = map[int][]string{
	1: {"0,0,1548,1080"},
	2: {"0,0,1094,541", "0,542,1094,1080"},
	3: {"0,0,961,720", "0,722,480,1080", "480,722,961,1080"},
	4: {"0,0,1080,720", "0,722,360,1080", "360,722,720,1080", "720,722,1080,1080"},
	5: {"0,0,1244,748", "0,751,311,1080", "311,751,622,1080", "622,751,933,1080", "933,748,1244,1080"},
	6: {"0,0,789,720", "0,722,394,1080", "394,722,792,1080", "789,394,1183,1080", "792,0,1183,360", "792,360,1183,720"},
	7: {"0,0,967,720", "0,722,322,1080", "322,722,645,1080", "645,722,967,1080", "969,0,1290,360", "969,360,1290,720", "969,722,1290,1080"},
	8: {"0,0,861,809", "0,809,288,1080", "288,809,576,1080", "576,809,864,1080", "861,0,1149,270", "861,270,1177,541", "861,541,1177,812", "861,812,1177,1080"},
	9: {"0,0,771,541", "0,541,385,809", "385,541,771,809", "0,809,385,1080", "385,809,771,1080", "771,0,1157,270", "771,270,1157,541", "771,541,1157,809", "771,809,1157,1080"},
}

// Top Left Bottom Right
// protrudeWidescreenLayout 突出宽屏样式
var protrudeWidescreenLayout = map[int][]string{
	1: {"0,0,1152,1080"},
	2: {"0,0,864,540", "0,540,864,1080"},
	3: {"0,0,864,720", "0,720,432,1080", "432,720,864,1080"},
	4: {"0,0,864,720", "0,720,288,1080", "288,720,576,1080", "576,720,864,1080"},
	5: {"0,0,864,748", "0,748,216,1080", "216,748,432,1080", "432,748,648,1080", "648,748,864,1080"},
	6: {"0,0,576,720", "0,720,288,1080", "288,720,576,1080", "576,720,864,1080", "576,0,864,360", "576,360,864,720"},
	7: {"0,0,648,720", "0,720,216,1080", "216,720,432,1080", "432,720,648,1080", "648,0,864,360", "648,360,864,720", "648,720,864,1080"},
	8: {"0,0,648,808", "0,808,216,1080", "216,808,432,1080", "432,808,648,1080", "648,0,864,270", "648,270,864,540", "648,540,864,810", "648,810,864,1080"},
	9: {"0,0,576,540", "0,540,288,812", "288,540,576,812", "0,812,288,1080", "288,812,576,1080", "576,0,864,270", "576,270,864,540", "576,540,864,810", "576,810,864,1080"},
}

//var protrudeLayout = map[int][]string{
//	1: {"0,0,1075,750"},
//	2: {"0,0,760,376", "0,377,760,750"},
//	3: {"0,0,668,500", "0,502,334,750", "334,502,668,750"},
//	4: {"0,0,750,500", "0,502,250,750", "250,502,500,750", "500,502,750,750"},
//	5: {"0,0,864,520", "0,522,216,750", "216,522,432,750", "432,522,648,750", "648,520,864,750"},
//	6: {"0,0,548,500", "0,502,274,750", "274,502,550,750", "548,274,822,750", "550,250,822,500", "550,502,822,750"},
//	7: {"0,0,672,500", "0,502,224,750", "224,502,448,750", "448,502,672,750", "673,0,896,250", "673,250,896,500", "673,502,896,750"},
//	8: {"0,0,598,562", "0,562,200,750", "200,562,400,750", "400,562,600,750", "598,0,798,188", "598,188,818,376", "598,376,818,564", "598,564,818,750"},
//	9: {"0,0,536,376", "0,376,268,562", "268,378,536,562", "536,376,804,562", "0,562,268,750", "268,562,536,750", "536,562,804,750", "536,0,804,188", "536,188,804,376"},
//}

// tileStyleLayout 平铺样式
var tileStyleLayout = map[int][]string{
	1: {"0,0,1075,750"},
	2: {"0,0,600,376", "0,376,600,750"},
	3: {"0,0,420,376", "0,378,420,750", "420,0,840,380"},
	4: {"0,0,420,376", "0,378,420,750", "420,0,840,380", "420,376,840,750"},
	5: {"0,0,410,250", "0,252,410,500", "0,502,410,750", "410,0,820,250", "410,250,820,500"},
	6: {"0,0,410,250", "0,252,410,500", "0,502,410,750", "410,0,820,250", "410,250,820,500", "410,500,820,750"},
	7: {"0,0,286,250", "0,252,288,500", "0,502,286,750", "286,0,572,250", "286,250,572,500", "286,500,572,750", "572,0,858,250"},
	8: {"0,0,286,250", "0,252,288,500", "0,502,286,750", "286,0,572,250", "286,250,572,500", "286,500,572,750", "572,0,858,250", "572,250,858,500"},
	9: {"0,0,286,250", "0,252,288,500", "0,502,286,750", "286,0,572,250", "286,250,572,500", "286,500,572,750", "572,0,858,250", "572,250,858,500", "572,500,858,750"},
}

// GetLayoutSet 获取多人房布局设置
// layoutSet 布局设置  1 突出 2 平铺 3 突出宽屏
// size 多人房人数
// index 多人房下标
func GetLayoutSet(layoutSet int64, size, index int) []string {
	if layoutSet == 1 {
		layoutSetStr := protrudeLayout[size][index]
		return strings.Split(layoutSetStr, ",")
	} else if layoutSet == 3 {
		layoutSetStr := protrudeWidescreenLayout[size][index]
		return strings.Split(layoutSetStr, ",")
	} else {
		layoutSetStr := tileStyleLayout[size][index]
		return strings.Split(layoutSetStr, ",")
	}
}

var CharmList = []string{"Pesona", "Charm", "Pesona cindo", "Pesona cosplay", "Pesona singing", "Pesona dance", "Pesona DJ"}

var CharmListV2 = []string{"Pesona", "Pesona cindo", "Pesona cosplay", "Pesona singing", "Pesona dance", "Pesona DJ"}

func IsCharmAnchor(ctx context.Context, label string) bool {
	if arrays.StringsContains(CharmList, label) != -1 {
		return true
	}
	return false
}

func GenerateOrderNo(userId int64) string {
	return fmt.Sprintf("himi-%v%v", userId, time.Now().UnixMilli())
}

func ConvertXenditBankName(bankName string) (bankCode string) {
	return "ID_BCA"
}

func GetPictureAndText(ctx context.Context, winType, multiple int, userId, diamonds int64) (picture string, text string) {
	switch winType {
	case 1:
		picture += "1.png"
		text = fmt.Sprintf("Selamat! Saya mendapat bigwin dalam Space War, mendapat hadiah %v kali lipat, %v diamond", multiple, diamonds)
		picture = fmt.Sprintf(picture+"?x-oss-process=image/auto-orient,1/resize,p_83/quality,q_90/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,text_%v,color_ffff7c,size_40,g_south,t_80,x_10,y_70", gbase64.EncodeString(gconv.String(diamonds)))
	case 2:
		picture += "2.png"
		text = fmt.Sprintf("Surprise！Saya mendapat megawin dalam Space War, mendapat bonus %v kali lipat, %v diamond", multiple, diamonds)
		picture = fmt.Sprintf(picture+"?x-oss-process=image/auto-orient,1/resize,p_83/quality,q_90/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,text_%v,color_ffff7c,size_40,g_south,t_80,x_10,y_85", gbase64.EncodeString(gconv.String(diamonds)))
	case 3:
		picture += "3.png"
		picture = fmt.Sprintf(picture+"?x-oss-process=image/auto-orient,1/resize,p_83/quality,q_90/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,text_%v,color_ffff7c,size_40,g_south,t_80,x_10,y_35", gbase64.EncodeString(gconv.String(diamonds)))
		text = fmt.Sprintf("How amazing！Saya mendapatkan turntable koin emas dalam Space War，dan mendapatkan %v diamond", diamonds)
	case 4:
		picture += "4.png"
		picture = fmt.Sprintf(picture+"?x-oss-process=image/auto-orient,1/resize,p_83/quality,q_90/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,text_%v,color_ffff7c,size_40,g_south,t_80,x_10,y_35", gbase64.EncodeString(gconv.String(diamonds)))
		text = fmt.Sprintf("Unbelievable! Saya mendapat jackpot dalam Space War, mendapat %v diamond", diamonds)
	}
	return
}

func GetGameAnchorNotice(ctx context.Context, winType, multiple int, userId, diamonds int64, nickname string) (noticeText string) {
	switch winType {
	case 1:
		return fmt.Sprintf(
			"Selamat kepada %v (%v) karena telah memenangkan bonus diamond %v kali lipat dan mendapatkan %v diamond", nickname, userId, multiple, diamonds)
	case 2:
		return fmt.Sprintf(
			"Selamat kepada %v (%v) karena telah memenangkan cash wheel dan mendapatkan %v diamond", nickname, userId, diamonds)
	case 3:
		return fmt.Sprintf(
			"Selamat kepada %v (%v) yang telah memenangkan jackpot dan mendapatkan %v diamond", nickname, userId, diamonds)
	case 4:
		return fmt.Sprintf("Selamat kepada %v (%v) karena telah mendapatkan 10 game gratis ", nickname, userId)
	}
	return
}

func GetDayDiamonds(giftRank int, gid, diamonds int64) int64 {
	if gid == 217 || gid == 219 || gid == 224 || gid == 226 || gid == 228 || gid == 231 || gid == 234 ||
		gid == 237 || gid == 242 || gid == 243 || gid == 248 || gid == 253 || gid == 256 || gid == 259 || gid == 260 { //10钻石
		if giftRank == 1 { //钻石奖励（钻石*n%）0.7% 徽章  *1 hari
			diamonds = diamonds * 7 / 1000
		} else if giftRank == 2 { //钻石奖励 0.5%  徽章  *1 hari
			diamonds = diamonds * 5 / 1000
		} else if giftRank == 3 { //钻石奖励 0.3%  徽章  *1 hari
			diamonds = diamonds * 3 / 1000
		} else {
			diamonds = 0
		}
	} else if gid == 217 || gid == 220 || gid == 222 || gid == 225 || gid == 229 || gid == 232 || gid == 235 ||
		gid == 238 || gid == 241 || gid == 244 || gid == 246 || gid == 252 || gid == 255 || gid == 258 || gid == 261 { //20钻石
		if giftRank == 1 { //钻石奖励（钻石*n%）0.9% 徽章  *1 hari
			diamonds = diamonds * 9 / 1000
		} else if giftRank == 2 { //钻石奖励 0.7%  徽章  *1 hari
			diamonds = diamonds * 7 / 1000
		} else if giftRank == 3 { //钻石奖励 0.5%  徽章  *1 hari
			diamonds = diamonds * 5 / 1000
		} else {
			diamonds = 0
		}
	} else if gid == 216 || gid == 221 || gid == 223 || gid == 227 || gid == 230 || gid == 233 || gid == 236 ||
		gid == 239 || gid == 240 || gid == 245 || gid == 247 || gid == 251 || gid == 254 || gid == 257 || gid == 262 { //30钻石
		if giftRank == 1 { //钻石奖励（钻石*n%）1.1% 徽章  *1 hari
			diamonds = diamonds * 11 / 1000
		} else if giftRank == 2 { //钻石奖励 0.9%  徽章  *1 hari
			diamonds = diamonds * 9 / 1000
		} else if giftRank == 3 { //钻石奖励 0.7%  徽章  *1 hari
			diamonds = diamonds * 7 / 1000
		} else {
			diamonds = 0
		}
	} else {
		diamonds = 0
	}
	return diamonds
}

func VersionOrdinal(ctx context.Context, version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			g.Log().Error(ctx, "判断版本号失败")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)

}

// 用户版本渠道是否符合新消息通知
func CheckNewMessage(ctx context.Context, appChannel string, version string) bool {
	if VersionOrdinal(ctx, "1.2.2") <= VersionOrdinal(ctx, version) && appChannel == "fuya_google" {
		return true
	}
	if VersionOrdinal(ctx, "1.0.5") <= VersionOrdinal(ctx, version) && appChannel == "fuya_apple" {
		return true
	}
	return false
}

package utility

import (
	"context"
	"embed"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/druidcaesa/ztool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"

	"fuya-ark/internal/consts"
)

var searcher *xdb.Searcher

func InitIp2region(ctx context.Context, fs embed.FS) {
	cBuff, err := fs.Open("resource/ip2region.xdb")
	//var err error
	//cBuff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		g.Log().Errorf(ctx, "failed to load content from `%s`: %s\n", fs, err)
		return
	}
	res, _ := io.ReadAll(cBuff)
	searcher, err = xdb.NewWithBuffer(res)
	if err != nil {
		g.Log().Errorf(ctx, "failed to create searcher with content: %s\n", err)
		return
	}
	g.Log().Info(ctx, "init ip2region success")
}

// SearchIp 根据ip转换地址 region 国家中文  isOfficial 是否官方isp
func SearchIp(ip string) (region string, isOfficial bool) {
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	ctx := gctx.New()
	if err != nil {
		g.Log().Errorf(ctx, "failed to SearchIP(%s): %s", ip, err)
		return
	}
	g.Log().Debugf(ctx, "ip2region---ip:%s region: %s, took: %s}", ip, region, time.Since(tStart))
	isOfficial = getIsOfficial(region)
	region = getRegionHead(region)
	if strings.Contains(region, consts.IndonesiaCn) ||
		strings.Contains(region, consts.MalaysiaCn) {
		return
	}
	var areaJson = ""
	areaJsonVar, err := g.Redis().Get(ctx, consts.IpCacheKey+ip)
	if len(areaJsonVar.String()) < 10 {
		//{"status":"success","country":"印度尼西亚","countryCode":"ID","region":"JK","regionName":"Jakarta","city":"雅加达","zip":"","lat":-6.2114,"lon":106.8446,"timezone":"Asia/Jakarta","isp":"Linknet-Fastnet ASN","org":"PT. Link Net Tbk","as":"AS23700 Linknet-Fastnet ASN","query":"149.113.254.104"}
		//不能改成https、只支持http
		areaJson, err = ztool.HttpUtils.Get("http://ip-api.com/json/" + ip + "?lang=zh-CN")
		if err != nil {
			g.Log().Errorf(ctx, "ip-api---ip:%s  err: %s}", ip, err)
		} else {
			if len(areaJson) > 0 && strings.Contains(areaJson, "success") {
				g.Redis().Set(ctx, consts.IpCacheKey+ip, areaJson)
			}
		}
		g.Log().Debugf(ctx, "ip-api---ip:%s region: %s, took: %s}", ip, areaJson, time.Since(tStart))
	} else {
		areaJson = areaJsonVar.String()
		g.Log().Debugf(ctx, "ip-api-oss---ip:%s region: %s, took: %s}", ip, areaJson, time.Since(tStart))
	}
	areaMap, ipAPiIsOfficial := parseIPApiResp(areaJson)
	if !isOfficial {
		isOfficial = ipAPiIsOfficial
	}
	region = areaMap["country"]
	if len(region) <= 0 {
		region = consts.USACn
	}
	return
}

func parseIPApiResp(areaJson string) (areaMap map[string]string, isOfficial bool) {
	areaMap = g.MapStrStr{}
	_ = json.Unmarshal([]byte(areaJson), &areaMap)
	isOfficial = getIsOfficial(areaJson)
	return
}

func getRegionHead(region string) (t string) {
	region = strings.TrimSpace(region)
	if len(region) <= 0 {
		return ""
	}
	firstSegEndIndex := strings.Index(region, "|")
	return gstr.SubStr(region, 0, firstSegEndIndex)
}

// getIsOfficial 判断region是否官方的ISP
func getIsOfficial(region string) (isOfficial bool) {
	if len(region) <= 0 {
		return false
	}
	region = strings.ToLower(region)
	isOfficial = strings.Contains(region, "谷歌")
	if !isOfficial {
		isOfficial = strings.Contains(region, "苹果")
	}
	if !isOfficial {
		isOfficial = strings.Contains(region, "脸书")
	}
	if !isOfficial {
		isOfficial = strings.Contains(region, "google")
	}
	if !isOfficial {
		isOfficial = strings.Contains(region, "facebook")
	}
	if !isOfficial {
		isOfficial = strings.Contains(region, "apple")
	}
	return
}

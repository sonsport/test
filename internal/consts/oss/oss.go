package oss

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/aliyun-sts-go-sdk/sts"
	"github.com/gogf/gf/v2/frame/g"

	"fuya-ark/internal/model"
)

var client *sts.Client

type ossCfg struct {
	Access struct {
		AppId  string `json:"id"`
		Secret string `json:"secret"`
	} `json:"userAccess"`
	RoleArn struct {
		Arn string `json:"arn"`
	} `json:"uploadRole"`
	Bucket      string `json:"publicBucket"`
	UserDir     string `json:"userDir"`
	OssEndpoint string `json:"oss_endpoint"`
	PublicHost  string `json:"publicHost"`
	StsEndpoint string `json:"sts_endpoint"`
}

var Cfg *ossCfg

func GetCfg() *ossCfg {
	return Cfg
}

func init() {
	cfgByte, err := g.Cfg().Get(context.Background(), "himi_aliyun.oss")
	if err != nil {
		g.Log().Errorf(context.Background(), "getUploadToken err:%v", err)
	}
	err = cfgByte.Scan(&Cfg)
	if err != nil {
		g.Log().Errorf(context.Background(), "getUploadToken err:%v", err)
	}
	g.Log().Info(context.TODO(), "init oss config success")
}

func GetUploadToken(ctx context.Context, uid int64) (*model.OssToken, error) {
	val, err := g.Cfg().Get(ctx, "himi_online_mode.value")
	strUid := strconv.Itoa(int(uid))
	if !val.Bool() {
		strUid += "dev"
	}
	client = sts.NewClient(Cfg.Access.AppId, Cfg.Access.Secret, Cfg.RoleArn.Arn, Cfg.Bucket+"."+strUid)
	uploadPath := Cfg.UserDir + strUid + "/upload"
	expire := 3600
	res, err := client.AssumeRole(uint(expire))
	if err != nil {
		g.Log().Errorf(ctx, "getUploadToken err:%v", err)
		return nil, err
	}

	return &model.OssToken{
		EndPoint: Cfg.OssEndpoint,
		Bucket:   Cfg.Bucket,
		Path:     uploadPath,
		Key:      res.Credentials.AccessKeyId,
		Secret:   res.Credentials.AccessKeySecret,
		Token:    res.Credentials.SecurityToken,
		Expire:   time.Now().Unix() + int64(expire),
	}, err
}

// GetFullUrl 获取正常图
func GetFullUrl(path string) string {
	if len(path) == 0 {
		return ""
	}
	path = strings.ReplaceAll(path, "https://himi-public.oss-ap-southeast-1.aliyuncs.com", "")
	if strings.Contains(path, "http") {
		return path
	}
	publicHost := Cfg.PublicHost
	return publicHost + path
}

// GetSmallUrl 获取小尺寸图，由阿里云OSS处理
func GetSmallUrl(path string, size int) string {
	if len(path) == 0 {
		return ""
	}
	if size == 0 {
		size = 200
	}
	if size != 0 && size != 600 && size != 100 && size != 300 {
		size = 200
	}
	path = strings.ReplaceAll(path, "https://himi-public.oss-ap-southeast-1.aliyuncs.com", "")
	if strings.Contains(path, "http") {
		return path
	}
	publicHost := Cfg.PublicHost
	return publicHost + path + "!zip-" + strconv.Itoa(size)
}

//// GetAliYunUploadKey 上传sts临时凭证
//func GetAliYunUploadKey(ctx context.Context, r service.RequestParam) (resp interface{}, gCode gcode.Code) {
//	accessKey := Cfg.Access.AppId
//	accessSecret := Cfg.Access.Secret
//	host := Cfg.PublicHost
//	policyBase64 := gbase64.EncodeString(`{"expiration":"2030-01-01T12:00:00.000Z","conditions":[["content-length-range",0,1048576000]]}`)
//	signature := encryption.HMACSHA1(accessSecret, policyBase64)
//	return g.Map{
//		"accessId":  accessKey,
//		"policy":    policyBase64,
//		"signature": signature,
//		"host":      host,
//	}, code.Succeed
//}

// DeleteObjects 删除文件
func DeleteObjects(paths ...string) {
	if len(paths) == 0 {
		return
	}
	for i := 0; i < len(paths); i++ {
		paths[i] = strings.ReplaceAll(paths[i], "https://himi-public.oss-ap-southeast-1.aliyuncs.com", "")
		paths[i] = strings.ReplaceAll(paths[i], Cfg.PublicHost, "")
	}
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(Cfg.OssEndpoint, Cfg.Access.AppId, Cfg.Access.Secret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// objectName表示删除OSS文件时需要指定包含文件后缀，不包含Bucket名称在内的完整路径，例如exampledir/exampleobject.txt。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	bucket, err := client.Bucket(Cfg.Bucket)
	// 删除单个文件。
	bucket.DeleteObjects(paths)
}

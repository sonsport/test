package encryption

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/text/gstr"
)

// Md5Encode 加密
func Md5Encode(s string) string {
	md5Str, _ := gmd5.Encrypt(s)
	return gstr.ToLower(md5Str)
}

// ComputeHmacSha256 HmacSha256加密
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
	//hex.EncodeToString(h.Sum(nil))
	//return base64.StdEncoding.EncodeToString([]byte(sha))
}

// HmacSha256 计算HmacSha256
// key 是加密所使用的key
// data 是加密的内容
func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))
	return mac.Sum(nil)
}

// HmacSha256ToHex 将加密后的二进制转16进制字符串
func HmacSha256ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}

// Sha1ToHex 将加密后的二进制转16进制字符串
func Sha1ToHex(data string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HmacSha256ToBase64 将加密后的二进制转Base64字符串
func HmacSha256ToBase64(key string, data string) string {
	return base64.StdEncoding.EncodeToString(HmacSha256(key, data))
}

/*
HMACSHA1
//  keyStr 密钥
//  value  消息内容
*/
func HMACSHA1(keyStr, value string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	//进行base64编码
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

/*
SHA1
*/
func SHA1(s string) string {
	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}

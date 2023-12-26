package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RSASign 私钥签名
func RSASign(data, keyBytes []byte) (string, error) {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha256.New()
	h.Write(data)
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// RSAVerify 公钥验证
func RSAVerify(originalData, signData, pubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	hash := sha256.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash.Sum(nil), sign)
}

func VerifyGoogleSign(data, sign, publicKey string) (bool, error) {
	decodePublic, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return false, err
	}

	pubInterface, err := x509.ParsePKIXPublicKey(decodePublic)
	if err != nil {
		return false, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	decodeSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	sh1 := sha1.New()
	sh1.Write([]byte(data))
	hashData := sh1.Sum(nil)

	result := rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashData, decodeSign)
	if result != nil {
		return false, err
	}
	return true, nil
}

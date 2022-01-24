package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/golang/glog"
)

//EncodingBase64 ... base64加密
func EncodingBase64(content []byte) string {
	return base64.StdEncoding.EncodeToString(content)
}

//DecodingBase64 ... base64解密
func DecodingBase64(str string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		glog.Errorf("DecodeString failed,err:%+v", err)
		return "", err
	}
	return string(decoded), nil
}

//EncodingPassword ... 用户密码加密
func EncodingPassword(str string) string {
	salt := "qpnmPX7PrQWJy88zjA7tbmbhf2WkxFEM"
	h := md5.New()
	h.Write([]byte(salt + str))
	return hex.EncodeToString(h.Sum(nil))
}

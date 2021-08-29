package md5

import (
	"api-go/lib/config"
	"crypto/md5"
	"encoding/hex"
)

// MD5 加密
func ToMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// 加盐 MD5
func ToMD5Salt(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	m.Write([]byte(config.AppConfig.MD5Salt))
	return hex.EncodeToString(m.Sum(nil))
}

package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.Sum([]byte(data))

	return hex.EncodeToString(h[:])
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// ValidatePasswd 校验用户密码
func ValidatePasswd(plainPassword, salt, passwd string) bool {
	return Md5Encode(plainPassword+salt) == passwd
}

// MakePasswd 生成用户密码
func MakePasswd(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}

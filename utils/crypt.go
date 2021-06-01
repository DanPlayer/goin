package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5hex(str string) string {
	ret := md5.Sum([]byte(str))
	return hex.EncodeToString(ret[:])
}

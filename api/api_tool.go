package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func MakeSign(method, appId, appSecret, version string, timestamp int64) string {
	value := fmt.Sprintf("%v?appId=%v&timestamp=%v&version=%v%v", method, appId, timestamp, version, appSecret)
	sign := Md5(value)
	return sign
}

func Md5(str string) string {
	m := md5.New()
	_, _ = io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}

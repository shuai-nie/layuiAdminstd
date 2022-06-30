package util
// 组件
import (
	"crypto/md5"
	"encoding/hex"
)
// 加密
func EncodeMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}



package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"hash/crc32"
	"math/rand"
	"strings"
	"time"
)

func Md5(encryptString string) string {
	md5Object := md5.New()
	md5Object.Write([]byte(encryptString + beego.AppConfig.String("md5Code")))
	return hex.EncodeToString(md5Object.Sum(nil))
}

func DateFormat(currentTimeStamp int64) string {
	unixTime := time.Unix(currentTimeStamp, 0)
	return unixTime.Format("2006-01-02")
}


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandSeq 随机字符串
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Trim 移除两端空格和制表符
func Trim(str string) string {
	str = strings.Trim(str, "\n")
	return strings.Trim(str, " ")
}

// HashStr2Int 将字符串转换成一串唯一编码
func HashStr2Int(s string) (code int) {
	code = int(crc32.ChecksumIEEE([]byte(s)))
	if code >= 0 {
		return
	}
	if -code >= 0 {
		return -code
	}
	// v == MinInt
	return
}


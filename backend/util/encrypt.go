package util

import (
	"crypto/md5"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/scrypt"
	"math/rand"
	"strconv"
	"time"
)

func Encrypt(password string) string {
	salt, err := beego.AppConfig.String("salt")
	if err != nil {
		panic(err)
	}
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	str := fmt.Sprintf("%x", dk)
	return str
}
func SelfMd5(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}

/*
RandStringWithTime
带有时间的随机值
*/
func RandStringWithTime() string {
	now := time.Now()
	str := fmt.Sprintf("%2d%2d%7d", now.Month(), now.Day(), now.Nanosecond())
	num := rand.Int()
	itoa := strconv.Itoa(num)
	str = Encrypt(str + itoa)
	return str
}

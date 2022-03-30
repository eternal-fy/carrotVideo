package nosql

import (
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/garyburd/redigo/redis"
)

func GetRedisConnect() (redis.Conn, error) {
	host, _ := beego.AppConfig.String("redis-host")
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return nil, err
	}
	redisPassword, _ := beego.AppConfig.String("redis-password")

	flag, err := conn.Do("AUTH", redisPassword)
	if flag != "OK" {
		fmt.Println("redis-password error")
		return nil, errors.New("redis-password error")
	}

	fmt.Println("redis conn success", flag)
	return conn, nil
}

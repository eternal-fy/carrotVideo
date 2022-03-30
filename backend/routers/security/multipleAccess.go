package security

import (
	"backend/dao/nosql"
	_ "backend/mylog"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/garyburd/redigo/redis"
)

/*
	为了防止同一个ip在短时间内频繁访问导致服务器崩溃
	通过redis缓存ip，只允许5分钟内最多1000次访问
	如果超过则不进行服务，需要等存放ip的key过期
	未超过则进行访问次数自增
*/
const (
	expireSeconds = 5 * 60
	limitTimes    = 1000
	stopRunMeg    = "访问次数过于频繁，请稍后再试"
)

func MultipleAccess() beego.FilterFunc {
	return func(ctx *context.Context) {
		connect, err := nosql.GetRedisConnect()
		if err != nil {
			fmt.Println("redis connect fail")
		}
		ip := ctx.Request.RemoteAddr
		flag, err := CheckOverWhelm(connect, ip)
		if err != nil {
			panic(err)
		}
		if !flag {
			panic(stopRunMeg)
		}
	}
}

func CheckOverWhelm(conn redis.Conn, ip string) (bool, error) {
	count, err := redis.Int(conn.Do("incr", ip))
	if err != nil {
		return false, err
	}
	if count == 1 {
		conn.Do("expire", expireSeconds)
	}
	if count > limitTimes {
		return false, nil
	}
	return true, nil

}

package test

import (
	"backend/dao/nosql"
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	connect, err := nosql.GetRedisConnect()
	if err != nil {
		fmt.Println("redis 链接成功----")
	}
	do, _ := connect.Do("ping")
	fmt.Println(do)
}

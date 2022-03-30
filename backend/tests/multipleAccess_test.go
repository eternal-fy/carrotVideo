package test

import (
	"backend/dao/nosql"
	"backend/routers/security"
	"testing"
)

func TestMultipleAccess(t *testing.T) {
	connect, _ := nosql.GetRedisConnect()

	whelm, _ := security.CheckOverWhelm(connect, "106.12.174.72")
	print(whelm)

}

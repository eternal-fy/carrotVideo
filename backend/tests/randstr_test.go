package test

import (
	"backend/util"
	"testing"
)

func TestRansStr(t *testing.T) {
	time := util.RandStringWithTime()
	println(time)
}

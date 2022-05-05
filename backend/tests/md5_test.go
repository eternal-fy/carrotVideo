package test

import (
	"backend/util"
	"testing"
)

func TestMd5(t *testing.T) {
	println("----")
	md5 := util.SelfMd5("199935")
	println(md5)
}

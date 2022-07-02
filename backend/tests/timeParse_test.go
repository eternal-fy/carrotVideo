package test

import (
	"backend/util"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	now := time.Now()
	parseTime := util.ParseTime(now)
	println(parseTime)
}

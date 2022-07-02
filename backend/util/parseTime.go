package util

import (
	"fmt"
	"time"
)

func ParseTime(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	timeString := fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, min, sec)
	return timeString
}

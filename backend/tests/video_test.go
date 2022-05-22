package test

import (
	"backend/dao/sql/videoDao"
	"fmt"
	"testing"
)

func TestVideo(t *testing.T) {
	videoDao.Before()
}

func TestSelectVideo(t *testing.T) {
	username := ""
	page := 1
	size := 5
	byUsername, err := videoDao.SelectVideoByUsername(username, "video", page, size)
	if err != nil {
		panic(err)
	}
	for item := range byUsername {
		fmt.Printf("%v", item)

	}
}

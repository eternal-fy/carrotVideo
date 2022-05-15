package test

import (
	"backend/dao/bosService"
	"io/ioutil"
	"os"
	"testing"
)

func TestBosServiceUpload(t *testing.T) {
	open, err := os.Open("D:\\GoCode\\MyGoer\\carrotVideo\\front\\src\\assets\\face.jpg")
	if err != nil {
		panic("open fail")
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		panic("open fail")
	}
	bosService.BosUpload(all, "ld", "ldimg")
}
func TestBosServiceGetUrl(t *testing.T) {
	download := bosService.BosGetUrl("star")
	println(download)
}

func TestBosServiceGetObject(t *testing.T) {
	object := bosService.BosGetObject("firstObject")
	stream := object.Body
	defer stream.Close()
	all, err := ioutil.ReadAll(stream)
	if err != nil {
		panic(err)
	}
	println(string(all))
}

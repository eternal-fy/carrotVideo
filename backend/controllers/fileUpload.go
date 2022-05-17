package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) UploadFiles() {
	_, header, _ := c.GetFile("file")
	open, err := header.Open()
	if err != nil {
		panic("文件打开失败")
	}
	bytes, err := ioutil.ReadAll(open)
	println(bytes)
	if err != nil {
		panic("文件读取失败")
	}
	println(header)
}

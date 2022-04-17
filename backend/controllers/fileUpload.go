package controllers

import beego "github.com/beego/beego/v2/server/web"

type FileController struct {
	beego.Controller
}

func (f *FileController) UploadFiles() {
	println(f.Ctx)
}

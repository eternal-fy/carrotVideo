package controllers

import (
	"backend/dao/bosService"
	"backend/dao/sql/videoDao"
	"backend/models/videoInfo"
	"backend/util"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"log"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) UploadFiles() {
	username := c.Ctx.GetCookie("name")
	username = util.StringDecode(username)
	_, fileHead, _ := c.GetFile("videoFile")
	_, imgHead, _ := c.GetFile("videoImg")
	title := c.GetString("title")
	typeName := c.GetString("type")
	des := c.GetString("description")

	open, err := fileHead.Open()
	open1, err1 := imgHead.Open()
	if err != nil || err1 != nil {
		panic("文件打开失败")
	}

	fileBytes, err := ioutil.ReadAll(open)
	imgBytes, err := ioutil.ReadAll(open1)

	if err != nil {
		panic("文件读取失败")
	}
	go func() {
		fileId := util.RandStringWithTime()
		imgId := util.RandStringWithTime()

		video := &videoInfo.Video{}
		video.Username = username
		video.Type = typeName
		video.Title = title
		video.VideoId = fileId
		video.ImageId = imgId
		video.Description = des
		bosService.BosUpload(fileBytes, username, fileId)
		bosService.BosUpload(imgBytes, username, imgId)
		videoDao.CreateVideo(video)
	}()

	log.Println("上传成功！")
	response := &ResponseData{}
	response.Code = SUCCESS
	response.Msg = "上传成功！"
	c.Data["json"] = &response
	c.ServeJSON()
}

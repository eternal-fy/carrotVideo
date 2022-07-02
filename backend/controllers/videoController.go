package controllers

import (
	"backend/dao/bosService"
	"backend/dao/sql/userDao"
	"backend/dao/sql/videoDao"
	"backend/models/videoInfo"
	"backend/util"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	beego.Controller
}
type VideoInfoDTO struct {
	VideoModel    videoInfo.Video
	AuthorImg     string
	VideoResource string
	ImgResource   string
}
type SendData struct {
	VideoType string
	Username  string
	Page      int
}

//GetVideoInfos 获取视频信息
func (c *VideoController) GetVideoInfos() {
	var data SendData
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if data.VideoType == "index" {
		data.VideoType = ""
	}
	size := 8

	usernameCookie := c.Ctx.GetCookie("name")
	username := data.Username
	if username == usernameCookie {
		data.Username = util.StringDecode(username)
	}

	videos, err := videoDao.SelectVideoByUsername(data.Username, data.VideoType, data.Page, size)
	if err != nil {
		panic(err)
	}
	infoDTO := make([]VideoInfoDTO, len(videos))
	for index, _ := range infoDTO {
		infoDTO[index].VideoModel = videos[index]
		videoResource := bosService.BosGetUrl(videos[index].VideoId)
		imgResource := bosService.BosGetUrl(videos[index].ImageId)

		infoDTO[index].AuthorImg = userDao.GetProfileImgUrl(videos[index].Username)
		infoDTO[index].VideoResource = videoResource
		infoDTO[index].ImgResource = imgResource
	}
	c.Data["json"] = infoDTO
	c.ServeJSON()
}

//PostMsg 添加评论
func (c *VideoController) PostMsg() {
	username := c.Ctx.GetCookie("name")
	username = util.StringDecode(username)
	videoid := c.GetString("videoid")
	des := c.GetString("desc")
	response := &ResponseData{}

	err := videoDao.PostMsg(videoid, username, des)
	response.Code = SUCCESS
	if err != nil {
		response.Code = FAIL
		response.Msg = err.Error()
	}
	c.Data["json"] = response
	c.ServeJSON()
}

//PostMsg 添加评论
func (c *VideoController) GetMsg() {
	videoid := c.GetString("videoid")
	response := &ResponseData{}

	list, err := videoDao.GetMsg(videoid)
	response.Code = SUCCESS
	if err != nil {
		response.Code = FAIL
		response.Msg = err.Error()
		return
	}
	response.TransData = list
	c.Data["json"] = response
	c.ServeJSON()
}

//Star 点赞
func (c *VideoController) Star() {
	videoid := c.GetString("videoid")
	author := c.GetString("author")
	username := c.Ctx.GetCookie("name")
	username = util.StringDecode(username)

	response := &ResponseData{}
	flag := userDao.WhetherStar(videoid, username, author)
	var err error
	if !flag {
		err = userDao.Star(videoid, username, author)
	} else {
		err = userDao.UnStar(videoid, username)
	}

	response.Code = SUCCESS
	if err != nil {
		response.Code = FAIL
		response.Msg = err.Error()
		return
	}
	c.Data["json"] = response
	c.ServeJSON()
}

//GetStar 获取点赞情况
func (c *VideoController) GetStar() {
	videoid := c.GetString("videoid")
	username := c.Ctx.GetCookie("name")
	username = util.StringDecode(username)

	response := &ResponseData{}

	count := userDao.GetStar(videoid, username)
	response.TransData = count
	response.Code = SUCCESS
	c.Data["json"] = response
	c.ServeJSON()
}

//GetStar 获取点赞情况
func (c *VideoController) GetStarCount() {
	videoid := c.GetString("videoid")

	response := &ResponseData{}

	count := userDao.GetStarCount(videoid)
	response.TransData = count
	response.Code = SUCCESS
	c.Data["json"] = response
	c.ServeJSON()
}

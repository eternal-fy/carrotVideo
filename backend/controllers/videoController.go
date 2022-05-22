package controllers

import (
	"backend/dao/bosService"
	"backend/dao/sql/videoDao"
	"backend/models/videoInfo"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	beego.Controller
}
type VideoInfoDTO struct {
	VideoModel    videoInfo.Video
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
		data.VideoType = "video"
	}
	size := 13

	videos, err := videoDao.SelectVideoByUsername(data.Username, data.VideoType, data.Page, size)
	if err != nil {
		panic(err)
	}
	infoDTO := make([]VideoInfoDTO, size)
	for index, _ := range infoDTO {
		infoDTO[index].VideoModel = videos[index]
		videoResource := bosService.BosGetUrl(videos[index].VideoId)
		imgResource := bosService.BosGetUrl(videos[index].ImageId)
		infoDTO[index].VideoResource = videoResource
		infoDTO[index].ImgResource = imgResource
	}
	c.Data["json"] = infoDTO
	c.ServeJSON()
}

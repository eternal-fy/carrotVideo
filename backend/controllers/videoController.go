package controllers

import (
	"backend/dao/bosService"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	beego.Controller
}
type VideoInfoDTO struct {
	VideoResource string
	ImgResource   string
	InfoResource  string
}
type SendData struct {
	VideoType string
}

//QQ联合登陆验证
func (c *VideoController) GetVideoInfos() {
	var data SendData
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	if data.VideoType != "index" {
		panic("传递类型错误")
	}
	videoResource := bosService.BosGetUrl("star")
	imgResource := bosService.BosGetUrl("star_png")
	infoResource := bosService.GetStringByName("firstObject")
	infoDTO := make([]VideoInfoDTO, 10)
	for i := 0; i < 10; i++ {
		infoDTO[i] = VideoInfoDTO{videoResource, imgResource, infoResource}
	}
	c.Data["json"] = infoDTO
	c.ServeJSON()
}

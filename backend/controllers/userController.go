package controllers

import (
	"backend/dao/bosService"
	dao "backend/dao/sql/userDao"
	"backend/models/userInfo"
	"backend/util"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"log"
)

type UserController struct {
	beego.Controller
}
type UserInfo struct {
	PersonalInfo userInfo.User
}

func (c *UserController) SaveInformation() {
	validated := CheckUser(c)
	username := c.Ctx.GetCookie("name")
	response := &ResponseData{}
	defer LastSolve(c, response)
	if !validated {
		response.Code = FAIL
		response.Msg = "密码错误！"
		return
	}
	var user UserInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic("数据异常")
	}

	dao.SaveUser(user.PersonalInfo, util.StringDecode(username))
	response.Code = SUCCESS

}

//postprofileimg
func (c *UserController) Postprofileimg() {
	username := c.Ctx.GetCookie("name")
	username = util.StringDecode(username)
	_, imgHead, _ := c.GetFile("profileImg")
	open1, err := imgHead.Open()
	if err != nil || err != nil {
		panic("文件打开失败")
	}

	imgBytes, err := ioutil.ReadAll(open1)

	if err != nil {
		panic("文件读取失败")
	}
	imgId := util.RandStringWithTime()
	err = dao.ChangeProfileImgUrl(username, imgId)
	if err != nil {
		panic(err)
	}
	bosService.BosUpload(imgBytes, username, imgId)

	log.Println("上传成功！")
	response := &ResponseData{}
	response.Code = SUCCESS
	response.Msg = "上传成功！"
	c.Data["json"] = &response
	c.ServeJSON()
}

// GetInformation 获取用户的基本信息
func (c *UserController) GetInformation() {
	cookieusername := c.Ctx.GetCookie("name")
	username := c.GetString("username")

	response := &ResponseData{}
	if cookieusername == username {
		validated := CheckUser(c)
		if !validated {
			response.Code = FAIL
			response.Msg = "密码错误！"
			return
		}
		username = util.StringDecode(username)
	}

	defer LastSolve(c, response)

	user := dao.GetUser(username)
	user.Password = ""
	response.TransData = user
	response.Code = SUCCESS
}

// GetUserImgUrl 获取用户的头像
func (c *UserController) GetUserImgUrl() {
	username := c.GetString("username")
	cookieusername := c.Ctx.GetCookie("name")

	response := &ResponseData{}
	if cookieusername == username {
		validated := CheckUser(c)
		if !validated {
			response.Code = FAIL
			response.Msg = "密码错误！"
			return
		}
		username = util.StringDecode(username)
	}
	defer LastSolve(c, response)
	url := dao.GetProfileImgUrl(username)
	response.TransData = url
	response.Code = SUCCESS
}

//getusernickname
func (c *UserController) GetUserNickname() {
	username := c.GetString("username")
	response := &ResponseData{}
	defer LastSolve(c, response)

	nickName := dao.GetUser(username).Name
	response.TransData = nickName
	response.Code = SUCCESS
}

//CheckUser cookie值与session值是否对应
func CheckUser(c *UserController) bool {
	username := c.Ctx.GetCookie("name")
	randSequence := c.Ctx.GetCookie("rand-sequence")
	session := c.GetSession(username)
	if session != randSequence {
		return false
	}
	return true
}
func LastSolve(this *UserController, response *ResponseData) {
	this.Data["json"] = response
	this.ServeJSON()
}

package controllers

import (
	dao "backend/dao/sql/user"
	"backend/models/userInfo"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}
type UserInfo struct {
	PersonalInfo userInfo.User
}

func (c *UserController) SaveInformation() {
	validated := CheckUser(c)
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
	dao.SaveUser(user.PersonalInfo, "ld")
	response.Code = SUCCESS

}

// GetInformation 获取用户的基本信息
func (c *UserController) GetInformation() {
	username := c.Ctx.GetCookie("name")
	validated := CheckUser(c)
	response := &ResponseData{}
	defer LastSolve(c, response)
	if !validated {
		response.Code = FAIL
		response.Msg = "密码错误！"
		return
	}
	user := dao.GetUser(username)
	user.Password = ""
	response.TransData = user
	response.Code = SUCCESS
}

// GetUserImgUrl 获取用户的头像
func (c *UserController) GetUserImgUrl() {
	username := c.Ctx.GetCookie("name")
	validated := CheckUser(c)
	response := &ResponseData{}
	defer LastSolve(c, response)
	if !validated {
		response.Code = FAIL
		response.Msg = "密码错误！"
		return
	}
	url := dao.GetProfileImgUrl(username)
	response.TransData = url
	response.Code = SUCCESS
}

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

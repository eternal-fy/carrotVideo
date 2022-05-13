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
	username := c.Ctx.GetCookie("name")
	randSequence := c.Ctx.GetCookie("rand-sequence")
	session := c.GetSession(username)
	if randSequence == session {
		println("success")
	}
	var user UserInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic("数据异常")
	}
	dao.SaveUser(user.PersonalInfo, "ld")
	response := &ResponseData{}
	response.Code = SUCCESS
	c.Data["json"] = &response
	c.ServeJSON()

}

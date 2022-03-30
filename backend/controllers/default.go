package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Result struct {
	Name string
}

func (c *MainController) Post() {
	all, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	result := &Result{}
	json.Unmarshal(all, result)
	fmt.Println(result)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

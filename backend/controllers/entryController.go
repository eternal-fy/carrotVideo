package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type EntryController struct {
	beego.Controller
}

//QQ联合登陆验证
func (c *EntryController) Auth() {
	uri := c.Ctx.Request.RequestURI
	println(uri)
	c.Redirect("http://www.eternalfy.site/main-page/main-context/index", http.StatusFound)
	/*	params := url.Values{}
		appId := "102004026"
		params.Add("response_type", "code")
		params.Add("client_id", appId)
		params.Add("state", "test")
		redirectURI := "/login"
		str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
		loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)
		println(loginURL)*/
	//c.Redirect(loginURL, http.StatusFound)
}

func (c *EntryController) Login() {
	c.Ctx.WriteString("login success")
}

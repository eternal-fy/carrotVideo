package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EntryController struct {
	beego.Controller
}

//QQ联合登陆验证
func (c *EntryController) Auth() {
	params := url.Values{}
	AppId := "102005140"
	params.Add("response_type", "code")
	params.Add("client_id", AppId)
	params.Add("state", "test")
	redirectURI := "http%3A%2F%2Feternalfy.site%2Fapi%2Fentry%2Fgetcode"
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)
	c.Ctx.ResponseWriter.Write([]byte(loginURL))
}

func (c *EntryController) GetCode() {
	code := c.Ctx.Input.Query("code")
	params := url.Values{}
	AppId := "102005140"
	Secret := "DE7mIahZGj03uWwE"
	params.Add("client_id", AppId)
	params.Add("client_secret", Secret)
	params.Add("code", code)
	redirectURI := "http%3A%2F%2Feternalfy.site%2Fapi%2Fentry%2Fgettoken"
	loginURL := fmt.Sprintf("%s&%s&redirect_uri=%s", "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code", params.Encode(), redirectURI)
	response, _ := http.Get(loginURL)
	defer response.Body.Close()
	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)
	println(body)
}
func (c *EntryController) GetToken() {
	token := c.Ctx.Input.Query("access_token")
	println(token)

}

func (c *EntryController) Login() {
	c.Ctx.WriteString("login success")
}

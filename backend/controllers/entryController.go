package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EntryController struct {
	beego.Controller
}

// LoginInfo 获取token
type LoginInfo struct {
	UserName string
	Password string
}

// Access 获取token
type Access struct {
	Access_token  string
	Expires_in    string
	Refresh_token string
}

//ClientId 获取openID
type ClientId struct {
	Client_id string
	Openid    string
}

// Auth QQ联合登陆验证
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
	params.Add("fmt", "json")
	redirectURI := "http%3A%2F%2Feternalfy.site%2Fmain-page%2Fmain-context%2Findex"
	loginURL := fmt.Sprintf("%s&%s&redirect_uri=%s", "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code", params.Encode(), redirectURI)
	return
	response, err := http.Get(loginURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bytes, _ := ioutil.ReadAll(response.Body)
	var access Access
	json.Unmarshal(bytes, &access)

	//获取client_id和openid
	paramsToken := url.Values{}
	paramsToken.Add("access_token", access.Access_token)
	paramsToken.Add("fmt", "json")
	tokenUri := fmt.Sprintf("%s%s", "https://graph.qq.com/oauth2.0/me?", paramsToken.Encode())
	get, err := http.Get(tokenUri)
	if err != nil {
		println(err)
	}
	defer get.Body.Close()
	clientBytes, _ := ioutil.ReadAll(get.Body)
	var clientId ClientId
	json.Unmarshal(clientBytes, &clientId)
	infoUri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", access.Access_token, AppId, clientId.Openid)
	resp, err := http.Get(infoUri)
	all, _ := ioutil.ReadAll(resp.Body)
	ctx := c.Ctx
	ctx.ResponseWriter.Write(all)
	localurl := "/main-page/main-context/index"
	http.Redirect(ctx.ResponseWriter, ctx.Request, localurl, http.StatusFound)
}
func (c *EntryController) GetToken() {
	token := c.Ctx.Input.Query("access_token")
	println(token)

}

func (c *EntryController) Login() {
	var result LoginInfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &result)
	if err != nil {
		panic(err)
	}
	c.Ctx.WriteString("login success")
}

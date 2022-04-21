package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
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
	redirectURI := "http%3A%2F%2Feternalfy.site%2Fapi%2Fentry%2Fgettoken"
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)
	c.Ctx.ResponseWriter.Write([]byte(loginURL))
}

/*func (c *EntryController) GetToken() {
	code := c.FormValue("code")
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", AppId)
	params.Add("client_secret", AppKey)
	params.Add("code", code)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/token", str)

	response, err := http.Get(loginURL)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)

	resultMap := convertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	GetOpenId(info, w)

}*/

func (c *EntryController) Login() {
	c.Ctx.WriteString("login success")
}

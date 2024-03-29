package controllers

import (
	"backend/dao/bosService"
	. "backend/dao/sql/userDao"
	"backend/models/userInfo"
	"backend/util"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	SUCCESS = "9999"
	FAIL    = "0000"
	MAXAGE  = 1000 * 60 * 60 * 24 * 7
)

type EntryController struct {
	beego.Controller
}

// LoginInfo 获取token
type ResponseData struct {
	Code      string      //返回Code,其中9999为成功，0000为失败，其余码为
	Msg       string      //返回的消息
	TransData interface{} //需要传输的数据
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

//Qresult 获取QQ三方登录的信息
type Qresult struct {
	Nickname       string
	Gender         string
	Figureurl_qq_2 string
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
	loginURL := fmt.Sprintf("%s&%s&redirect_uri=%s。", "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code", params.Encode(), redirectURI)
	ctx := c.Ctx
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
	user := GetUserByAppid(clientId.Openid)

	infoUri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", access.Access_token, AppId, clientId.Openid)
	resp, err := http.Get(infoUri)
	all, _ := ioutil.ReadAll(resp.Body)

	var qresult Qresult
	err = json.Unmarshal(all, &qresult)
	if err != nil {
		panic(err)
	}

	var newRecord userInfo.User
	var username string
	if user.ID == 0 {
		username = util.RandStringWithTime()
		newRecord.Username = username
		newRecord.Appid = clientId.Openid
		var genderNum int32
		genderNum = 0
		if qresult.Gender == "男" {
			genderNum = 1
		}
		newRecord.Gender = genderNum
		newRecord.Name = qresult.Nickname
		newRecord.Profileimgname = username + "img"
		newRecord.Password = util.RandStringWithTime()
		CreateUser(newRecord)
		bosService.BosUploadByUrl(qresult.Figureurl_qq_2, username, newRecord.Profileimgname)
	} else {
		username = user.Username
	}
	localurl := "http://eternalfy.site/main-page/main-context/index"
	randSequence := util.RandStringWithTime()
	c.Ctx.SetCookie("name", util.StringEncode(username), "/")
	c.Ctx.SetCookie("rand-sequence", randSequence, "/")
	c.SetSession(util.StringEncode(username), randSequence)
	http.Redirect(ctx.ResponseWriter, ctx.Request, localurl, http.StatusFound)

}

func (c *EntryController) Register() {
	var user userInfo.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic(err)
	}

	response := &ResponseData{}
	response.Code = SUCCESS
	response.Msg = "注册成功！"

	defer func(this *EntryController, response *ResponseData) {
		this.Data["json"] = response
		this.ServeJSON()
	}(c, response)

	username := user.Username
	password := user.Password
	validated := c.DataValidated(username, password, response)
	if !validated {
		return
	}

	validate := UserNameValidated(username)
	if !validate {
		response.Code = FAIL
		response.Msg = "用户名已存在！"
	} else {
		user.Password = util.Encrypt(user.Password)
		err := CreateUser(user)
		if err != nil {
			response.Code = FAIL
			response.Msg = err.Error()
		}
	}

}

func (c *EntryController) Login() {
	var user userInfo.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic(err)
	}
	response := &ResponseData{}
	response.Code = SUCCESS
	response.Msg = "登陆成功！"
	defer func(this *EntryController, response *ResponseData) {
		this.Data["json"] = response
		this.ServeJSON()
	}(c, response)

	username := user.Username
	password := user.Password
	validated := c.DataValidated(username, password, response)
	if !validated {
		return
	}

	tPassword := GetUser(username).Password
	encrypt := util.Encrypt(password)
	if tPassword != encrypt {
		response.Code = FAIL
		response.Msg = "密码错误"
		return
	}
	randSequence := util.RandStringWithTime()
	c.Ctx.SetCookie("name", util.StringEncode(username), "/")
	c.Ctx.SetCookie("rand-sequence", randSequence, "/")
	c.SetSession(util.StringEncode(username), randSequence)
}
func (c *EntryController) CheckUsername() {
	var user userInfo.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		panic(err)
	}
	validate := UserNameValidated(user.Username)
	response := &ResponseData{}
	response.Code = FAIL
	if validate {
		response.Code = SUCCESS
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

/*
DataValidated
判断用户名和密码数据的有效性
*/
func (c *EntryController) DataValidated(username, password string, response *ResponseData) bool {
	usernameLen := len(username)
	passwordLen := len(password)
	if usernameLen < 2 || usernameLen > 16 || passwordLen < 6 || passwordLen > 16 {
		response.Code = FAIL
		response.Msg = "用户名或者密码长度异常！"
		return false
	}
	return true
}

/*
GetLoginStatus
判断cookie是否为登录状态，存在session中

func (c *EntryController) GetLoginStatus() {
	username := c.Ctx.GetCookie("name")
	randSequence := c.Ctx.GetCookie("rand-sequence")
	session := c.GetSession(username)
	response := &ResponseData{}
	response.Code = SUCCESS
	response.Msg = "成功登录！"
	if session != randSequence {
		response.Code = FAIL
		response.Msg = "登录失败！"
	}
	c.Data["json"] = response
	c.ServeJSON()
}*/

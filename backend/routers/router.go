package routers

import (
	. "backend/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &MainController{})
	beego.AutoRouter(&EntryController{})
}

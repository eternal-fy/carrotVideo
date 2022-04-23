package routers

import (
	. "backend/controllers"
	"backend/routers/security"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, security.MultipleAccess())
	beego.Router("/", &MainController{})
	beego.AutoRouter(&EntryController{})
	beego.AutoRouter(&FileController{})
	beego.AutoRouter(&VideoController{})

}

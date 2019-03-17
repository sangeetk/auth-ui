package routers

import (
	"git.urantiatech.com/auth/login/controllers"
	"github.com/urantiatech/beego"
)

func init() {
	beego.Router("/auth/register", &controllers.RegisterController{}, "get:RegisterForm")
	beego.Router("/auth/register", &controllers.RegisterController{}, "post:RegisterUser")

	beego.Router("/auth/confirm", &controllers.ConfirmController{}, "get:ConfirmUser")

	beego.Router("/auth/forgot", &controllers.ForgotController{}, "get:ForgotForm")
	beego.Router("/auth/forgot", &controllers.ForgotController{}, "post:ResetLink")

	beego.Router("/auth/reset", &controllers.ResetController{}, "get:ResetForm")
	beego.Router("/auth/reset", &controllers.ResetController{}, "post:ResetPassword")

	beego.Router("/auth/login", &controllers.LoginController{}, "get:LoginForm")
	beego.Router("/auth/login", &controllers.LoginController{}, "post:LoginUser")

	beego.Router("/auth/logout", &controllers.LogoutController{}, "get:LogoutUser")

	beego.ErrorController(&controllers.ErrorController{})

	beego.SetStaticPath("/auth/css", "static/css")
	beego.SetStaticPath("/auth/js", "static/js")
	beego.SetStaticPath("/auth/img", "static/img")
	beego.SetStaticPath("/auth/font", "static/font")
}

package routers

import (
	"git.urantiatech.com/auth/login/controllers"
	"github.com/urantiatech/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{}, "get:RegisterForm")
	beego.Router("/register", &controllers.RegisterController{}, "post:RegisterUser")

	beego.Router("/confirm", &controllers.ConfirmController{}, "get:ConfirmUser")

	beego.Router("/forgot", &controllers.ForgotController{}, "get:ForgotForm")
	beego.Router("/forgot", &controllers.ForgotController{}, "post:ResetLink")

	beego.Router("/reset", &controllers.ResetController{}, "get:ResetForm")
	beego.Router("/reset", &controllers.ResetController{}, "get:ResetPassword")

	beego.Router("/login", &controllers.LoginController{}, "get:LoginForm")
	beego.Router("/login", &controllers.LoginController{}, "post:LoginUser")

	beego.Router("/logout", &controllers.LogoutController{}, "get:LogoutUser")
}

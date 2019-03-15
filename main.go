package main

import (
	"os"

	// "git.urantiatech.com/auth/login/middlewares"
	_ "git.urantiatech.com/auth/login/routers"
	"github.com/urantiatech/beego"
)

func main() {
	beego.AddFuncMap("getenv", os.Getenv)
	beego.AddFuncMap("siteName", siteName)
	beego.AddFuncMap("logo", logo)
	beego.AddFuncMap("domain", domain)
	beego.AddFuncMap("backgroundImage", backgroundImage)
	beego.AddFuncMap("backgroundColor", backgroundColor)
	beego.Run()
	//beego.RunWithMiddleWares("", middlewares.AuthMiddleware{})
}

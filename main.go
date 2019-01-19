package main

import (
	_ "git.urantiatech.com/auth/login/routers"
	"github.com/urantiatech/beego"
)

func main() {
	beego.AddFuncMap("siteName", siteName)
	beego.AddFuncMap("logo", logo)
	beego.AddFuncMap("domain", domain)
	beego.AddFuncMap("backgroundImage", backgroundImage)
	beego.AddFuncMap("backgroundColor", backgroundColor)
	beego.Run()
}

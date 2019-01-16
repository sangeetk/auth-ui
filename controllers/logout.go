package controllers

import (
	"github.com/urantiatech/beego"
)

// LogoutController provides user functions
type LogoutController struct {
	beego.Controller
}

// LogoutUser logs the current user out
func (c *LogoutController) LogoutUser() {
	c.TplName = "logout.tpl"
}
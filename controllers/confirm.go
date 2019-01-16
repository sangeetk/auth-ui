package controllers

import (
	"github.com/urantiatech/beego"
)

// ConfirmController provides user functions
type ConfirmController struct {
	beego.Controller
}

// ConfirmUser the user account after registeration
func (c *ConfirmController) ConfirmUser() {
	c.TplName = "confirm.tpl"
}

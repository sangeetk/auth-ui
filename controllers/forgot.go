package controllers

import (
	"github.com/urantiatech/beego"
)

// ForgotController provides user functions
type ForgotController struct {
	beego.Controller
}

// ForgotForm shows the form
func (c *ForgotController) ForgotForm() {
	c.TplName = "forgot.tpl"
}

// ResetLink emails the reset link
func (c *ForgotController) ResetLink() {
	c.TplName = "forgot.tpl"
}

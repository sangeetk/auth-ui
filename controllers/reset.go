package controllers

import (
	"github.com/urantiatech/beego"
)

// ResetController provides user functions
type ResetController struct {
	beego.Controller
}

// ResetForm shows reset form
func (c *ResetController) ResetForm() {
	c.TplName = "reset.tpl"
}

// ResetPassword resets the password
func (c *ResetController) ResetPassword() {
	c.TplName = "reset.tpl"
}

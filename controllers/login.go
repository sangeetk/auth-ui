package controllers

import (
	"log"

	"github.com/urantiatech/beego"
)

// LoginController provides login functions
type LoginController struct {
	beego.Controller
}

// LoginForm provides login form
func (c *LoginController) LoginForm() {
	c.TplName = "login.tpl"
	log.Println("Login Form")
}

// LoginUser authenticates the user
func (c *LoginController) LoginUser() {
	c.TplName = "login.tpl"
	log.Println("Login User")
}

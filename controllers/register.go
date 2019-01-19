package controllers

import (
	"log"

	"github.com/urantiatech/beego"
)

// RegisterController provides registration functions
type RegisterController struct {
	beego.Controller
}

// RegisterForm new user
func (c *RegisterController) RegisterForm() {
	c.TplName = "register1.tpl"

	log.Println("Register Form")
}

// RegisterUser new user
func (c *RegisterController) RegisterUser() {
	c.TplName = "register.tpl"

	log.Println("Register User")
}

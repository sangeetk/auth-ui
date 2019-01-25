package controllers

import (
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// ConfirmController provides new user confirmation functions
type ConfirmController struct {
	beego.Controller
}

// ConfirmUser the user account after registeration
func (c *ConfirmController) ConfirmUser() {
	c.TplName = "confirm.tpl"
	flash := beego.NewFlash()

	token := c.GetString("token")
	if token == "" {
		flash.Error("Invalid activation link.")
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	var req = authapi.ConfirmRequest{ConfirmToken: token}
	resp, err := authapi.Confirm(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		c.Redirect("/", http.StatusSeeOther)
		return
	}
	if resp.Err != "" {
		flash.Error("Invalid or expired link, sorry cannot activate the account.")
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	flash.Success("Account successfully activated, you may now <a href=\"/auth/login\">login</a> to your account.")
	c.Redirect("/", http.StatusSeeOther)
	return
}

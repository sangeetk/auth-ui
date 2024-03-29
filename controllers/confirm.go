package controllers

import (
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/astaxie/beego"
)

// ConfirmController provides new user confirmation functions
type ConfirmController struct {
	beego.Controller
}

// ConfirmUser the user account after registeration
func (c *ConfirmController) ConfirmUser() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	flash := beego.NewFlash()

	token := c.GetString("token")
	if token == "" {
		flash.Error("Invalid activation link.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	var req = authapi.ConfirmRequest{ConfirmToken: token}
	resp, err := authapi.Confirm(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}
	if resp.Err != "" {
		flash.Error("Invalid or expired link, unable to activate the account.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	flash.Success("Account activated")
	flash.Store(&c.Controller)
	c.Redirect("/auth/login", http.StatusSeeOther)
	return
}

package controllers

import (
	"net/http"
	"net/mail"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// ForgotController provides forgot password functions
type ForgotController struct {
	beego.Controller
}

// ForgotForm shows the form
func (c *ForgotController) ForgotForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.TplName = "page/forgot.tpl"
}

// ResetLink emails the reset link
func (c *ForgotController) ResetLink() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.ForgotRequest
	var err error
	flash := beego.NewFlash()

	email := c.GetString("email")

	var e *mail.Address
	if e, err = mail.ParseAddress(email); err != nil {
		c.Data["Error"] = "Invalid email address"
		c.TplName = "page/forgot.tpl"
		return
	}

	req.Username = e.Address
	req.Domain = os.Getenv("DOMAIN")

	resp, err := authapi.Forgot(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	// Check for error
	if resp.Err != "" {
		flash.Error(resp.Err)
		flash.Store(&c.Controller)
		c.TplName = "error.tpl"
		return
	}

	// Render the next form
	flash.Success("Further instruction to reset your password has been emailed to you.")
	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusSeeOther)
}

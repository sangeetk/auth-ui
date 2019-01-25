package controllers

import (
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
	c.TplName = "forgot.tpl"
}

// ResetLink emails the reset link
func (c *ForgotController) ResetLink() {
	var req authapi.ForgotRequest
	var err error

	email := c.GetString("email")

	var e *mail.Address
	if e, err = mail.ParseAddress(email); err != nil {
		c.Data["Error"] = "Invalid email address"
	}
	req.Username = e.Address
	req.Domain = os.Getenv("DOMAIN")

	resp, err := authapi.Forgot(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		c.Data["Error"] = "Unable to process your request. Please try again after some time."
		c.TplName = "error.tpl"
		return
	}

	// Check for error
	if resp.Err != "" {
		c.Data["Error"] = resp.Err
		c.TplName = "error.tpl"
		return
	}

	// Render the next form
	c.Data["Message"] = "Further instruction to reset your password has been emailed to you."
	c.TplName = "thankyou.tpl"
	return

}

package controllers

import (
	"log"
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/astaxie/beego"
)

// ResetController provides reset functions
type ResetController struct {
	beego.Controller
}

// ResetForm shows reset form
func (c *ResetController) ResetForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	// Show the password rest form on GET request
	flash := beego.NewFlash()

	token := c.GetString("token")
	if token == "" {
		flash.Error("Invalid link.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	c.Data["Token"] = token
	c.TplName = "page/reset.tpl"
	return
}

// ResetPassword resets the password
func (c *ResetController) ResetPassword() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.ResetRequest
	flash := beego.NewFlash()

	req.ResetToken = c.GetString("token")
	req.NewPassword = c.GetString("password")
	password2 := c.GetString("password2")

	log.Printf("Password=[%s]\n", req.NewPassword)
	log.Printf("Password2=[%s]\n", password2)

	if req.NewPassword != password2 {
		c.Data["Error"] = "Passwords do not match"
		c.Data["Token"] = req.ResetToken
		c.TplName = "page/reset.tpl"
		return
	}

	resp, err := authapi.Reset(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	// Check for registration error
	if resp.Err != "" {
		flash.Error("Invalid or expired link, please try <a href=\"/auth/forgot\">forgot password</a> again.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	flash.Success("Your password been reset successfully, you may now <a href=\"/auth/login\">login</a> using the new password.")
	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusSeeOther)
	return
}

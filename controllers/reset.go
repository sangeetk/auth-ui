package controllers

import (
	"log"
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// ResetController provides reset functions
type ResetController struct {
	beego.Controller
}

// ResetForm shows reset form
func (c *ResetController) ResetForm() {
	// Show the password rest form on GET request
	token := c.GetString("token")
	if token == "" {
		flash := beego.NewFlash()
		flash.Error("Invalid link.")
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	c.Data["Token"] = token
	c.TplName = "reset.tpl"
	return
}

// ResetPassword resets the password
func (c *ResetController) ResetPassword() {
	var req authapi.ResetRequest
	req.ResetToken = c.GetString("token")
	req.NewPassword = c.GetString("password")
	password2 := c.GetString("password2")

	log.Printf("Password=[%s]\n", req.NewPassword)
	log.Printf("Password2=[%s]\n", password2)

	if req.NewPassword != password2 {
		c.Data["Error"] = "Passwords do not match"
		c.Data["Token"] = req.ResetToken
		c.TplName = "reset.tpl"
		return
	}

	resp, err := authapi.Reset(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		log.Println(err.Error())
		// AUTH is unreachable
		c.Data["Error"] = "Unable to process your request. Please try again after some time."
		c.TplName = "error.tpl"
		return
	}

	flash := beego.NewFlash()

	// Check for registration error
	if resp.Err != "" {
		flash.Error("Invalid or expired link, please try <a href=\"/auth/forgot\">forgot password</a> again.")
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	flash.Success("Your password been reset successfully, you may now <a href=\"/auth/login\">login</a> using the new password.")
	c.Redirect("/", http.StatusSeeOther)
	return

}

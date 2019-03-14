package controllers

import (
	"log"
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// LoginController provides login functions
type LoginController struct {
	beego.Controller
}

// LoginForm provides login form
func (c *LoginController) LoginForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.TplName = "login.tpl"
}

// LoginUser authenticates the user
func (c *LoginController) LoginUser() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.LoginRequest
	flash := beego.NewFlash()

	req.Username = c.GetString("email")
	req.Password = c.GetString("password")
	req.Domain = os.Getenv("DOMAIN")

	resp, err := authapi.Login(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		log.Println(err)
		// AUTH is unreachable
		c.TplName = "500.tpl"
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", http.StatusSeeOther)
		return
	}

	// Check for login error
	if resp.Err != "" {
		flash.Error(resp.Err)
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", http.StatusSeeOther)
		return
	}

	// Set the Access & Refresh Tokens in Cookie
	c.SetSecureCookie(os.Getenv("SECRET"), "AccessToken", resp.AccessToken)
	c.SetSecureCookie(os.Getenv("SECRET"), "RefreshToken", resp.RefreshToken)

	// flash.Success("You have now logged in to the system.")
	// flash.Store(&c.Controller)

	// Redirect to Home page
	c.Redirect("/", http.StatusSeeOther)
	return

}

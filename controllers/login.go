package controllers

import (
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/astaxie/beego"
)

// LoginController provides login functions
type LoginController struct {
	beego.Controller
}

// LoginForm provides login form
func (c *LoginController) LoginForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.TplName = "page/login.tpl"
}

// LoginUser authenticates the user
func (c *LoginController) LoginUser() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.LoginRequest
	flash := beego.NewFlash()

	req.Username = c.GetString("email")
	req.Password = c.GetString("password")
	req.Domain = os.Getenv("DOMAIN")
	if c.GetString("remember") == "on" {
		req.RememberMe = true
	}

	resp, err := authapi.Login(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	// Check for login error
	if resp.Err != "" {
		c.Data["Error"] = resp.Err
		c.TplName = "page/login.tpl"
		return
	}

	// Set the Access & Refresh Tokens in Cookie
	c.SetSecureCookie(os.Getenv("SECRET"), "AccessToken", resp.AccessToken)
	c.SetSecureCookie(os.Getenv("SECRET"), "RefreshToken", resp.RefreshToken)
	c.SetSecureCookie(os.Getenv("SECRET"), "SuggestedURI", "/dashboard")

	// Redirect to Home page
	c.Redirect("/", http.StatusSeeOther)
	return

}

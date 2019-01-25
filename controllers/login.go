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
	c.TplName = "login.tpl"
}

// LoginUser authenticates the user
func (c *LoginController) LoginUser() {
	var req authapi.LoginRequest
	req.Username = c.GetString("email")
	req.Password = c.GetString("password")
	req.Domain = os.Getenv("DOMAIN")

	resp, err := authapi.Login(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		log.Println(err)
		// AUTH is unreachable
		c.TplName = "500.tpl"
		c.Data["Error"] = "Unable to process your request. Please try again after some time."
		return
	}

	// Check for login error
	if resp.Err != "" {
		c.Data["Error"] = resp.Err
		c.TplName = "login.tpl"
		return
	}

	// Set the Access & Refresh Tokens in Cookie
	c.SetSecureCookie(os.Getenv("SECRET"), "AccessToken", resp.AccessToken)
	c.SetSecureCookie(os.Getenv("SECRET"), "RefreshToken", resp.RefreshToken)

	flash := beego.NewFlash()
	flash.Success("Success:You have now logged in to the system.")

	// Redirect to Home page
	c.Redirect("/", http.StatusSeeOther)
	return

}

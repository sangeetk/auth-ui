package controllers

import (
	"net/http"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// LogoutController provides logout functions
type LogoutController struct {
	beego.Controller
}

// LogoutUser logs the current user out
func (c *LogoutController) LogoutUser() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.LogoutRequest
	flash := beego.NewFlash()

	req.AccessToken, _ = c.GetSecureCookie(os.Getenv("SECRET"), "AccessToken")
	req.RefreshToken, _ = c.GetSecureCookie(os.Getenv("SECRET"), "RefreshToken")

	_, _ = authapi.Logout(&req, os.Getenv("AUTH_SVC"))

	// Set Access & Refresh tokens to empty values
	c.SetSecureCookie(os.Getenv("SECRET"), "AccessToken", "")
	c.SetSecureCookie(os.Getenv("SECRET"), "RefreshToken", "")

	flash.Success("You have been logged out of the system.")
	flash.Store(&c.Controller)

	// Redirect to Home page
	c.Redirect("/", http.StatusSeeOther)
	return
}

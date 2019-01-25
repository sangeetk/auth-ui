package controllers

import (
	"log"
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
	var req authapi.LogoutRequest

	req.AccessToken, _ = c.GetSecureCookie(os.Getenv("SECRET"), "AccessToken")
	req.RefreshToken, _ = c.GetSecureCookie(os.Getenv("SECRET"), "RefreshToken")

	resp, err := authapi.Logout(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		log.Println(resp.Err)
	}

	// Set Access & Refresh tokens to empty values
	c.SetSecureCookie(os.Getenv("SECRET"), "AccessToken", "")
	c.SetSecureCookie(os.Getenv("SECRET"), "RefreshToken", "")

	flash := beego.NewFlash()
	flash.Success("You have been logged out of the system.")

	// Redirect to Home page
	c.Redirect("/", http.StatusSeeOther)
	return
}

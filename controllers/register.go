package controllers

import (
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/astaxie/beego"
)

// RegisterController provides registration functions
type RegisterController struct {
	beego.Controller
}

// RegisterForm new user
func (c *RegisterController) RegisterForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.TplName = "page/register1.tpl"
}

// RegisterUser new user
func (c *RegisterController) RegisterUser() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var err error
	flash := beego.NewFlash()
	step := c.GetString("step")

	// Register user
	if step == "1" {
		var req = &authapi.RegisterRequest{}
		var err error

		req.FirstName = c.GetString("fname")
		req.LastName = c.GetString("lname")
		req.Name = req.FirstName + " " + req.LastName

		email := c.GetString("email")

		req.Password = c.GetString("password")
		password2 := c.GetString("password2")

		if req.Password != password2 {
			c.Data["Error"] = "Passwords do not match"
			c.TplName = "page/register1.tpl"
			return
		}

		var e *mail.Address
		if e, err = mail.ParseAddress(email); err != nil {
			c.Data["Error"] = "Invalid email address"
			c.TplName = "page/register1.tpl"
			return
		}

		req.Email = e.Address
		req.Username = req.Email
		req.Domain = os.Getenv("DOMAIN")
		req.CacheReq = true

		resp, err := authapi.Register(req, os.Getenv("AUTH_SVC"))
		if err != nil {
			// AUTH is unreachable
			log.Println("Error:", err)
			c.Data["Error"] = "Unable to process your request. Please try again after some time."
			c.TplName = "page/error.tpl"
			return
		}

		// Check for registration error
		if resp.Err != "" {
			if resp.Err == "Already Registered" {
				c.Data["Error"] = "You are already registered, please login to continue"
				c.TplName = "page/login.tpl"
			} else {
				c.Data["Error"] = resp.Err
				c.TplName = "page/error.tpl"
			}
			return
		}

		// Render the next form
		c.Data["Token"] = resp.CacheKey
		c.TplName = "page/register2.tpl"
		return
	}

	// Update user address
	if step == "2" {
		var req = &authapi.RegisterRequest{}
		birthday := c.GetString("birthday")
		log.Println(birthday)
		req.Birthday, err = time.Parse("2006-01-02", birthday)
		if err != nil {
			c.Data["Error"] = "Invalid Date"
			c.TplName = "page/register2.tpl"
			return
		}

		req.Address.AddressType = "default"
		req.Address.Address1 = c.GetString("address")
		req.Address.City = c.GetString("city")
		req.Address.State = c.GetString("state")
		req.Address.Country = c.GetString("country")
		req.CacheKey = c.GetString("token")
		req.CacheReq = true

		resp, err := authapi.Register(req, os.Getenv("AUTH_SVC"))
		if err != nil {
			// AUTH is unreachable
			c.Data["Error"] = "Unable to process your request. Please try again after some time."
			c.TplName = "page/error.tpl"
			return
		}

		// Check for update error
		if resp.Err != "" {
			c.Data["Error"] = resp.Err
			c.TplName = "page/error.tpl"
			return
		}

		// Render the next form
		c.Data["Token"] = resp.CacheKey
		c.TplName = "page/register3.tpl"
		return
	}

	// Update user profile
	if step == "3" {
		var req = &authapi.RegisterRequest{}
		req.Profile = make(map[string]string)
		req.Profile["profession"] = c.GetString("profession")
		req.Profile["introduction"] = c.GetString("introduction")
		req.CacheKey = c.GetString("token")
		req.CacheReq = false

		resp, err := authapi.Register(req, os.Getenv("AUTH_SVC"))
		if err != nil {
			// AUTH is unreachable
			c.Data["Error"] = "Unable to process your request. Please try again after some time."
			c.TplName = "page/error.tpl"
			return
		}

		// Check for update error
		if resp.Err != "" {
			c.Data["Error"] = resp.Err
			c.TplName = "page/error.tpl"
			return
		}

		// Render the thankyou screen
		flash.Success("Further instruction to activate your account has been emailed to you.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
	}
}

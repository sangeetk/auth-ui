package controllers

import (
	"log"
	"net/mail"
	"os"
	"time"

	authapi "git.urantiatech.com/auth/auth/api"
	"github.com/urantiatech/beego"
)

// RegisterController provides registration functions
type RegisterController struct {
	beego.Controller
}

// RegisterForm new user
func (c *RegisterController) RegisterForm() {
	c.TplName = "register1.tpl"

	log.Println("Register Form")
}

// RegisterUser new user
func (c *RegisterController) RegisterUser() {
	var err error
	step := c.GetString("step")

	// Register user
	if step == "1" {
		var req authapi.RegisterRequest
		var err error

		req.FirstName = c.GetString("fname")
		req.LastName = c.GetString("lname")
		req.Name = req.FirstName + " " + req.LastName

		email := c.GetString("email")

		req.Password = c.GetString("password")
		password2 := c.GetString("password2")

		if req.Password != password2 {
			c.Data["Error"] = "Passwords do not match"
			c.TplName = "register1.tpl"
			return
		}

		var e *mail.Address
		if e, err = mail.ParseAddress(email); err != nil {
			c.Data["Error"] = "Invalid email address"
			c.TplName = "register1.tpl"
			return
		}

		req.Email = e.Address
		req.Username = req.Email
		req.Domain = os.Getenv("DOMAIN")

		resp, err := authapi.Register(&req, os.Getenv("AUTH_SVC"))
		if err != nil {
			log.Println(err.Error())
			// AUTH is unreachable
			c.Data["Error"] = "Unable to process your request. Please try again after some time."
			c.TplName = "500.tpl"
			return
		}

		// Check for registration error
		if resp.Err != "" {
			c.Data["Error"] = resp.Err
			c.TplName = "register1.tpl"
			return
		}

		// Render the next form
		c.Data["Token"] = resp.UpdateToken
		c.TplName = "register2.tpl"
		return
	}

	// Update user address
	if step == "2" {
		var req authapi.UpdateRequest
		req.UpdateToken = c.GetString("token")
		birthday := c.GetString("birthday")
		log.Println(birthday)
		req.Birthday, err = time.Parse("2 January, 2006", birthday)
		if err != nil {
			c.Data["Error"] = "Invalid Date"
			c.TplName = "register2.tpl"
			return
		}

		req.Address.AddressType = "default"
		req.Address.Address1 = c.GetString("address")
		req.Address.City = c.GetString("city")
		req.Address.State = c.GetString("state")
		req.Address.Country = c.GetString("country")

		resp, err := authapi.Update(&req, os.Getenv("AUTH_SVC"))
		if err != nil {
			log.Println(err.Error())
			// AUTH is unreachable
			c.Data["Error"] = "Network error"
			c.TplName = "500.tpl"
			return
		}

		// Check for update error
		if resp.Err != "" {
			c.Data["Error"] = resp.Err
			c.TplName = "register2.tpl"
			return
		}

		// Render the next form
		c.Data["Token"] = resp.UpdateToken
		c.TplName = "register3.tpl"
		return
	}

	// Update user profile
	if step == "3" {
		var req authapi.UpdateRequest
		req.UpdateToken = c.GetString("token")
		req.Profile = make(map[string]string)
		req.Profile["profession"] = c.GetString("profession")
		req.Profile["introduction"] = c.GetString("introduction")

		resp, err := authapi.Update(&req, os.Getenv("AUTH_SVC"))
		if err != nil {
			log.Println(err.Error())
			// AUTH is unreachable
			c.Data["Error"] = "Network error"
			c.TplName = "500.tpl"
			return
		}

		// Check for update error
		if resp.Err != "" {
			c.Data["Error"] = resp.Err
			c.TplName = "500.tpl"
			return
		}

		// Render the thankyou screen
		c.Data["Message"] = "Further instruction to activate your account has been emailed to you."
		c.TplName = "thankyou.tpl"

	}
}

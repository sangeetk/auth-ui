package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"

	authapi "git.urantiatech.com/auth/auth/api"
	"git.urantiatech.com/auth/login/emails"
	mailapi "git.urantiatech.com/mail/mail/api"
	"github.com/urantiatech/beego"
)

// ForgotController provides forgot password functions
type ForgotController struct {
	beego.Controller
}

// ForgotForm shows the form
func (c *ForgotController) ForgotForm() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.TplName = "page/forgot.tpl"
}

// ResetLink emails the reset link
func (c *ForgotController) ResetLink() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data

	var req authapi.ForgotRequest
	var err error
	flash := beego.NewFlash()

	email := c.GetString("email")

	var e *mail.Address
	if e, err = mail.ParseAddress(email); err != nil {
		c.Data["Error"] = "Invalid email address"
		c.TplName = "page/forgot.tpl"
		return
	}

	req.Username = e.Address
	req.Domain = os.Getenv("DOMAIN")

	resp, err := authapi.Forgot(&req, os.Getenv("AUTH_SVC"))
	if err != nil {
		// AUTH is unreachable
		flash.Error("Unable to process your request. Please try again after some time.")
		flash.Store(&c.Controller)
		c.Redirect("/", http.StatusSeeOther)
		return
	}

	// Check for error
	if resp.Err != "" {
		if resp.Err == "Not Found" {
			c.Data["Error"] = "This email address is not registered in our database."
		} else {
			c.Data["Error"] = resp.Err
		}
		c.TplName = "page/error.tpl"
		return
	}

	// Preapre PasswordReset link mail
	data := make(map[string]interface{})
	data["Domain"] = os.Getenv("DOMAIN")
	data["Token"] = resp.ResetToken
	data["Name"] = resp.FirstName

	var html bytes.Buffer
	if err := emails.Emails[emails.PasswordReset].Execute(&html, data); err != nil {
		log.Fatal(err.Error())
	}

	// Send the PasswordReset link mail
	mail := mailapi.Mail{
		From:    fmt.Sprintf("%s <contact@%s>", os.Getenv("SITE_NAME"), os.Getenv("DOMAIN")),
		To:      fmt.Sprintf("%s", resp.Email),
		Subject: fmt.Sprintf("%s password reset link", os.Getenv("SITE_NAME")),
		HTML:    html.String(),
	}
	err = mailapi.SendMail(&mail, os.Getenv("MAIL_SVC"))

	// Render the next form
	flash.Success("Further instruction to reset your password has been emailed to you.")
	flash.Store(&c.Controller)
	c.Redirect("/", http.StatusSeeOther)
}

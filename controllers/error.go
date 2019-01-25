package controllers

import (
	"github.com/urantiatech/beego"
)

// ErrorController handles 404, 401, 403, 500, 503 errors
type ErrorController struct {
	beego.Controller
}

// Error404 handles 404
func (c *ErrorController) Error404() {
	c.Data["content"] = "page not found"
	c.TplName = "404.tpl"
}

// Error500 handles 500
func (c *ErrorController) Error500() {
	c.Data["content"] = "internal server error"
	c.TplName = "500.tpl"
}

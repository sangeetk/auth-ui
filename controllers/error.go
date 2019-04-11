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
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.Data["content"] = "page not found"
	c.TplName = "page/404.tpl"
}

// Error500 handles 500
func (c *ErrorController) Error500() {
	c.Data["Flash"] = beego.ReadFromRequest(&c.Controller).Data
	c.Data["content"] = "internal server error"
	c.TplName = "page/500.tpl"
}

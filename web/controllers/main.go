package controllers

import (
	"lazyper/web/views"

	"github.com/kataras/iris/v12"
)

type MainController struct {
	Ctx iris.Context
}

//主页
func (c *MainController) Get() {
	c.Ctx.WriteString(views.HtmlPage)
}

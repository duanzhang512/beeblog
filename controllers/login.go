package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	if c.Input().Get("exit") == "true" {
		c.DelSession("uname")
		c.DelSession("pwd")
		c.DestroySession()
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")

	if uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd") {

		c.SetSession("uname", uname)
		c.SetSession("pwd", pwd)
	}

	c.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {

	uname := ctx.Input.CruSession.Get("uname")
	pwd := ctx.Input.CruSession.Get("pwd")

	if uname == nil || pwd == nil {
		return false
	}

	return uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd")
}

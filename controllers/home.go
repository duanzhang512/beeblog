package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(c.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories
}

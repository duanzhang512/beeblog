package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {

	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topic"] = topics
}

func (c *TopicController) Post() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	tid := c.Input().Get("tid")
	title := c.Input().Get("title")
	category := c.Input().Get("category")
	content := c.Input().Get("content")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}

	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}

func (c *TopicController) Add() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "topic_add.html"
}

func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/topic", 302)
}

func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"

	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
	c.Data["Tid"] = tid
}

func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 302)
		return
	}

	c.Data["Topic"] = topic
}

package controllers

import (
	"github.com/astaxie/beego"
)

type ListBulkController struct {
	beego.Controller
}

func (c *ListBulkController) Get() {
	c.TplName = "list_bulk.html"
}

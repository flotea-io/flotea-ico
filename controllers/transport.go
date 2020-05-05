package controllers

import (
	"github.com/astaxie/beego"
)


type TransportController struct {
	beego.Controller
}

func (c *TransportController) Get() {
    c.Layout = "layouts/main.tpl"
    c.LayoutSections = make(map[string]string)
    c.LayoutSections["Header"] = "layouts/header.tpl"
    c.LayoutSections["Footer"] = "layouts/footer.tpl"
    c.TplName = "transport.tpl"
}
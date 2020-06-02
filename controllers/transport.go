/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

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

/*
* Project: FLOTEA - Decentralized passenger transport system
* Copyright (c) 2020 Flotea, All Rights Reserved
* For conditions of distribution and use, see copyright notice in LICENSE
*/

package controllers

import (
	"github.com/astaxie/beego"
)

type VoteController struct {
	beego.Controller
}

func (c *VoteController) setLayout() {
	c.Layout = "layouts/main.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "layouts/header.tpl"
	c.LayoutSections["Footer"] = "layouts/footer.tpl"
	c.Data["TplName"] = c.TplName[:len(c.TplName)-4]
}

func (c *VoteController) Carriers() {
	c.TplName = "carriers.tpl"
	var Javascript = []string{
		"https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/carriers.js",
	}
	c.Data["Javascript"] = Javascript
	c.setLayout()
}

func (c *VoteController) Add() {
	c.TplName = "addCarrier.tpl"
	var Javascript = []string{
		"https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/voting.js",
	}
	c.Data["Javascript"] = Javascript
	c.setLayout()
}

func (c *VoteController) Remove() {
	c.TplName = "removeCarrier.tpl"
	var Javascript = []string{
		"https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/voting.js",
	}
	c.Data["Javascript"] = Javascript
	c.setLayout()
}

func (c *VoteController) Applicants() {
	c.TplName = "applicants.tpl"
	var Javascript = []string{
		"https://cdnjs.cloudflare.com/ajax/libs/jquery-jgrowl/1.4.6/jquery.jgrowl.min.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_main_voting.js",
		"/static/js/vendors_applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/vendors_applicants_icoVoting.js",
		"/static/js/vendors_applicants.js",
		"/static/js/applicants_carriers_icoBuy_icoVoting_voting.js",
		"/static/js/applicants.js",
	}
	c.Data["Javascript"] = Javascript
	c.setLayout()
}

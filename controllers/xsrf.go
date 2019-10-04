package controllers

import "github.com/astaxie/beego"

type MyXSRF struct {
	beego.Controller
}


func (this * MyXSRF) Get() {
	this.Data["xsrf"] = this.XSRFToken()
	this.TplName = "XSRF.html"
}
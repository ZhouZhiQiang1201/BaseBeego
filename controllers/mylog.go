package controllers

import "github.com/astaxie/beego"

type MyLog struct {
	beego.Controller
}

func (this *MyLog) Get() {

	this.TplName = "log.html"
}
package controllers

import "github.com/astaxie/beego"

type MyFlash struct {
	beego.Controller
}

func (this * MyFlash) Get(){
	flash:=beego.NewFlash()
	flash.Error("flash的使用")
	this.TplName = "flash.html"
}
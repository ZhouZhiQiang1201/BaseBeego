package controllers

import "github.com/astaxie/beego"

type MyJson struct {
	beego.Controller
}

func (this *MyJson) Get() {
	data := map[string]string{"name": "beego", "age": "1"}
	this.Data["json"] = data
	this.ServeJSON()
}
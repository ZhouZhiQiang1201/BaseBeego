package controllers

import "github.com/astaxie/beego"

type MySession struct {
	beego.Controller
}

func (this *MySession) Get() {
	this.SetSession("mybeego", "zzq")
	this.TplName = "session.html"

}
